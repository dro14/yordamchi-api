package provider

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

const model = "gemini-2.5-flash"

type Provider struct {
	Clients []*genai.Client
	index   int
}

func New() *Provider {
	ctx := context.Background()
	var clients []*genai.Client
	for i := 0; ; i++ {
		key := fmt.Sprintf("GOOGLE_API_KEY_%d", i)
		apiKey, ok := os.LookupEnv(key)
		if !ok {
			break
		}
		client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
		if err != nil {
			log.Printf("can't create google client: %s", err)
			continue
		}
		clients = append(clients, client)
	}

	return &Provider{
		Clients: clients,
		index:   0,
	}
}
