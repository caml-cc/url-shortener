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

	server.HandleFunc("/shorten", shorten.ShortenURL)
	server.HandleFunc("/{alias}", shorten.RedirectURL)

	fmt.Printf("Server running on port: %s", conf.PORT)
	switch conf.ENV {
	case "production":
		http.ListenAndServeTLS(":"+conf.PORT, conf.CRT, conf.KEY, server)
	case "staging":
		log.Fatal(http.ListenAndServeTLS(":"+conf.PORT, conf.CRT, conf.KEY, server))
	default:
		log.Fatal(http.ListenAndServe(":"+conf.PORT, server))
	}
}
