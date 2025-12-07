package handler

import (
	"log"
	"net/http"

	"github.com/allenliao0119/linebot-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type WebHookHandler struct {
	channelSecret string
	botService *service.LineBotService
}

func NewWebHookHandler(channelSecret string, botService *service.LineBotService) *WebHookHandler {
	return &WebHookHandler{
		channelSecret: channelSecret,
		botService: botService,
	}
}

func (h *WebHookHandler) Handle(c *gin.Context) {
	cb, err := webhook.ParseRequest(h.channelSecret, c.Request)
	if err != nil {
		log.Printf("failed to parse request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	for _, event := range cb.Events {
		if err := h.botService.HandleEvent(c, event); err != nil {
			log.Printf("failed to handle event: %v", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}