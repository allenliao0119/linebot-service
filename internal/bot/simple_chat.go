package bot

import (
	"context"
	"strings"
)

type SimpleChatBot struct{}

func NewSimpleChatBot() *SimpleChatBot {
	return &SimpleChatBot{}
}

func (b *SimpleChatBot) GetResponse(ctx context.Context, userMessage string, userID string) (string, error) {
	msg := strings.ToLower(strings.TrimSpace(userMessage))

	// Simple chat for test
	switch {
    case strings.Contains(msg, "你好") || strings.Contains(msg, "hi") || strings.Contains(msg, "hello"):
        return "Hello! 👋 How can I help you?", nil
    case strings.Contains(msg, "謝謝") || strings.Contains(msg, "感謝") || strings.Contains(msg, "thanks") || strings.Contains(msg, "thank"):
        return "You're welcome! 😊 Happy to help!", nil
    case strings.Contains(msg, "再見") || strings.Contains(msg, "bye") || strings.Contains(msg, "goodbye"):
        return "Goodbye! 👋 Looking forward to chatting again!", nil
    case strings.Contains(msg, "幫助") || strings.Contains(msg, "help"):
        return "I can answer your questions or just chat with you! 💬 Feel free to ask me anything!", nil
    default:
        return "I received your message: 「" + userMessage + "」\nI'm still learning and will get smarter! 🤖✨", nil
    }
}