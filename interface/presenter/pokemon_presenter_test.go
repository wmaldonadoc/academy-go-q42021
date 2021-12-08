package presenter

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/vendors"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
)

func TestResponsePokemon(t *testing.T) {
	tests := []struct {
		name    string
		pokemon *model.Pokemon
	}{
		{name: "Return the same pokemon passed as parameter"},
	}

	for _, test := range tests {
		pr := &mocks.PokemonPresenter{}

		id, _ := faker.RandomInt(1, 100)
		test.pokemon = &model.Pokemon{
			ID:      id[0],
			Ability: faker.Word(),
			Name:    faker.Name(),
		}
		pr.On("ResponsePokemon", test.pokemon).Return(test.pokemon)

		result := pr.ResponsePokemon(test.pokemon)

		assert.Equal(t, test.pokemon.Ability, result.Ability)
		assert.Equal(t, test.pokemon.Name, result.Name)
		assert.Equal(t, test.pokemon.ID, result.ID)

		pr.AssertExpectations(t)
	}
}

func mockAPIResponse() *vendors.Response {
	var mPokemons []vendors.Pokemon
	var mAbilities []vendors.PokemonAbilities
	for i := 1; i <= 10; i++ {
		// Mocking Pokemon
		mockSpecie := vendors.PokemonSpecies{
			Name: faker.Word(),
		}
		mockPokemon := vendors.Pokemon{
			EntryNo: i,
			Species: mockSpecie,
		}
		mPokemons = append(mPokemons, mockPokemon)

		// Mocking pokemon abilities
		mockAbility := vendors.PokemonAbility{
			Name: faker.Name(),
			URL:  faker.URL(),
		}
		mockPokemonAbolities := vendors.PokemonAbilities{
			Ability:  mockAbility,
			IsHidden: true,
			Slot:     i,
		}

		mAbilities = append(mAbilities, mockPokemonAbolities)
	}

	mockResponse := vendors.Response{
		Name:      faker.Name(),
		Pokemon:   mPokemons,
		Abilities: mAbilities,
	}

	return &mockResponse
}

func TestResponseMappedPokemonFromAPI(t *testing.T) {
	tests := []struct {
		name       string
		response   *api.APIResponse
		pokemon    *model.Pokemon
		HTTPStatus int
	}{
		{name: "Return the same pokemon passed as parameter", HTTPStatus: http.StatusOK},
	}

	for _, test := range tests {
		mockResp := mockAPIResponse()
		json, _ := json.Marshal(mockResp)
		test.response = &api.APIResponse{
			HTTPStatus: test.HTTPStatus,
			Body:       string(json),
		}
		id, _ := faker.RandomInt(1, 100)
		test.pokemon = &model.Pokemon{
			ID:      id[0],
			Ability: faker.Word(),
			Name:    faker.Name(),
		}

		pr := &mocks.PokemonPresenter{}
		pr.On("ResponseMappedPokemonFromAPI", test.response).Return(test.pokemon)
		result := pr.ResponseMappedPokemonFromAPI(test.response)

		assert.Equal(t, test.pokemon.Ability, result.Ability)
		assert.Equal(t, test.pokemon.Name, result.Name)
		assert.Equal(t, test.pokemon.ID, result.ID)

		pr.AssertExpectations(t)
	}
}
