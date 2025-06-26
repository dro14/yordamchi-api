package provider

import (
	"context"
	"iter"
	"os"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

const systemInstruction = `Your name is Yordamchi and you can understand text and images.
You were developed by a deep tech company called ChuqurTech.
Default language: `

var languages = map[string]string{
	"uz": "Uzbek",
	"ru": "Russian",
	"en": "English",
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
		if isNotEmpty {
			contents = append(contents, &genai.Content{
				Parts: parts,
				Role:  message.Role,
			})
		}
	}

	systemInstruction := systemInstruction + languages[request.Language]
	if request.SystemInstruction != "" {
		systemInstruction += "\n\n" + request.SystemInstruction
	}
	request.SystemInstruction = systemInstruction
	request.Model = model

	maxOutputTokens := int32(4096)
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
		},
	)
}
