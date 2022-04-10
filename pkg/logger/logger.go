package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(debug bool) {
	var cfg zap.Config
	if debug {
		cfg = devLoggerConfig()
	} else {
		cfg = prodLoggerConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Errorf("building logger: %w", err))
	}
	zap.ReplaceGlobals(logger)
}

func devLoggerConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.FunctionKey = zapcore.OmitKey
	return cfg
}

func prodLoggerConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.FunctionKey = "func"
	return cfg
}

func S() *zap.SugaredLogger {
	return zap.S()
}

func L() *zap.Logger {
	return zap.L()
}
