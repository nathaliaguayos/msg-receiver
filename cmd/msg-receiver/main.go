package main

import (
	"context"
	"github.com/nathaliaguayos/msg-receiver/internal/config"
	"github.com/nathaliaguayos/msg-receiver/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// initialize application
	cfg := config.New()
	log := logger.New(cfg)
	log.Info().Msg("msg-receiver loading")

	// Create a context that will be canceled on shutdown signal
	_, cancel := context.WithCancel(context.Background())

	// Channel to listen for shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle shutdown signals
	go func() {
		<-sigChan
		log.Info().Msg("shutdown signal received, shutting down gracefully...")
		cancel()
	}()
}
