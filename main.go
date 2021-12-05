package main

import (
	"log"
	"os"

	"github.com/wmaldonadoc/academy-go-q42021/config"
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
	port := config.GetEnvVariable("PORT")
	loggerMgr := bootingGlobalLogger()
	zap.ReplaceGlobals(loggerMgr)
	defer func() {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatalf("Error setting logger %s", err)
		}
	}()
	logger := loggerMgr.Sugar()
	var openFunc = os.Open
	filePath := config.GetEnvVariable("FILE_LOCATION")
	db, err := datastore.NewCSV(filePath, openFunc)
	if err != nil {
		zap.S().Error("Error bootstraping CSV file")
	}
	r := registry.NewRegistry(db)

	logger.Debug("Booting routes...")
	rtr := router.NewRouter(r.NewAppController())

	if err := rtr.Run(port); err != nil {
		zap.S().Error("Error bootstraping router", err)
	}

	zap.S().Infof("Route booted successfully")
}
