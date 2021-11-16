package router

import (
	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.AppController) {
	port := config.GetEnvVariable("PORT")
	router := gin.Default()
	mainGroup := router.Group("/api/v1")
	{
		pokemonRoutes := mainGroup.Group("/pokemons")
		{
			pokemonRoutes.GET("/:id", func(context *gin.Context) { controller.Pokemon.GetByID(context) })
		}

		healthRoutes := mainGroup.Group("/health")
		{
			healthRoutes.GET("/", func(context *gin.Context) { controller.Health.GetServiceHealth(context) })
		}
	}
	router.Run(port)
}
