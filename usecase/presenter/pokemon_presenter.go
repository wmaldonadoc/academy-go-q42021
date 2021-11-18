package presenter

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
)

// PokemonPresenter - Holds the abstractions of the presenter methods.
type PokemonPresenter interface {
	// ResponsePokemon - Returns the Pokemon given.
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
	// ResponseMappedPokemonFromAPI - Receives an ApiResponse, deserialize the JSON string and mapped to Pokemon model.
	// It will generate a random ID (between 10 - 1000) and pick the first ability in the response.
	ResponseMappedPokemonFromAPI(p *api.ApiResponse) *model.Pokemon
}
