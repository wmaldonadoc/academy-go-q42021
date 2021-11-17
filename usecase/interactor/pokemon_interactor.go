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
	GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError)
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *ucExceptions.UseCaseError)
	GetRequest(url string) (*model.Pokemon, *ucExceptions.UseCaseError)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter, client vendors.HTTPClient) *pokemonInteractor {
	return &pokemonInteractor{r, p, client}
}

func (pi *pokemonInteractor) GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError) {
	p, err := pi.PokemonRepository.FindById(id)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error getting pokemon %s", err.Err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

func (pi *pokemonInteractor) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *ucExceptions.UseCaseError) {
	p, err := pi.PokemonRepository.CreateOne(pokemon)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error storing pokemon record %s", err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}
	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

func (pi *pokemonInteractor) GetRequest(url string) (*model.Pokemon, *ucExceptions.UseCaseError) {
	resp, err := pi.HTTPClient.Get(url)
	if resp.HTTPStatus != http.StatusOK {
		zap.S().Errorf("INTEREACTOR: Third part API response with an error status: %i", resp.HTTPStatus)
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

	return pi.PokemonPresenter.MappedPokemonFromAPI(resp), nil
}
