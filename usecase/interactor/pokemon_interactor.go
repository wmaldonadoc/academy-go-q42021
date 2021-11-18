package interactor

import (
	"errors"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	ucExceptions "github.com/wmaldonadoc/academy-go-q42021/usecase/exceptions"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/repository"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/vendors"

	"go.uber.org/zap"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
	HTTPClient        vendors.HTTPClient
}

type PokemonInteractor interface {
	// GetById - Returns a pokemon given an ID.
	GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError)
	// CreateOne - Append the new pokemon row to CSV file
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *ucExceptions.UseCaseError)
	// GetPokemonByName - Get the pokemon from PokeAPI given the name and return it as Pokemon model.
	GetPokemonByName(name string) (*model.Pokemon, *ucExceptions.UseCaseError)
}

// NewPokemonInteractor - Receive a repository, presenter and HTTPClient and returns a concret instance of PokemonInteractor.
func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter, client vendors.HTTPClient) *pokemonInteractor {
	return &pokemonInteractor{r, p, client}
}

// GetById - Returns a pokemon given an ID.
func (pi *pokemonInteractor) GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError) {
	p, err := pi.PokemonRepository.FindById(id)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error getting pokemon %s", err.Err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

// CreateOne - Append the new pokemon row to CSV file.
func (pi *pokemonInteractor) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *ucExceptions.UseCaseError) {
	p, err := pi.PokemonRepository.CreateOne(pokemon)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error storing pokemon record %s", err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}
	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

// GetPokemonByName - Get the pokemon from PokeAPI given the name and return it as Pokemon model.
func (pi *pokemonInteractor) GetPokemonByName(name string) (*model.Pokemon, *ucExceptions.UseCaseError) {
	resp, err := pi.HTTPClient.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if resp.HTTPStatus != http.StatusOK {
		zap.S().Errorf("INTEREACTOR: Third part API response with an error status: ", resp.HTTPStatus)
		useCaseException := ucExceptions.NewErrorWrapper(
			constants.ThirdPartAPIExceptionCode,
			resp.HTTPStatus,
			errors.New("request error"),
			resp.Body,
		)
		return nil, &useCaseException
	}
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error making request %s", err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponseMappedPokemonFromAPI(resp), nil
}
