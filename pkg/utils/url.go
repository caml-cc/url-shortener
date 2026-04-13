package utils

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrInvalidURL     = errors.New("invalid url")
	ErrUnsupportedURL = errors.New("unsupported url scheme")
)

func NormalizeURL(input string) (*url.URL, error) {
	input = strings.TrimSpace(input)

	// If no scheme, assume http
	if !strings.Contains(input, "://") {
		input = "http://" + input
	}

	parsed, err := url.Parse(input)
	if err != nil {
		return nil, ErrInvalidURL
	}

	// Must have host
	if parsed.Host == "" {
		return nil, ErrInvalidURL
	}

	parsed.Host = strings.ToLower(parsed.Host)

	switch parsed.Scheme {
	case "http", "https":
		return parsed, nil
	default:
		return nil, ErrUnsupportedURL
	}
}
