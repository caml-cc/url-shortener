package main

import (
	"url-shortener/api"
	"url-shortener/internal/database"
	"url-shortener/pkg/utils"
)

func main() {
	utils.LoadConfig()
	err := database.InitSQLiteDB()
	if err != nil {
		panic(err)
	}
	defer database.DbClose()

	api.StartServer(utils.Conf)
}
