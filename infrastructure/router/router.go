package router

import (
	"log"

	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewRouter - Setup & returns the API routes.
func NewRouter(controller controller.AppController) {
	port := config.GetEnvVariable("PORT")
	router := gin.Default()
	mainGroup := router.Group("/api/v1")
	{
		pokemonRoutes := mainGroup.Group("/pokemons")
		{
			pokemonRoutes.GET("/:id", func(context *gin.Context) { controller.Pokemon.GetByID(context) })
			pokemonRoutes.GET("/name/:name", func(context *gin.Context) { controller.Pokemon.GetByName(context) })
			pokemonRoutes.GET("/", func(context *gin.Context) { controller.Pokemon.BatchSearching(context) })
		}

		healthRoutes := mainGroup.Group("/health")
		{
			healthRoutes.GET("/", func(context *gin.Context) { controller.Health.GetServiceHealth(context) })
		}
	}
	if err := router.Run(port); err == nil {
		zap.S().Infof("API running at port", port)
	} else {
		log.Fatal("Error running router")
	}
}
