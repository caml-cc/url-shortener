package db

import "url-shortener/internal/database"

func AddURL(alias string, url string) error {
	_, err := database.DB.Exec("INSERT INTO ALIAS (alias, url) VALUES ($1, $2);", alias, url)
	return err
}

func GetURL(alias string) (string, error) {
	var url string
	err := database.DB.QueryRow("SELECT url FROM ALIAS WHERE alias = ?;", alias).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
