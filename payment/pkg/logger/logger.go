package logger

import (
	"fmt"

	"github.com/isutare412/hexago/payment/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	fastLogger   *zap.Logger
	slowLogger   *zap.SugaredLogger
	accessLogger *zap.Logger
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

	accLogger, err := accessLoggerConfig(zCfg).Build()
	if err != nil {
		panic(fmt.Errorf("building access logger: %w", err))
	}

	fastLogger = logger
	slowLogger = logger.Sugar()
	accessLogger = accLogger
}

func A() *zap.Logger {
	return accessLogger
}

func S() *zap.SugaredLogger {
	return slowLogger
}

func L() *zap.Logger {
	return fastLogger
}

func Sync() {
	fastLogger.Sync()
	slowLogger.Sync()
	accessLogger.Sync()
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

func accessLoggerConfig(cfg zap.Config) zap.Config {
	cfg.EncoderConfig.CallerKey = zapcore.OmitKey
	cfg.EncoderConfig.FunctionKey = zapcore.OmitKey
	return cfg
}
