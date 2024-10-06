package db

import (
	"database/sql"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
)

func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to establish connection to Postgres: %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping Postgres: %s", err.Error())
	}
	slog.Info("Connected to Postgres database")
	return db
}
