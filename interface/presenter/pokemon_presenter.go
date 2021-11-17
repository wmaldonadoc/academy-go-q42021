package presenter

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/vendors"
	"go.uber.org/zap"
)

type pokemonPresenter struct{}

type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
	MappedPokemonFromAPI(p *api.ApiResponse) *model.Pokemon
}

func NewPokemonPresenter() *pokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	return p
}

func (pp *pokemonPresenter) MappedPokemonFromAPI(p *api.ApiResponse) *model.Pokemon {
	// Generating random id
	rand.Seed(time.Now().UnixNano())
	max := 1000
	min := 10
	id := rand.Intn(max-min+1) + min
	zap.S().Infof("PRESENTER: ID generate %i", id)
	// mapping & deserialize JSON
	var responseObject vendors.Response
	json.Unmarshal([]byte(p.Body), &responseObject)
	pokemon := model.Pokemon{
		ID:      id,
		Name:    responseObject.Name,
		Ability: "test",
	}
	zap.S().Infof("PRESENTER: New pokemon generated %s", pokemon)

	return &pokemon
}
