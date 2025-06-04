package provider

import (
	"context"
	"iter"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

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

	var systemInstruction *genai.Content
	if request.SystemInstruction != "" {
		systemInstruction = &genai.Content{
			Parts: []*genai.Part{
				{Text: request.SystemInstruction},
			},
		}
	}

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
			SystemInstruction: systemInstruction,
			MaxOutputTokens:   maxOutputTokens,
			Temperature:       temperature,
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: thinkingBudget,
			},
		},
	)
}
