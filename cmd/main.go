package main

import (
	"context"
	"madastore/analytics/internal/app"
	"madastore/analytics/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize app")
	}

	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdownCh
		log.Info().Msg("shutting down")

		appCancel()
		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		_ = application.Shutdown(ctx)
	}()

	if err := application.Run(appCtx); err != nil {
		log.Fatal().Err(err).Msg("server exited")
	}
}
