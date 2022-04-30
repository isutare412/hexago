package main

import (
	"context"
	"flag"
	"time"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

func main() {
	flag.Parse()
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}
	logger.Initialize(cfg.Logger)
	defer logger.S().Sync()

	logger.S().Info("Start dependency injection")
	diCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	beans, err := dependencyInjection(diCtx, cfg)
	if err != nil {
		logger.S().Fatalf("Failed to inject dependencies: %v", err)
	}
	logger.S().Info("Done dependency injection")

	logger.S().Info("Start graceful shutdown")
	ctx := context.Background()
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	shutdown(shutdownCtx, beans)
	logger.S().Info("Done graceful shutdown")
}

func shutdown(ctx context.Context, beans *beans) {
	if err := beans.mongoRepo.Close(ctx); err != nil {
		logger.S().Error("Failed to close mongodb: %v", err)
	}
}
