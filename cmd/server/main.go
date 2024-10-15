package main

import (
	"log/slog"

	"github.com/kevalsabhani/campus-connect-backend/internal/config"
	"github.com/kevalsabhani/campus-connect-backend/internal/db"
)

var Version = "1.0.0"

func main() {
	slog.Info("Campus Connect API", "version", Version)

	// Load config
	cfg := config.MustLoad()

	// Database setup
	db := db.InitDB(cfg.Database.Dsn)
	defer db.Close()
}
