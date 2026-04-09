package utils

import (
	"os"
	"url-shortener/internal/models"

	"github.com/joho/godotenv"
)

var Conf models.Config

func LoadConfig() {
	godotenv.Load()

	Conf = models.Config{
		ENV:     fallback(os.Getenv("ENV"), "development"),
		CRT:     os.Getenv("CRT"),
		KEY:     os.Getenv("KEY"),
		PORT:    fallback(os.Getenv("PORT"), "5099"),
		API_KEY: os.Getenv("API_KEY"),
	}
}

func fallback(value, def string) string {
	if value == "" {
		return def
	}
	return value
}
