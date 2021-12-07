package interactor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/repository"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/vendors"
	"github.com/wmaldonadoc/academy-go-q42021/workers"
	"github.com/wmaldonadoc/academy-go-q42021/workers/pool"

	"go.uber.org/zap"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
	HTTPClient        vendors.HTTPClient
	WorkerPool        workers.Dispatcher
}

type PokemonInteractor interface {
	// GetById - Returns a pokemon given an ID.
	GetByID(id int) (*model.Pokemon, *pokerrors.UseCaseError)
	// CreateOne - Append the new pokemon row to CSV file
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *pokerrors.UseCaseError)
	// GetPokemonByName - Get the pokemon from PokeAPI given the name and return it as Pokemon model.
	GetPokemonByName(name string) (*model.Pokemon, *pokerrors.UseCaseError)
	// BatchFilter - Create the worker pool and dispatch the jobs to recover the items from CSV.
	BatchFilter(disc string, items int, itemsPerworker int) []*model.Pokemon
}

// NewPokemonInteractor - Receive a repository, presenter and HTTPClient and returns a conrete instance of PokemonInteractor.
func NewPokemonInteractor(
	r repository.PokemonRepository,
	p presenter.PokemonPresenter,
	client vendors.HTTPClient,
	disp workers.Dispatcher,
) *pokemonInteractor {
	return &pokemonInteractor{r, p, client, disp}
}

// GetById - Returns a pokemon given an ID.
func (pi *pokemonInteractor) GetByID(id int) (*model.Pokemon, *pokerrors.UseCaseError) {
	p, err := pi.PokemonRepository.FindByID(id)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error getting pokemon %s", err.Err)
		useCaseException := pokerrors.GenerateUseCaseError("Error getting pokemon")

		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

// CreateOne - Append the new pokemon row to CSV file.
func (pi *pokemonInteractor) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *pokerrors.UseCaseError) {
	p, err := pi.PokemonRepository.CreateOne(pokemon)
	if err != nil {
		zap.S().Errorf("INTERACTOR: Error storing pokemon record %s", err)
		useCaseException := pokerrors.GenerateUseCaseError("Error storing pokemon record")
		return nil, &useCaseException
	}
	return pi.PokemonPresenter.ResponsePokemon(p), nil
}

// GetPokemonByName - Get the pokemon from PokeAPI given the name and return it as Pokemon model.
func (pi *pokemonInteractor) GetPokemonByName(name string) (*model.Pokemon, *pokerrors.UseCaseError) {
	resp, err := pi.HTTPClient.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if resp.HTTPStatus != http.StatusOK {
		zap.S().Errorf("INTEREACTOR: Third part API response with an error status: ", resp.HTTPStatus)
		useCaseException := pokerrors.GenerateUseCaseError("Request error")

		return nil, &useCaseException
	}

	if err != nil {
		zap.S().Errorf("INTERACTOR: Error making request %s", err)
		useCaseException := pokerrors.GenerateUseCaseError("Error making request")

		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponseMappedPokemonFromAPI(resp), nil
}

// BatchFilter - Create the worker pool and dispatch the jobs to recover the items from CSV.
func (pi *pokemonInteractor) BatchFilter(disc string, items int, itemsPerworker int) []*model.Pokemon {
	disp := pi.WorkerPool.SetPoolSize(items, itemsPerworker, disc).Start()
	defer disp.Stop()

	disp.Submit(pool.Job{
		ID:        1,
		Name:      fmt.Sprintf("JobID::%d", 1),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	resp := <-disp.OutputChannel
	return resp
}
