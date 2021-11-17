package presenter

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
)

type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
	MappedPokemonFromAPI(p *api.ApiResponse) *model.Pokemon
}
