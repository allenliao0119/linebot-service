package main

import (
	"log"

	"github.com/allenliao0119/linebot-service/internal/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}