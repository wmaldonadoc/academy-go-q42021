package repository

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
)

type pokemonRepository struct {
	db []*model.Pokemon
}

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, error)
}

func NewPokemonRepository(db []*model.Pokemon) PokemonRepository {
	return &pokemonRepository{db}
}

func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, error) {

	for _, poke := range pr.db {
		if poke.ID == id {
			return poke, nil
		}
	}
	return nil, nil
}
