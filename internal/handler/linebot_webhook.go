package handler

import (
	"io"
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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	signature := c.GetHeader("X-Line-Signature")
	if !webhook.ValidateSignature(h.channelSecret, signature, body) {
		log.Println("invalid signature")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	cb, err := webhook.ParseRequest(h.channelSecret, c.Request)
	if err != nil {
		log.Printf("failed to parse request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse error"})
		return
	}

	for _, event := range cb.Events {
		if err := h.botService.HandleEvent(c, event); err != nil {
			log.Printf("failed to handle event: %v", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}