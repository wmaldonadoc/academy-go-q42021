package presenter

import "github.com/wmaldonadoc/academy-go-q42021/domain/model"

type pokemonPresenter struct{}

type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	return p
}
