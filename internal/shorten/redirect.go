package shorten

import (
	"net/http"
	"url-shortener/pkg/db"

	"github.com/gorilla/mux"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	url, err := db.GetURL(alias)
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
