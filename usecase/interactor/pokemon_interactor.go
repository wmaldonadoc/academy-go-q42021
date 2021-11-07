package interactor

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetById(id int) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pi *pokemonInteractor) GetById(id int) (*model.Pokemon, error) {
	p, err := pi.PokemonRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return pi.PokemonPresenter.ResponsePokemon(p), nil
}
