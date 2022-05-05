package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	appCtx := context.Background()
	runAndWait(appCtx, components)

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

func runAndWait(ctx context.Context, components *Components) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	donationConsumerFails := components.donationConsumer.Run(ctx)
	logger.S().Info("Run kafka donation consumer")

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-donationConsumerFails:
		logger.S().Errorf("Error from donation consumer: %v", err)
	case sig := <-signals:
		logger.S().Infof("Caught signal[%s]", sig.String())
	}
}

func shutdown(ctx context.Context, components *Components) {
	if err := components.donationConsumer.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown kafka donation consumer: %v", err)
	}

	if err := components.mongoRepo.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown MongoDB: %v", err)
	}
}
