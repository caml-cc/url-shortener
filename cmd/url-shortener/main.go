package main

import (
	"os"
	"url-shortener/api"
	"url-shortener/internal/database"
	"url-shortener/pkg/utils"
)

func main() {
	err := utils.LoadConfig()
	if err != nil {
		println("ERROR: " + err.Error())
		os.Exit(1)
	}

	err = database.InitSQLiteDB()
	if err != nil {
		println(err)
		os.Exit(1)
	}

	api.StartServer(utils.Conf)

	database.DbClose()
}
