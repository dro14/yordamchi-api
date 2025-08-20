package provider

import (
	"context"
	"iter"
	"os"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

const systemInstruction = "Your name is Yordamchi. Default language: "

var languages = map[string]string{
	"uz": "Uzbek",
	"en": "English",
	"ru": "Russian",
}

func (p *Provider) ContentStream(request *models.Request) iter.Seq2[*genai.GenerateContentResponse, error] {
	var contents []*genai.Content
	for _, message := range request.Contents {
		parts := []*genai.Part{}
		if len(message.Text) > 0 {
			parts = append(parts, &genai.Part{Text: message.Text})
		}
		for _, image := range message.Images {
			imageData, _ := os.ReadFile("rasmlar/" + image)
			parts = append(parts, &genai.Part{
				InlineData: &genai.Blob{
					MIMEType: "image/jpeg",
					Data:     imageData,
				},
			})
		}
		for _, call := range message.Calls {
			parts = append(parts, &genai.Part{
				FunctionCall: call,
			})
		}
		for _, response := range message.Responses {
			parts = append(parts, &genai.Part{
				FunctionResponse: response,
			})
		}
		contents = append(contents, &genai.Content{
			Parts: parts,
			Role:  message.Role,
		})
	}

	systemInstruction := systemInstruction + languages[request.Language]
	if request.SystemInstruction != "" {
		systemInstruction += "\n\n" + request.SystemInstruction
	}
	request.SystemInstruction = systemInstruction
	request.Model = model

	p.index = (p.index + 1) % len(p.clients)
	return p.clients[p.index].Models.GenerateContentStream(
		context.Background(),
		model,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.Text(systemInstruction)[0],
			MaxOutputTokens:   int32(4096),
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: new(int32),
			},
			Tools: []*genai.Tool{{
				FunctionDeclarations: []*genai.FunctionDeclaration{{
					Name:        "google_search",
					Description: "Provides real-time, up-to-date information",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"query": {Type: genai.TypeString},
						},
					},
				}},
			}},
		},
	)
}
