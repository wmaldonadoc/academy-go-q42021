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

// PokemonController - Holds the abstraction of controller methods.
type PokemonController interface {
	/*
		GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
		It will return an error as HTTP response if something goes wrong.
	*/
	GetByID(c Context)
	/*
		GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
		It will return an error as HTTP response if something goes wrong.
		If the HTTP request fails nothing will gonna be stored.
	*/
	GetByName(c Context)
}

// NewPokemonController - Receive the controller interactor and returns a concret instance of the controller.
func NewPokemonController(pi interactor.PokemonInteractor) *pokemonController {
	return &pokemonController{pi}
}

/*
	GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
	It will return an error as HTTP response if something goes wrong.
*/
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

/*
	GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
	It will return an error as HTTP response if something goes wrong.
	If the HTTP request fails nothing will gonna be stored.
*/
func (pc *pokemonController) GetByName(c Context) {
	pokemonName := c.Param("name")
	response, err := pc.pokemonInteractor.GetPokemonByName(pokemonName)
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
	record, repositoryError := pc.pokemonInteractor.CreateOne(response)
	if repositoryError != nil {
		zap.S().Error("CONTROLLER: Error storing pokemon")
		genericException := exceptions.GenericException(
			repositoryError.Message,
			repositoryError.HTTPStatus,
			repositoryError.Code,
		)
		c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

		return
	}

	c.JSON(http.StatusOK, record)
}
