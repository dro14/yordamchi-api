package provider

import (
	"context"
	"os"

	"github.com/dro14/yordamchi-api/models"
	"google.golang.org/genai"
)

const followUpInstruction = `Come up with three follow-up questions that the user is likely to ask next.
The questions should be in the same language as the user's message.`

func (p *Provider) FollowUp(request *models.Request) (*genai.GenerateContentResponse, error) {
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

	request.SystemInstruction = followUpInstruction
	request.Model = model

	p.index = (p.index + 1) % len(p.clients)
	return p.clients[p.index].Models.GenerateContent(
		context.Background(),
		model,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.Text(followUpInstruction)[0],
			MaxOutputTokens:   int32(4096),
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: new(int32),
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
}
