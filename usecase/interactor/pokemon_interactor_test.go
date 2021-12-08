package interactor_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
)

func TestGetByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		pokemon *model.Pokemon
		err     *pokerrors.UseCaseError
	}{
		{name: "Return a valid pokemon", id: 1, err: nil, pokemon: &model.Pokemon{
			ID:      1,
			Name:    faker.Word(),
			Ability: faker.Word(),
		}},
		{name: "Return an error", id: 1, err: &pokerrors.UseCaseError{
			Message:    "error getting pokemon",
			Code:       constants.DefaultExceptionCode,
			HTTPStatus: http.StatusNotFound,
			Err:        errors.New("error getting pokemon"),
		}, pokemon: nil},
	}

	for _, test := range tests {
		itr := &mocks.PokemonInteractor{}

		itr.On("GetByID", test.id).Return(test.pokemon, test.err)

		result, err := itr.GetByID(test.id)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.err, err)

		itr.AssertExpectations(t)
	}

}

func TestCreateOne(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		pokemon *model.Pokemon
		err     *pokerrors.UseCaseError
	}{
		{name: "Return a valid pokemon", id: 1, err: nil, pokemon: &model.Pokemon{
			ID:      1,
			Name:    faker.Word(),
			Ability: faker.Word(),
		}},
		{name: "Return an error", id: 1, err: &pokerrors.UseCaseError{
			Message:    "error storing pokemon",
			Code:       constants.DefaultExceptionCode,
			HTTPStatus: http.StatusNotFound,
			Err:        errors.New("error storing pokemon"),
		}, pokemon: nil},
	}

	for _, test := range tests {
		itr := &mocks.PokemonInteractor{}

		itr.On("CreateOne", test.pokemon).Return(test.pokemon, test.err)

		result, err := itr.CreateOne(test.pokemon)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.err, err)

		itr.AssertExpectations(t)
	}

}

func TestGetPokemonByName(t *testing.T) {
	tests := []struct {
		name        string
		pokemonName string
		pokemon     *model.Pokemon
		err         *pokerrors.UseCaseError
	}{
		{name: "Return a valid pokemon", pokemonName: "ditto", err: nil, pokemon: &model.Pokemon{
			ID:      1,
			Name:    faker.Word(),
			Ability: faker.Word(),
		}},
		{name: "Return an error", pokemonName: "ditto", err: &pokerrors.UseCaseError{
			Message:    "error getting pokemon",
			Code:       constants.DefaultExceptionCode,
			HTTPStatus: http.StatusNotFound,
			Err:        errors.New("error getting pokemon"),
		}, pokemon: nil},
	}

	for _, test := range tests {
		itr := &mocks.PokemonInteractor{}

		itr.On("GetPokemonByName", test.name).Return(test.pokemon, test.err)

		result, err := itr.GetPokemonByName(test.name)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.err, err)

		itr.AssertExpectations(t)
	}

}

func TestBatchFilter(t *testing.T) {
	var p []*model.Pokemon
	tests := []struct {
		name           string
		disc           string
		items          int
		itemsPerworker int
		pokemon        []*model.Pokemon
		err            *pokerrors.UseCaseError
	}{
		{name: "Return a valid pokemon", items: 2, disc: "even", itemsPerworker: 2, err: nil},
	}

	for _, test := range tests {

		for i := 1; i <= test.items; i++ {
			pk := &model.Pokemon{
				ID:      i,
				Name:    faker.Word(),
				Ability: faker.Word(),
			}
			p = append(p, pk)
		}
		test.pokemon = p

		itr := &mocks.PokemonInteractor{}

		itr.On("BatchFilter", test.disc, test.items, test.itemsPerworker).Return(test.pokemon, test.err)

		result := itr.BatchFilter(test.disc, test.items, test.itemsPerworker)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.items, len(result))

		itr.AssertExpectations(t)
	}

}
