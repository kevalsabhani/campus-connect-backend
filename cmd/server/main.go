package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: buildHandler(),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("starting server", "port", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("failed to start server", "error", err)
		}
	}()

	<-done
	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("server shutdown successfully")
}

func buildHandler() *mux.Router {
	// setup router
	r := mux.NewRouter()

	v1Router := r.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	return r
}
