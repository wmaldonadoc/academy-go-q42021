package main

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/router"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func bootingGlobalLogger() *zap.Logger {
	// Replacing logger global
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger

}

func main() {
	loggerMgr := bootingGlobalLogger()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()

	logger.Debug("Booting routes...")
	router.NewRouter()
}
