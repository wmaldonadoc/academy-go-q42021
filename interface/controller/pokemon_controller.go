package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
	"go.uber.org/zap"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetById(c Context)
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetById(c Context) {
	requestId := c.Param("id")
	if pokemonId, err := strconv.Atoi(requestId); err == nil {
		zap.S().Infof("Searching Pokemon with id %s", pokemonId)
		p, err := pc.pokemonInteractor.GetById(pokemonId)
		if err != nil {
			zap.S().Errorf("Error seaching pokemon by id %s", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		if p == nil {
			zap.S().Errorf("Pokemon not found with id %s", requestId)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "NotFound"})
			return
		}
		c.JSON(http.StatusOK, p)
	}

}
