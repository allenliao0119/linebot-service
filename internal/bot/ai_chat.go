package bot

import (
	"context"
	"log"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type AIChatBot struct{
	apiKey string
	model  string
}

func NewAIChatBot(apiKey string, model string) *AIChatBot {
	if apiKey == "" {
		log.Fatal("OpenAI API key is required")
	}
	return &AIChatBot{
		apiKey: apiKey,
		model: model,
	}
}

func (b *AIChatBot) GetResponse(ctx context.Context, userMessage string, userID string) (string, error) {
	client := openai.NewClient(
		option.WithAPIKey(b.apiKey),
	)
	
	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model: b.model,
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(userMessage)},
	})
	if err != nil {
		panic(err.Error())
	}

	return resp.OutputText(), nil
}