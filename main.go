package main

import (
	"log"

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
	defer func() {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatalf("Error setting logger %s", err)
		}
	}()
	logger := loggerMgr.Sugar()

	db := datastore.NewCSV()
	r := registry.NewRegistry(db)

	logger.Debug("Booting routes...")
	router.NewRouter(r.NewAppController())
}
