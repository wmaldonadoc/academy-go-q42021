package main

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/router"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Booting router...")
	router.NewRouter()
}
