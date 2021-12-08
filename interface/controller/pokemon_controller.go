package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/interface/schemas"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"

	"go.uber.org/zap"
)

type ControllerResponse struct {
	HTTPStatus int         `json:"status"`
	Data       interface{} `json:"data"`
}

type ControllerPokemon struct {
	PokemonInteractor interactor.PokemonInteractor
}

// PokemonController - Holds the abstraction of controller methods.
type PokemonController interface {

	// GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
	// It will return an error as HTTP response if something goes wrong.
	GetByID(c Context) *ControllerResponse
	// GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
	// It will return an error as HTTP response if something goes wrong.
	// If the HTTP request fails nothing will gonna be stored.
	GetByName(c Context) *ControllerResponse
	// FilterSearching - Reads the CSV file using a worker pool and return the requested items.
	// It will return an error as HTTP response if something goes wrong.
	FilterSearching(c Context) *ControllerResponse
}

// NewPokemonController - Receive the controller interactor and returns a conrete instance of the controller.
func NewPokemonController(pi interactor.PokemonInteractor) *ControllerPokemon {
	return &ControllerPokemon{pi}
}

// GeyByID - Receive the HTTP request context, find the pokemon by ID and return it.
// It will return an error as HTTP response if something goes wrong.
func (pc *ControllerPokemon) GetByID(c Context) *ControllerResponse {
	requestError := pokerrors.DefaultError{}
	response := ControllerResponse{}
	var req schemas.GetPokemonById
	if err := c.BindUri(&req); err != nil {
		zap.S().Errorf("CONTROLLER: The id should be a integer %s", err)
		requestError.Code = constants.UnprocessableEntityExceptionCode
		requestError.HTTPStatus = http.StatusUnprocessableEntity
		requestError.Message = "The id should be a integer"

		response.HTTPStatus = requestError.HTTPStatus
		response.Data = requestError

		return &response
	}
	p, err := pc.PokemonInteractor.GetByID(req.ID)
	if err != nil {
		zap.S().Errorf("CONTROLLER: Error searching pokemon by id %s", err)
		requestError.Code = constants.NotFoundExceptionCode
		requestError.HTTPStatus = http.StatusNotFound
		requestError.Message = "Error searching pokemon by id"

		response.HTTPStatus = requestError.HTTPStatus
		response.Data = requestError

		return &response
	}

	response.HTTPStatus = http.StatusOK
	response.Data = p

	return &response
}

// GetByName - Receive the HTTP request context, it will request the pokemon name from an API then will stored it and finally it will return it as HTTP response.
// It will return an error as HTTP response if something goes wrong.
// If the HTTP request fails nothing will gonna be stored.
func (pc *ControllerPokemon) GetByName(c Context) *ControllerResponse {
	requestError := pokerrors.DefaultError{}
	response := ControllerResponse{}
	var req schemas.GetPokemonByName
	if err := c.BindUri(&req); err != nil {
		zap.S().Errorf("CONTROLLER: The name its mandatory %s", err)
		requestError.Code = constants.UnprocessableEntityExceptionCode
		requestError.HTTPStatus = http.StatusUnprocessableEntity
		requestError.Message = "The name its mandatory"

		response.HTTPStatus = requestError.HTTPStatus
		response.Data = requestError

		return &response

	}
	resp, err := pc.PokemonInteractor.GetPokemonByName(req.Name)
	if err != nil {
		zap.S().Errorf("CONTROLLER: Error getting pokemon %s", req.Name)
		requestError.Code = constants.DefaultExceptionCode
		requestError.HTTPStatus = http.StatusInternalServerError
		requestError.Message = "Error getting pokemon " + req.Name

		response.HTTPStatus = requestError.HTTPStatus
		response.Data = requestError

		return &response

	}
	record, repositoryError := pc.PokemonInteractor.CreateOne(resp)
	if repositoryError != nil {
		zap.S().Error("CONTROLLER: Error storing pokemon")
		requestError.Code = constants.DefaultExceptionCode
		requestError.HTTPStatus = repositoryError.HTTPStatus
		requestError.Message = "Error storing pokemon " + req.Name

		response.HTTPStatus = requestError.HTTPStatus
		response.Data = requestError

		return &response
	}

	response.HTTPStatus = http.StatusOK
	response.Data = record

	return &response
}

// FilterSearching - Reads the CSV file using a worker pool and return the requested items.
// It will return an error as HTTP response if something goes wrong.
func (pc *ControllerPokemon) FilterSearching(c Context) *ControllerResponse {
	requestError := pokerrors.DefaultError{}
	response := ControllerResponse{}
	var req schemas.BatchSearchingSchema
	if err := c.ShouldBindQuery(&req); err != nil {
		for _, field := range err.(validator.ValidationErrors) {
			zap.S().Debugf("CONTROLLER: Request error ", field.Error())
			message := fmt.Sprint("Missing query string param " + field.StructField() + " condition: " + field.ActualTag())
			requestError.Code = constants.UnprocessableEntityExceptionCode
			requestError.HTTPStatus = http.StatusUnprocessableEntity
			requestError.Message = message

			response.HTTPStatus = requestError.HTTPStatus
			response.Data = requestError

			return &response
		}
	}
	zap.S().Infof("Request: ", req.Items)
	zap.S().Infof("Request: ", req.ItemsPerWorker)
	zap.S().Infof("Request: ", req.Type)
	resp := pc.PokemonInteractor.BatchFilter(req.Type, req.Items, req.ItemsPerWorker)
	response.HTTPStatus = http.StatusOK
	response.Data = resp

	return &response
}
