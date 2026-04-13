package api

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/shorten"

	"github.com/gorilla/mux"
)

func StartServer(conf models.Config) {
	server := mux.NewRouter()

	server.HandleFunc("/create", shorten.CreateURL)
	server.HandleFunc("/delete", shorten.DeleteURL)
	server.HandleFunc("/{alias}", shorten.RedirectURL)

	fmt.Printf("Server running on port: %s", conf.PORT)
	log.Fatal(http.ListenAndServe(":"+conf.PORT, server))
}
