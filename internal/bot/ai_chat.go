package bot

import (
	"context"
	"log"

	"github.com/allenliao0119/linebot-service/internal/helper"
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

	systemPrompt := "你是一個LINE Bot助手，請使用繁體中文回答用戶的問題。你應該友善，並提供清晰簡潔的回答。如果需要標示段落、數字或是要表示重點時，適度會使用icon、幫助用戶快速抓到重點"

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model:        b.model,
		Instructions: openai.String(systemPrompt),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(userMessage),
		},
		Tools: []responses.ToolUnionParam{
			responses.ToolParamOfWebSearch(responses.WebSearchToolTypeWebSearch),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	return helper.ConvertMarkdownToLineText(resp.OutputText()), nil
}