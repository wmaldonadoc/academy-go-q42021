package router

import (
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"

	"github.com/gin-gonic/gin"
)

// NewRouter - Setup & returns the API routes.
func NewRouter(controller controller.AppController) *gin.Engine {
	router := gin.Default()
	mainGroup := router.Group("/api/v1")
	{
		pokemonRoutes := mainGroup.Group("/pokemons")
		{
			pokemonRoutes.GET("/id/:id", func(context *gin.Context) { controller.Pokemon.GetByID(context) })
			pokemonRoutes.GET("/name/:name", func(context *gin.Context) { controller.Pokemon.GetByName(context) })
			pokemonRoutes.GET("/filter", func(context *gin.Context) { controller.Pokemon.FilterSearching(context) })
		}

		healthRoutes := mainGroup.Group("/health")
		{
			healthRoutes.GET("/", func(context *gin.Context) { controller.Health.GetServiceHealth(context) })
		}
	}

	return router
}
