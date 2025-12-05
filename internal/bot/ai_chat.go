package bot

import (
	"context"
	"fmt"
)

type AIChatBot struct{
	apiKey string
	model  string
}

func NewAIChatBot(apiKey string, model string) *AIChatBot {
	return &AIChatBot{
		apiKey: apiKey,
		model: model,
	}
}

func (b *AIChatBot) GetResponse(ctx context.Context, userMessage string, userID string) (string, error) {
	return "", fmt.Errorf("AI chat not implement yet")
}