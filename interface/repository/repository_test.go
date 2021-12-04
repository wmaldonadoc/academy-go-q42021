package repository

import (
	"encoding/json"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"

	"github.com/bxcodec/faker/v3"
)

func TestFindByID(t *testing.T) {
	var pokemons []*model.Pokemon

	for i := 0; i < 10; i++ {
		pokemon := model.Pokemon{
			ID:      i,
			Name:    faker.Name(),
			Ability: faker.Word(),
		}
		pokemons = append(pokemons, &pokemon)
	}

	tests := []struct {
		name    string
		id      int
		errCode int
	}{
		{name: "Return a pokemon", id: 1},
		{name: "Return a not found error", id: 100, errCode: constants.NotFoundExceptionCode},
	}

	repo := NewPokemonRepository(pokemons)
	for _, test := range tests {
		result, err := repo.FindByID(test.id)

		if test.errCode != 0 {
			if !reflect.DeepEqual(err.Code, test.errCode) {
				t.Errorf("%s: Expected %d but got %d", test.name, test.errCode, err.Code)
			}
		}

		if test.errCode == 0 {
			if !reflect.DeepEqual(result.ID, test.id) {
				t.Errorf("%s: Expected %d but got %d", test.name, test.id, result.ID)
			}
		}
	}
}

func TestCreateOne(t *testing.T) {
	var pokemons []*model.Pokemon

	for i := 0; i < 10; i++ {
		pokemon := model.Pokemon{
			ID:      i,
			Name:    faker.Name(),
			Ability: faker.Word(),
		}
		pokemons = append(pokemons, &pokemon)
	}

	tests := []struct {
		name       string
		id         int
		errCode    int
		shouldFail bool
	}{
		{name: "Failed storing a pokemon", errCode: constants.DefaultExceptionCode, shouldFail: true},
		{name: "Storing a pokemon", shouldFail: false},
	}

	repo := NewPokemonRepository(pokemons)

	for i, test := range tests {
		json, _ := json.Marshal(pokemons)
		mockFS := fstest.MapFS{
			"data.csv": {
				Data: []byte(json),
			},
		}

		data, errFS := mockFS.ReadFile("data.csv")

		if errFS != nil {
			t.Error(errFS)
		}
		t.Log(data)
		pokemonAt := pokemons[i]
		test.id = pokemonAt.ID
		result, err := repo.CreateOne(pokemonAt)

		if test.shouldFail {
			if !reflect.DeepEqual(test.errCode, err.Code) {
				t.Errorf("%s: Expected error code %d but got %d", test.name, test.errCode, err.Code)
			}
		} else {
			if !reflect.DeepEqual(result.ID, test.id) {
				t.Errorf("%s: Expected %d but got %d", test.name, test.id, result.ID)
			}
		}
	}
}
