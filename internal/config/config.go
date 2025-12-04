package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server ServerConfig `envconfig:"SERVER"`
	LINE   LINEConfig   `envconfig:"LINE"`
	Chat   ChatConfig   `envconfig:"CHAT"`
	OpenAI OpenAIConfig `envconfig:"OPENAI"`
}

type ServerConfig struct {
	Port string `envconfig:"PORT" default:"8080"`
	Env  string `envconfig:"ENV" default:"development"` // development, production
}

type LINEConfig struct {
	ChannelSecret      string `envconfig:"CHANNEL_SECRET" required:"true"`
	ChannelAccessToken string `envconfig:"CHANNEL_ACCESS_TOKEN" required:"true"`
}

type ChatConfig struct {
	Mode string `envconfig:"MODE" default:"simple"` // simple, ai
}

type OpenAIConfig struct {
	APIKey string `envconfig:"API_KEY"`
	Model  string `envconfig:"MODEL" default:"gpt-4o-mini"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return  nil, fmt.Errorf("failed to load config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Validate() error {
	if cfg.Chat.Mode != "simple" && cfg.Chat.Mode != "ai" {
		return fmt.Errorf("invalid chat mode: %s (must be 'simple' or 'ai')", cfg.Chat.Mode)
	}

	if cfg.Chat.Mode == "ai" && cfg.OpenAI.APIKey == "" {
		return fmt.Errorf("OPENAI_API_KEY is required when CHAT_MODE is 'ai'")
	}

	return  nil
}

func (cfg *Config) IsProduction() bool {
	return cfg.Server.Env == "production"
}

func (cfg *Config) IsChatModeAI() bool {
	return cfg.Chat.Mode == "ai"
}