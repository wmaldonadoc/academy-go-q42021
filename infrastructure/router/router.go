package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
)

func NewRouter(c controller.AppController) {
	port := config.GetEnvVariable("PORT")

	router := gin.Default()

	mainGroup := router.Group("/api/v1")
	{
		pokemonRoutes := mainGroup.Group("/pokemons")
		{
			pokemonRoutes.GET("/:id", func(context *gin.Context) { c.Pokemon.GetById(context) })
		}

		healthRoutes := mainGroup.Group("/health")
		{
			healthRoutes.GET("/", func(context *gin.Context) { c.Health.GetServiceHealth(context) })
		}
	}

	router.Run(port)

}
