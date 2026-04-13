package shorten

import (
	"io"
	"net/http"
	"strings"
	"url-shortener/pkg/db"
	"url-shortener/pkg/utils"
)

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.Header.Get("K")
	if key != utils.Conf.API_KEY {
		http.Error(w, "401 unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "400 Bad request", http.StatusBadRequest)
		return
	}

	alias := strings.TrimSpace(string(body))

	if err := db.DeleteURL(alias); err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(alias + " deleted\n"))
}
