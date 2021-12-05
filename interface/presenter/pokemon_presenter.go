package presenter

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/vendors"

	"go.uber.org/zap"
)

type pokemonPresenter struct{}

// PokemonPresenter - Holds an abstraction of presenter methods.
type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
	ResponseMappedPokemonFromAPI(p *api.APIResponse) *model.Pokemon
}

// NewPokemonPresenter - Create and returns a conrete instance of pokemonPresenter.
func NewPokemonPresenter() *pokemonPresenter {
	return &pokemonPresenter{}
}

// ResponsePokemon - Returns the Pokemon given.
func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	return p
}

// ResponseMappedPokemonFromAPI - Receives an APIResponse, deserialize the JSON string and mapped to Pokemon model.
// It will generate a random ID (between 10 - 1000) and pick the first ability in the response.
func (pp *pokemonPresenter) ResponseMappedPokemonFromAPI(p *api.APIResponse) *model.Pokemon {
	// Generating random id
	rand.Seed(time.Now().UnixNano())
	max := constants.MAXIDALLOWED
	min := constants.MINIDALLOWED
	id := rand.Intn(max-min+1) + min
	zap.S().Info("PRESENTER: ID generate: ", id)
	// mapping & deserialize JSON
	var responseObject vendors.Response
	json.Unmarshal([]byte(p.Body), &responseObject)
	pokemon := model.Pokemon{
		ID:      id,
		Name:    responseObject.Name,
		Ability: responseObject.Abilities[0].Ability.Name,
	}
	zap.S().Infof("PRESENTER: New pokemon generated %s", pokemon)

	return &pokemon
}
