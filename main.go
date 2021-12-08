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

func bootingGlobalLogger() (*zap.Logger, error) {
	// Replacing logger global
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func main() {

	// Setting zap as global logger
	loggerMgr, zapError := bootingGlobalLogger()

	if zapError != nil {
		zap.S().Errorf("Zap bootstrap error", zapError)
	}

	zap.ReplaceGlobals(loggerMgr)
	defer func() {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatalf("Error setting logger %s", err)
		}
	}()
	logger := loggerMgr.Sugar()

	// Bootsraping CSV connection
	var openFunc = os.Open
	filePath, fpError := config.GetEnvVariable("FILE_LOCATION")
	if fpError != nil {
		zap.S().Errorf("Error getting env key", fpError)
	}
	db, err := datastore.NewCSV(filePath, openFunc)
	if err != nil {
		zap.S().Error("Error bootstraping CSV file")
	}
	r := registry.NewRegistry(db)

	// Running Router
	logger.Debug("Booting routes...")
	rtr := router.NewRouter(r.NewAppController())

	port, envError := config.GetEnvVariable("PORT")
	if envError != nil {
		zap.S().Errorf("Error getting env key", envError)
	}
	if err := rtr.Run(port); err != nil {
		zap.S().Error("Error bootstraping router", err)
	}

	zap.S().Infof("Route booted successfully")
}
