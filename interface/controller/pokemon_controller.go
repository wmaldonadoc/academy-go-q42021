package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wmaldonadoc/academy-go-q42021/interface/schemas"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"

	"go.uber.org/zap"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

// PokemonController - Holds the abstraction of controller methods.
type PokemonController interface {

	// GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
	// It will return an error as HTTP response if something goes wrong.
	GetByID(c Context)
	// GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
	// It will return an error as HTTP response if something goes wrong.
	// If the HTTP request fails nothing will gonna be stored.
	GetByName(c Context)
	// FilterSearching - Reads the CSV file using a worker pool and return the requested items.
	// It will return an error as HTTP response if something goes wrong.
	FilterSearching(c Context)
}

// NewPokemonController - Receive the controller interactor and returns a conrete instance of the controller.
func NewPokemonController(pi interactor.PokemonInteractor) *pokemonController {
	return &pokemonController{pi}
}

// GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
// It will return an error as HTTP response if something goes wrong.
func (pc *pokemonController) GetByID(c Context) {
	var req schemas.GetPokemonById
	if err := c.BindUri(&req); err != nil {
		zap.S().Errorf("CONTROLLER: The id should be a integer %s", err)
		parseError := pokerrors.GenerateUnprocessableEntityError("The id should be a integer")
		c.AbortWithStatusJSON(parseError.HTTPStatus, parseError)

		return
	}
	p, err := pc.pokemonInteractor.GetByID(req.ID)
	if err != nil {
		zap.S().Errorf("CONTROLLER: Error searching pokemon by id %s", err)
		genericException := pokerrors.GenerateNotFoundError("Error searching pokemon by id")
		c.AbortWithStatusJSON(genericException.HTTPStatus, &genericException)

		return
	}
	c.JSON(http.StatusOK, p)
}

// GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
// It will return an error as HTTP response if something goes wrong.
// If the HTTP request fails nothing will gonna be stored.
func (pc *pokemonController) GetByName(c Context) {
	var req schemas.GetPokemonByName
	if err := c.BindUri(&req); err != nil {
		zap.S().Errorf("CONTROLLER: The name its mandatory %s", err)
		parseError := pokerrors.GenerateUnprocessableEntityError("The name its mandatory")
		c.AbortWithStatusJSON(parseError.HTTPStatus, parseError)

		return
	}
	response, err := pc.pokemonInteractor.GetPokemonByName(req.Name)
	if err != nil {
		zap.S().Errorf("CONTROLLER: Error getting pokemon %s", req.Name)
		genericException := pokerrors.GenerateDefaultError("Error getting pokemon " + req.Name)
		c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

		return
	}
	record, repositoryError := pc.pokemonInteractor.CreateOne(response)
	if repositoryError != nil {
		zap.S().Error("CONTROLLER: Error storing pokemon")
		genericException := pokerrors.GenerateDefaultError("Error storing pokemon " + req.Name)
		c.AbortWithStatusJSON(genericException.HTTPStatus, genericException)

		return
	}

	c.JSON(http.StatusOK, record)
}

// FilterSearching - Reads the CSV file using a worker pool and return the requested items.
// It will return an error as HTTP response if something goes wrong.
func (pc *pokemonController) FilterSearching(c Context) {
	var req schemas.BatchSearchingSchema
	if err := c.ShouldBindQuery(&req); err != nil {
		for _, field := range err.(validator.ValidationErrors) {
			zap.S().Debugf("CONTROLLER: Request error ", field.Error())
			message := fmt.Sprint("Missing query string param " + field.StructField() + " condition: " + field.ActualTag())
			requestError := pokerrors.GenerateUnprocessableEntityError(message)
			c.AbortWithStatusJSON(requestError.HTTPStatus, requestError)
			return
		}
	}
	zap.S().Infof("Request: ", req.Items)
	zap.S().Infof("Request: ", req.ItemsPerWorker)
	zap.S().Infof("Request: ", req.Type)
	resp := pc.pokemonInteractor.BatchFilter(req.Type, req.Items, req.ItemsPerWorker)
	c.JSON(http.StatusOK, resp)
}
