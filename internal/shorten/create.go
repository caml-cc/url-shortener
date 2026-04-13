package shorten

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"url-shortener/pkg/db"
	"url-shortener/pkg/utils"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func parseShortenPayload(payload string) (string, string, bool) {
	trimmed := strings.TrimSpace(payload)

	if strings.Count(trimmed, "|") == 0 {
		return randomString(8), trimmed, true
	} else if strings.Count(trimmed, "|") != 1 {
		return "", "", false
	}

	rawURL, alias, found := strings.Cut(trimmed, "|")
	if !found {
		return "", "", false
	}

	alias = strings.TrimSpace(alias)
	rawURL = strings.TrimSpace(rawURL)
	if alias == "" || rawURL == "" {
		return "", "", false
	}

	return alias, rawURL, true
}

func CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.Header.Get("K")
	if key != utils.Conf.API_KEY {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	alias, rawURL, ok := parseShortenPayload(string(body))
	if !ok {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	url, err := utils.NormalizeURL(rawURL)
	if err != nil {
		switch err {
		case utils.ErrUnsupportedURL:
			http.Error(w, "only http and https are allowed", http.StatusBadRequest)
		default:
			http.Error(w, "invalid URL", http.StatusBadRequest)
		}
		return
	}

	if err := db.AddURL(alias, url.String()); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(r.Host + "/" + alias + "\n"))
}

func randomString(n int) string {
	b := make([]byte, n)
	rand.NewSource(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
