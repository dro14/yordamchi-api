package provider

import (
	"context"
	"os"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

const followUpInstruction = "Come up with three follow-up questions that the user is likely to ask next. The questions should be in the same language as the user's message."

func (p *Provider) FollowUp(request *models.Request) (string, error) {
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

	systemInstruction := systemInstructions[request.Language]
	if request.SystemInstruction != "" {
		systemInstruction += "\n\n" + request.SystemInstruction
	}
	request.SystemInstruction = systemInstruction
	request.Model = model

	maxOutputTokens := int32(3072)
	temperature := new(float32)
	*temperature = 1.0
	thinkingBudget := new(int32)
	*thinkingBudget = 0

	response, err := p.client.Models.GenerateContent(
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
			ResponseMIMEType: "application/json",
			ResponseSchema: &genai.Schema{
				Type: genai.TypeArray,
				Items: &genai.Schema{
					Type: genai.TypeString,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return response.Text(), nil
}
