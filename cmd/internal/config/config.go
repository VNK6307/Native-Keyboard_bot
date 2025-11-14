package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	BotURL   string
}

func LoadConfig() (*Config, error) {

	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("BOT_TOKEN environment variable not set")
	}

	baseUrl := os.Getenv("BOT_URL")
	if baseUrl == "" {
		return nil, errors.New("BOT_URL environment variable not set")
	}

	return &Config{
			token,
			baseUrl + token,
		},
		nil
}
