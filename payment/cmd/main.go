package main

import (
	"context"
	"flag"
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
	_, err = dependencyInjection(startupCtx, cfg)
	if err != nil {
		logger.S().Fatalf("Failed to inject dependencies: %v", err)
	}
	logger.S().Info("Done dependency injection")
}
