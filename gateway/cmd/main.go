package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

// @Title Hexago API Gateway
// @Version 0.1
// @Description API gateway for Hexago project.
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

func initialize(ctx context.Context, components *components) error {
	if err := components.mongoRepo.Migrate(ctx); err != nil {
		return fmt.Errorf("migrating MongoDB: %w", err)
	}
	return nil
}

func runAndWait(ctx context.Context, components *components) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	donationProducerFails := components.donationProducer.Run(ctx)
	logger.S().Infof("Run kafka donation producer")

	httpServerFails := components.httpServer.Run(ctx)
	logger.S().Infof("Run http server")

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-donationProducerFails:
		logger.S().Errorf("Error from donation producer: %v", err)
	case err := <-httpServerFails:
		logger.S().Errorf("Error from http server: %v", err)
	case sig := <-signals:
		logger.S().Infof("Caught signal[%s]", sig.String())
	}
}

func shutdown(ctx context.Context, components *components) {
	if err := components.httpServer.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown http server: %v", err)
	}

	if err := components.mongoRepo.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown MongoDB: %v", err)
	}

	if err := components.donationProducer.Shutdown(ctx); err != nil {
		logger.S().Errorf("Failed to shutdown kafka donation producer: %v", err)
	}
}
