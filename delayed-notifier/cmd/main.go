package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/app"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/config"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	zlog.Init()

	cfg := config.New()

	a, err := app.New(cfg)
	if err != nil {
		zlog.Logger.Error().
			Err(err).
			Msg("failed to create app")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errChan := make(chan error)
	go a.Start(ctx, errChan)

	select {
	case err := <-errChan:
		zlog.Logger.Error().
			Err(err).
			Msg("failed to start server")
		os.Exit(1)
	case <-ctx.Done():
		if err := a.Stop(); err != nil {
			zlog.Logger.Error().
				Err(err).
				Msg("failed to stop server")
			os.Exit(1)
		}
	}
}
