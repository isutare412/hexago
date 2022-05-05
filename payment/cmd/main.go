package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/isutare412/hexago/payment/pkg/config"
	"github.com/isutare412/hexago/payment/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

func main() {
	flag.Parse()
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}
	logger.Initialize(cfg.Logger)
	defer logger.Sync()

	startupCtx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.Timeout.Startup)*time.Second)
	defer cancel()

	logger.S().Info("Start dependency injection")
	components, err := dependencyInjection(startupCtx, cfg)
	if err != nil {
		logger.S().Fatalf("Failed to inject dependencies: %v", err)
	}
	logger.S().Info("Done dependency injection")

	logger.S().Info("Start initialization")
	if err := initialize(startupCtx, components); err != nil {
		logger.S().Fatalf("Failed to initialize components: %v", err)
	}
	logger.S().Info("Done initialization")

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.Timeout.Shutdown)*time.Second)
	defer cancel()

	logger.S().Info("Start graceful shutdown")
	shutdown(shutdownCtx, components)
	logger.S().Info("Done graceful shutdown")
}

func initialize(ctx context.Context, components *Components) error {
	if err := components.mongoRepo.Migrate(ctx); err != nil {
		return fmt.Errorf("migrating MongoDB: %w", err)
	}
	return nil
}

func shutdown(ctx context.Context, components *Components) {
	if err := components.mongoRepo.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown MongoDB: %v", err)
	}
}
