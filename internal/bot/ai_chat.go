package bot

import (
	"context"
	"fmt"
)

type AIChatBot struct{
	apiKey string
}

func NewAIChatBot(apiKey string) *AIChatBot {
	return &AIChatBot{
		apiKey: apiKey,
	}
}

func (b *AIChatBot) GetResponse(ctx context.Context, userMessage string, userID string) (string, error) {
	return "", fmt.Errorf("AI chat not implement yet")
}