package presenter

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/vendors"
)

func TestResponsePokemon(t *testing.T) {
	tests := []struct {
		name    string
		pokemon *model.Pokemon
	}{
		{name: "Return the same pokemon passed as parameter"},
	}

	pr := NewPokemonPresenter()

	for _, test := range tests {
		id, _ := faker.RandomInt(1, 100)
		test.pokemon = &model.Pokemon{
			ID:      id[0],
			Ability: faker.Word(),
			Name:    faker.Name(),
		}

		result := pr.ResponsePokemon(test.pokemon)

		if !reflect.DeepEqual(test.pokemon.ID, result.ID) {
			t.Errorf("%s: Expected ID %d but got %d", test.name, test.pokemon.ID, result.ID)
		}

		if !reflect.DeepEqual(test.pokemon.Ability, result.Ability) {
			t.Errorf("%s: Expected ability %s but got %s", test.name, test.pokemon.Ability, result.Ability)
		}

		if !reflect.DeepEqual(test.pokemon.Name, result.Name) {
			t.Errorf("%s: Expected name %s but got %s", test.name, test.pokemon.Name, result.Name)
		}
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
		HTTPStatus int
	}{
		{name: "Return the same pokemon passed as parameter", HTTPStatus: http.StatusOK},
	}

	pr := NewPokemonPresenter()

	for _, test := range tests {
		mockResp := mockAPIResponse()
		json, _ := json.Marshal(mockResp)
		test.response = &api.APIResponse{
			Headers:    faker.Word(),
			HTTPStatus: test.HTTPStatus,
			Body:       string(json),
		}

		result := pr.ResponseMappedPokemonFromAPI(test.response)

		if result.ID == 0 {
			t.Errorf("%s: The pokemon should have an ID", test.name)
		}

		if result.Name == "" {
			t.Errorf("%s: The pokemon should have a name", test.name)
		}

		if result.Ability == "" {
			t.Errorf("%s: The pokemon should have an ability", test.name)
		}
	}
}
