package repository

import "github.com/wmaldonadoc/academy-go-q42021/domain/model"

type pokemonRepository struct{}

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, error) {
	p := model.Pokemon{ID: 1, Name: "Poliwhirl", Ability: "Damp"}

	return &p, nil
}
