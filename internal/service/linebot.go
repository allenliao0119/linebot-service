package service

import (
	"context"
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type LineBotService struct {
	client *messaging_api.MessagingApiAPI
	chatService ChatService 
}

func NewLineBotService(channelToken string, chatService ChatService) (*LineBotService, error) {
	client, err := messaging_api.NewMessagingApiAPI(channelToken)
	if err != nil {
		return nil, err
	}

	return &LineBotService{
		client: client,
		chatService: chatService,
	}, nil
}

func (s *LineBotService) HandleEvent(ctx context.Context, event webhook.EventInterface) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
        return s.handleMessageEvent(ctx, e)
    case webhook.FollowEvent:
        return s.handleFollowEvent(ctx, e)
	default:
		log.Printf("Unsupported event type: %T", event)
        return nil
	}
}

func (s *LineBotService) handleMessageEvent(ctx context.Context, event webhook.MessageEvent) error {
    switch event.Message.(type) {
    case webhook.TextMessageContent:
        return s.handleTextMessage(ctx, event)
    default:
        log.Printf("Unsupported message type: %T", event.Message)
        return nil
    }
}

func (s *LineBotService) handleTextMessage(ctx context.Context, event webhook.MessageEvent) error {
	textMessage, ok := event.Message.(webhook.TextMessageContent)
	if !ok {
		return fmt.Errorf("failed to cast to TextMessageConent")
	}

	var userID string
    switch source := event.Source.(type) {
    case webhook.UserSource:
        userID = source.UserId
    case webhook.GroupSource:
        userID = source.UserId
    case webhook.RoomSource:
        userID = source.UserId
    default:
        userID = "unknown"
    }

    log.Printf("Received text message from user %s: %s", userID, textMessage.Text)

	replyMessage, err := s.chatService.GetResponse(ctx, textMessage.Text, userID)
	if err != nil {
		log.Printf("failed to get chat response: %v", err)
        replyMessage = "“Sorry, I’m not able to handle your message right now 😢"
	}

	_, err = s.client.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: event.ReplyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: replyMessage,
				},
			},
		},
	)
	if err != nil {
        log.Printf("failed to reply message: %v", err)
    }
	
	return nil
}

func (s *LineBotService) handleFollowEvent(ctx context.Context, event webhook.FollowEvent) error {
	if userSource, ok := event.Source.(webhook.UserSource); ok {
		log.Printf("user %s followed the bot", userSource.UserId) 
	}
	
	replyMessage := `
	👋 Welcome aboard!
	I'm your chatbot assistant, and it's great to meet you!
	Try sending me a message — I'll do my best to respond to your questions!
	`
	_, err := s.client.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: event.ReplyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: replyMessage,
				},
			},
		},
	)
	if err != nil {
		log.Printf("failed to reply message: %v", err)
	}
	return nil
}