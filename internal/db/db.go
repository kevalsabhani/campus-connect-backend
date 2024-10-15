package db

import (
	"database/sql"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
)

// InitDB initializes a connection to a Postgres database. The connection
// string should be in the form of a Postgres DSN. The function will log a
// fatal error and exit if the connection can not be established or if the
// database can not be pinged.
func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to establish connection to Postgres: %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping Postgres: %s", err.Error())
	}
	slog.Info("connected to postgres database")
	return db
}
