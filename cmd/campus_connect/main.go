package main

import (
	"log/slog"

	"github.com/kevalsabhani/campus-connect-backend/internal/config"
)

func main() {
	slog.Info("Campus Connect API", "version", "0.0.1")

	// Load config
	cfg := config.MustLoad()
	slog.Info("Loaded config", "config", cfg)

	// Database setup

	// setup router

	// Run server
}
