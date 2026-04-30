package utils

import (
	"errors"
	"os"
	"strconv"
	"url-shortener/internal/models"

	"github.com/joho/godotenv"
)

var Conf models.Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	temp_num, err := strconv.Atoi(fallback(os.Getenv("RAND_CHARS"), "8"))
	if err != nil || temp_num <= 0 {
		return errors.New("RAND_CHARS value must be a positive integer")
	}

	Conf = models.Config{
		PORT:       os.Getenv("PORT"),
		API_KEY:    os.Getenv("API_KEY"),
		RAND_CHARS: temp_num,
	}

	if Conf.API_KEY == "" || Conf.PORT == "" {
		return errors.New("env file is not filled in correctly")
	}

	return nil
}

func fallback(value, def string) string {
	if value == "" {
		return def
	}
	return value
}
