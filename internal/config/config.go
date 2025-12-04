package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerEnv string

const (
	ServerEnvDevelopment ServerEnv = "development"
	ServerEnvProduction  ServerEnv = "production"
)

type ChatMode string

const (
	ChatModeSimple ChatMode = "simple"
	ChatModeAI     ChatMode = "ai"
)

type Config struct {
	Server ServerConfig `envconfig:"SERVER"`
	LINE   LINEConfig   `envconfig:"LINE"`
	Chat   ChatConfig   `envconfig:"CHAT"`
	OpenAI OpenAIConfig `envconfig:"OPENAI"`
}

type ServerConfig struct {
	Port string    `envconfig:"PORT" default:"8080"`
	Env  ServerEnv `envconfig:"ENV" default:"development"`
}

type LINEConfig struct {
	ChannelSecret      string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
	ChannelAccessToken string `envconfig:"LINE_CHANNEL_ACCESS_TOKEN" required:"true"`
}

type ChatConfig struct {
	Mode ChatMode `envconfig:"MODE" default:"simple"`
}

type OpenAIConfig struct {
	APIKey string `envconfig:"API_KEY"`
	Model  string `envconfig:"MODEL" default:"gpt-4o-mini"`
}

func Load() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not found, using environment variables only")
	}

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
	// Validate server environment
	switch cfg.Server.Env {
	case ServerEnvDevelopment, ServerEnvProduction:
		// valid
	default:
		return fmt.Errorf("invalid server environment: %s (must be 'development' or 'production')", cfg.Server.Env)
	}

	// Validate chat mode
	switch cfg.Chat.Mode {
	case ChatModeSimple:
		// no additional validation needed
	case ChatModeAI:
		if cfg.OpenAI.APIKey == "" {
			return fmt.Errorf("openai api key is required when chat mode is 'ai'")
		}
	default:
		return fmt.Errorf("invalid chat mode: %s (must be 'simple' or 'ai')", cfg.Chat.Mode)
	}

	return nil
}

func (cfg *Config) IsProduction() bool {
	return cfg.Server.Env == ServerEnvProduction
}

func (cfg *Config) IsChatModeAI() bool {
	return cfg.Chat.Mode == ChatModeAI
}