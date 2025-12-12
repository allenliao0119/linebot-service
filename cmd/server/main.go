package main

import (
	"log"

	"github.com/allenliao0119/linebot-service/internal/bot"
	"github.com/allenliao0119/linebot-service/internal/config"
	"github.com/allenliao0119/linebot-service/internal/handler"
	"github.com/allenliao0119/linebot-service/internal/router"
	"github.com/allenliao0119/linebot-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	var chatService service.ChatService
	if cfg.IsChatModeAI() {
		chatService = bot.NewAIChatBot(cfg.OpenAI.APIKey, cfg.OpenAI.Model)
		log.Println("using AI chat mode with model:", cfg.OpenAI.Model)
	} else {
		chatService = bot.NewSimpleChatBot()
		log.Println("using simple chat mode")
	}

	botService, err := service.NewLineBotService(
		cfg.LINE.ChannelAccessToken, 
		chatService,
	)
	if err != nil {
		log.Fatalf("failed to create bot service: %v", err)
	}

	healthHandler := handler.NewHealthHandler()
	webhookHandler := handler.NewWebHookHandler(cfg.LINE.ChannelSecret, botService)

	r := router.NewRouter(healthHandler, webhookHandler)

	log.Printf("server starting on port %s (env: %s)", cfg.Server.Port, cfg.Server.Env)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}