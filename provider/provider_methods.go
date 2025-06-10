package provider

import (
	"context"
	"iter"
	"time"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

var systemInstructions = map[string]string{
	"uz":      "Sening isming Yordamchi, matn va rasmlarni tushuna oladigan, xushmuomala chatbotsan. ChuqurTech kompaniyasi tomonidan ishlab chiqilgansan. Standart til: O'zbekcha (lotin). Hozirgi vaqt: ",
	"uz_Cyrl": "Сенинг исминг Yordamchi, матн ва расмларни тушуна оладиган, хушмуомала чатботсан. ChuqurTech компанияси томонидан ишлаб чиқилгансан. Стандарт тил: Ўзбекча (кирил). Ҳозирги вақт: ",
	"ru":      "Ты являешься дружелюбным чатботом под именем Yordamchi, который понимает текст и изображения. Ты был разработан компанией ChuqurTech. Язык по умолчанию: Русский. Текущее время: ",
	"en":      "You are a friendly chatbot named Yordamchi, which understands text and images. You were developed by a company called ChuqurTech. Default language: English. Current time: ",
}

func (p *Provider) ContentStream(request *models.Request) iter.Seq2[*genai.GenerateContentResponse, error] {
	var contents []*genai.Content
	for _, message := range request.Contents {
		contents = append(contents, &genai.Content{
			Parts: []*genai.Part{
				{Text: message.Text},
			},
			Role: message.Role,
		})
	}

	systemInstruction := systemInstructions[request.Language] + time.Now().Format(time.DateTime)

	if request.SystemInstruction != "" {
		systemInstruction += "\n\n" + request.SystemInstruction
	}
	request.SystemInstruction = systemInstruction

	maxOutputTokens := int32(3072)
	temperature := new(float32)
	*temperature = 0.5
	thinkingBudget := new(int32)
	*thinkingBudget = 0

	return p.client.Models.GenerateContentStream(
		context.Background(),
		"gemini-2.5-flash-preview-05-20",
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.Text(request.SystemInstruction)[0],
			MaxOutputTokens:   maxOutputTokens,
			Temperature:       temperature,
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: thinkingBudget,
			},
		},
	)
}
