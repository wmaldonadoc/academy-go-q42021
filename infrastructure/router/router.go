package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
)

func NewRouter(c controller.AppController) {
	host := config.GetEnvVariable("BASE_HOST")

	router := gin.Default()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())

	router.GET("/api/pokemon/:id", func(context *gin.Context) { c.Pokemon.GetById(context) })

	router.Run(host)

}
