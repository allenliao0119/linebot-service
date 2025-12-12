package router

import (
	"github.com/allenliao0119/linebot-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	healthHandler *handler.HealthHandler, 
	webhookHandler *handler.WebHookHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthHandler.Check)
	r.POST("/webhook", webhookHandler.Handle)

	return r
}