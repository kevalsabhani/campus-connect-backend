package server

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevalsabhani/campus-connect-backend/internal/config"
)

type Server struct {
	addr string
	db   *sql.DB
	cfg  *config.Config
}

func NewServer(addr string, db *sql.DB, cfg *config.Config) *Server {
	return &Server{
		addr: addr,
		db:   db,
		cfg:  cfg,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler: mapHandlers(),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("starting server", "port", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("failed to start server", "error", err)
		}
	}()

	<-done
	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	slog.Info("server shutdown successfully")
	return nil
}
