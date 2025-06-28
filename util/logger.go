package util

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func Init(isProduction bool) error {
	// Log level
	level := zapcore.InfoLevel
	if !isProduction {
		level = zapcore.DebugLevel
	}

	// Create a file to write logs
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	// Encoder configuration (pretty for dev, json for prod)
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if isProduction {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create core for file output
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(logFile),
		level,
	)

	// Optional: also output to console
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	)

	// Combine both outputs
	core := zapcore.NewTee(fileCore, consoleCore)

	// Build logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log = logger.Sugar()

	return nil
}
