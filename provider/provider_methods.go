package provider

import (
	"context"
	"iter"
	"os"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

const model = "gemini-2.5-flash"

var systemInstructions = map[string]string{
	"uz":      "Sening isming Yordamchi, matn va rasmlarni tushuna oladigan, xushmuomala chatbotsan. Javob juda uzun bo'lmasligi kerak. Standart til: O'zbekcha (lotin)",
	"uz_Cyrl": "Сенинг исминг Yordamchi, матн ва расмларни тушуна оладиган, хушмуомала чатботсан. Жавоб жуда узун бўлмаслиги керак. Стандарт тил: Ўзбекча (кирил)",
	"ru":      "Ты являешься дружелюбным чатботом под именем Yordamchi, который понимает текст и изображения. Ответ не должен быть слишком длинным. Язык по умолчанию: Русский",
	"en":      "You are a friendly chatbot named Yordamchi, which understands text and images. Response should not be too long. Default language: English",
}

func (p *Provider) ContentStream(request *models.Request) iter.Seq2[*genai.GenerateContentResponse, error] {
	var contents []*genai.Content
	for _, message := range request.Contents {
		parts := []*genai.Part{}
		isNotEmpty := false
		if len(message.Text) > 0 {
			parts = append(parts, &genai.Part{Text: message.Text})
			isNotEmpty = true
		}
		if len(message.Images) > 0 {
			for _, image := range message.Images {
				imageData, _ := os.ReadFile("rasmlar/" + image)
				parts = append(parts, &genai.Part{
					InlineData: &genai.Blob{
						MIMEType: "image/jpeg",
						Data:     imageData,
					},
				})
			}
			isNotEmpty = true
		}
		if len(message.Calls) > 0 {
			for _, functionCall := range message.Calls {
				parts = append(parts, &genai.Part{
					FunctionCall: functionCall,
				})
			}
			isNotEmpty = true
		}
		if len(message.Responses) > 0 {
			for _, functionResponse := range message.Responses {
				parts = append(parts, &genai.Part{
					FunctionResponse: functionResponse,
				})
			}
			isNotEmpty = true
		}
		if isNotEmpty {
			contents = append(contents, &genai.Content{
				Parts: parts,
				Role:  message.Role,
			})
		}
	}

	systemInstruction := systemInstructions[request.Language]
	if request.SystemInstruction != "" {
		systemInstruction += "\n\n" + request.SystemInstruction
	}
	request.SystemInstruction = systemInstruction
	request.Model = model

	maxOutputTokens := int32(3072)
	temperature := new(float32)
	*temperature = 0.5
	thinkingBudget := new(int32)
	*thinkingBudget = 0

	return p.client.Models.GenerateContentStream(
		context.Background(),
		model,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.Text(request.SystemInstruction)[0],
			MaxOutputTokens:   maxOutputTokens,
			Temperature:       temperature,
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: thinkingBudget,
			},
			// Tools: []*genai.Tool{{
			// 	FunctionDeclarations: []*genai.FunctionDeclaration{{
			// 		Name:        "web_search",
			// 		Description: "Provides real-time, up-to-date information",
			// 		Parameters: &genai.Schema{
			// 			Type: "object",
			// 			Properties: map[string]*genai.Schema{
			// 				"query":         {Type: "string"},
			// 				"language_code": {Type: "string"},
			// 				"result_count":  {Type: "integer"},
			// 			},
			// 			Required: []string{"query", "language_code", "result_count"},
			// 		},
			// 	}},
			// }},
		},
	)
}
