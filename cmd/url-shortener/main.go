package urlshortener

import (
	"url-shortener/api"
	"url-shortener/internal/database"
	"url-shortener/pkg/utils"
)

func main() {
	DB, err := database.InitSQLiteDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	conf := utils.LoadConfig()

	api.StartServer(conf)
}
