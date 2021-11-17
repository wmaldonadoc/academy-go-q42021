package repository

import (
	"errors"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/exceptions"
	"go.uber.org/zap"
)

type pokemonRepository struct {
	db []*model.Pokemon
}

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, *exceptions.RepositoryError)
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *exceptions.RepositoryError)
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
		"Pokemon not found",
		http.StatusNotFound,
	)
	return nil, &repositoryError
}

func (pr *pokemonRepository) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *exceptions.RepositoryError) {
	pr.db = append(pr.db, pokemon)
	zap.S().Infof("New array: %s", pr.db)
	return nil, nil
}
