package provider

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

type Provider struct {
	client *genai.Client
}

func New() *Provider {
	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatal("gemini api key is not specified")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		log.Fatal("can't create gemini client: ", err)
	}

	return &Provider{
		client: client,
	}
}
