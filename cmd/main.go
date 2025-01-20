package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/nathaliaguayos/msg-receiver/config"
	"github.com/nathaliaguayos/msg-receiver/internal/handlers"
	"github.com/nathaliaguayos/msg-receiver/internal/rest"
	"github.com/nathaliaguayos/msg-receiver/internal/services"
	"github.com/nathaliaguayos/msg-receiver/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	// initialize application
	cfg := config.New()
	log := logger.New(cfg)
	log.Info().Msg("msg-receiver loading")

	jwtService := services.NewJWTService(cfg.SecretKey, cfg.Issuer)
	jwtHandler := handlers.NewJWTHandler(jwtService)

	restClient, err := rest.NewRestClient(log, jwtHandler, cfg.RateLimit)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating rest service")
	}

	log.Info().
		Str("host", cfg.Host).
		Uint("port", cfg.Port).
		Msg("Starting HTTP listener")

	port := strconv.FormatUint(uint64(cfg.Port), 10)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: restClient.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msg(fmt.Sprintf("error connecting the server: %v", err))
		}
	}()

	//Init shutting down gracefully
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Info().Msg("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to server shutdown due: %v", err))
	}
	//catching ctx.Done(). Timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info().Msg("timeout of 5 seconds")
	}
	log.Info().Msg("the server has been turned off gracefully")
}
