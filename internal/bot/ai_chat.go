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

	systemPrompt := "你是一個 LINE Bot 助手，請使用繁體中文回答用戶的問題。你應該友善、樂於助人，並提供清晰簡潔的回答。"

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model:        b.model,
		Instructions: openai.String(systemPrompt),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{

			},
			OfString: openai.String(userMessage),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	return resp.OutputText(), nil
}