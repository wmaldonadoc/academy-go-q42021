package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/config"
	"go.uber.org/zap"
)

func NewRouter() {
	host := config.GetEnvVariable("BASE_HOST")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		zap.S().Info("Testing global")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(host)

}
