package interactor

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	ucExceptions "github.com/wmaldonadoc/academy-go-q42021/usecase/exceptions"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/repository"

	"go.uber.org/zap"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) *pokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pi *pokemonInteractor) GetByID(id int) (*model.Pokemon, *ucExceptions.UseCaseError) {
	p, err := pi.PokemonRepository.FindById(id)
	if err != nil {
		zap.S().Errorf("Error getting pokemon %s", err.Err)
		useCaseException := ucExceptions.NewErrorWrapper(err.Code, err.HTTPStatus, err.Err, err.Message)
		return nil, &useCaseException
	}

	return pi.PokemonPresenter.ResponsePokemon(p), nil
}
