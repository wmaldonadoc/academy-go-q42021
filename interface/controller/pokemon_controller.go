package controller

import (
	"net/http"
	"strconv"

	"github.com/wmaldonadoc/academy-go-q42021/interface/exceptions"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"

	"go.uber.org/zap"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetById(c Context)
}

func NewPokemonController(pi interactor.PokemonInteractor) *pokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetById(c Context) {
	requestId := c.Param("id")
	if pokemonId, err := strconv.Atoi(requestId); err == nil {
		zap.S().Infof("Searching Pokemon with id %s", pokemonId)
		p, err := pc.pokemonInteractor.GetByID(pokemonId)
		zap.S().Infof("Pokemon given %s", p)
		zap.S().Infof("Error %s", err)
		if err != nil {
			zap.S().Errorf("Error searching pokemon by id %s", err)
			genericException := exceptions.GenericException(
				err.Message,
				err.HTTPStatus,
				err.Code,
			)
			c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

			return
		}
		if p == nil {
			zap.S().Errorf("Pokemon not found with id %s", requestId)
			notFound := exceptions.PokemonNotFoundException()
			c.AbortWithStatusJSON(http.StatusNotFound, notFound)

			return
		}
		c.JSON(http.StatusOK, p)
	} else {
		zap.S().Errorf("The id should be a integer %s", requestId)
		parseError := exceptions.ParseTypesException("string", "int")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, parseError)

		return
	}
}
