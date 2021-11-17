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
	GetByID(c Context)
	GetByName(c Context)
}

func NewPokemonController(pi interactor.PokemonInteractor) *pokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetByID(c Context) {
	requestId := c.Param("id")
	if pokemonId, err := strconv.Atoi(requestId); err == nil {
		p, err := pc.pokemonInteractor.GetByID(pokemonId)
		if err != nil {
			zap.S().Errorf("CONTROLLER: Error searching pokemon by id %s", err)
			genericException := exceptions.GenericException(
				err.Message,
				err.HTTPStatus,
				err.Code,
			)
			c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

			return
		}
		c.JSON(http.StatusOK, p)
	} else {
		zap.S().Errorf("CONTROLLER: The id should be a integer %s", requestId)
		parseError := exceptions.UnprocessableEntityException("The id should be a integer")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, parseError)

		return
	}
}

// TODO: Take the response from API and pase it to response
func (pc *pokemonController) GetByName(c Context) {
	pokemonName := c.Param("name")
	response, err := pc.pokemonInteractor.GetRequest("https://pokeapi.co/api/v2/pokemon/" + pokemonName)
	if err != nil {
		zap.S().Errorf("CONTROLLER: Error getting pokemon %s", pokemonName)
		genericException := exceptions.GenericException(
			err.Message,
			err.HTTPStatus,
			err.Code,
		)
		c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

		return
	}
	c.JSON(http.StatusOK, response)
}
