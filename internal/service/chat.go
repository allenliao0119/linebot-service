package service

import "context"

type ChatService interface {
	GetResponse(ctx context.Context, userMessage string, userID string) (string, error)
}