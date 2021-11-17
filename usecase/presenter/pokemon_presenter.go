package presenter

import "github.com/wmaldonadoc/academy-go-q42021/domain/model"

type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
}
