package repository

import (
	"errors"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/exceptions"
)

type pokemonRepository struct {
	db []*model.Pokemon
}

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, *exceptions.RepositoryError)
}

func NewPokemonRepository(db []*model.Pokemon) *pokemonRepository {
	return &pokemonRepository{db}
}

func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, *exceptions.RepositoryError) {
	for _, poke := range pr.db {
		if poke.ID == id {
			return poke, nil
		}
	}
	repositoryError := exceptions.NewErrorWrapper(
		constants.NotFoundExceptionCode,
		errors.New("pokemon not found"),
		"pokemon not found",
		http.StatusNotFound,
	)
	return nil, &repositoryError
}
