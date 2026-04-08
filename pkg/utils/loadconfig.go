package utils

import (
	"os"
	"url-shortener/internal/models"

	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {
	godotenv.Load()

	return models.Config{
		ENV:  fallback(os.Getenv("ENV"), "development"),
		CRT:  os.Getenv("CRT"),
		KEY:  os.Getenv("KEY"),
		PORT: fallback(os.Getenv("PORT"), "5098"),
	}
}

func fallback(value, def string) string {
	if value == "" {
		return def
	}
	return value
}
