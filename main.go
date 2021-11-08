package main

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/datastore"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/router"
	"github.com/wmaldonadoc/academy-go-q42021/registry"
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
	defer loggerMgr.Sync()
	logger := loggerMgr.Sugar()

	datastore.NewDb()
	r := registry.NewRegistry()

	logger.Debug("Booting routes...")
	router.NewRouter(r.NewAppController())
}
