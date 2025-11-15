package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sariya23/manage_pr_service/internal/app/app"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	config := cfg.MustLoad()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	application := app.NewApp(ctx, logger, config)
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		sig := <-exit
		logger.Info("get signal", slog.String("signal", sig.String()))
		cancel()
	}()
	var wg sync.WaitGroup
	wg.Go(func() {
		application.MustStart()
	})
	wg.Go(func() {
		<-ctx.Done()
		application.GracefulStop(context.Background())
	})
	wg.Wait()
	logger.Info("application stopped")
}
