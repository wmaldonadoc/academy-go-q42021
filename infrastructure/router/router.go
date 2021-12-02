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
			pokemonRoutes.GET("/id/:id", func(context *gin.Context) {
				resp := controller.Pokemon.GetByID(context)
				context.JSON(resp.HTTPStatus, resp.Data)
			})
			pokemonRoutes.GET("/name/:name", func(context *gin.Context) {
				resp := controller.Pokemon.GetByName(context)
				context.JSON(resp.HTTPStatus, resp.Data)
			})
			pokemonRoutes.GET("/filter", func(context *gin.Context) {
				resp := controller.Pokemon.FilterSearching(context)
				context.JSON(resp.HTTPStatus, resp.Data)
			})
		}

		healthRoutes := mainGroup.Group("/health")
		{
			healthRoutes.GET("/", func(context *gin.Context) {
				resp := controller.Health.GetServiceHealth(context)
				context.JSON(resp.HTTPStatus, resp.Data)
			})
		}
	}

	return router
}
