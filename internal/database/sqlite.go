package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func InitSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./internal/database/url-shortener.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "internal/database/migrations",
	}

	_, err = migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	return db, err
}
