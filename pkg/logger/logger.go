package logger

import (
	"fmt"

	"github.com/isutare412/hexago/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(cfg *config.LoggerConfig) {
	var zCfg zap.Config
	if cfg.Format == config.LogFormatJson {
		zCfg = jsonLoggerConfig()
	} else {
		zCfg = textLoggerConfig()
	}
	zCfg.DisableStacktrace = !cfg.StackTrace

	logger, err := zCfg.Build()
	if err != nil {
		panic(fmt.Errorf("building logger: %w", err))
	}
	zap.ReplaceGlobals(logger)
}

func textLoggerConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.FunctionKey = zapcore.OmitKey
	return cfg
}

func jsonLoggerConfig() zap.Config {
	cfg := zap.NewProductionConfig()
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
