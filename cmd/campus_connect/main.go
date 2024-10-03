package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kevalsabhani/campus-connect-backend/internal/config"
)

func main() {
	slog.Info("Campus Connect API", "version", "0.0.1")

	// Load config
	cfg := config.MustLoad()

	// Database setup

	// setup router
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	server := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("starting server", "host", cfg.Server.Host, "port", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("failed to start server", "error", err)
		}
	}()

	<-done
	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("server shutdown successfully")
}
