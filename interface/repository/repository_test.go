package repository_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"github.com/bxcodec/faker/v3"
)

func TestFindByID(t *testing.T) {

	tests := []struct {
		name    string
		id      int
		pokemon *model.Pokemon
		err     *pokerrors.RepositoryError
	}{
		{name: "Return a pokemon", id: 1, err: nil, pokemon: &model.Pokemon{
			ID:      1,
			Name:    faker.Name(),
			Ability: faker.Word(),
		}},
		{name: "Return a not found error", id: 100, pokemon: nil, err: &pokerrors.RepositoryError{
			Message:    "Pokemon not found",
			HTTPStatus: http.StatusNotFound,
			Code:       constants.NotFoundExceptionCode,
			Err:        errors.New("pokemon not found"),
		}},
	}

	for _, test := range tests {
		repo := &mocks.PokemonRepository{}

		repo.On("FindByID", test.id).Return(test.pokemon, test.err)
		result, err := repo.FindByID(test.id)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.err, err)

		repo.AssertExpectations(t)

	}
}

func TestCreateOne(t *testing.T) {

	tests := []struct {
		name    string
		id      int
		pokemon *model.Pokemon
		err     *pokerrors.RepositoryError
	}{
		{name: "Failed storing a pokemon", pokemon: nil, err: &pokerrors.RepositoryError{
			Message:    "Pokemon not found",
			HTTPStatus: http.StatusNotFound,
			Code:       constants.NotFoundExceptionCode,
			Err:        errors.New("pokemon not found"),
		}},
		{name: "Storing a pokemon", err: nil, pokemon: &model.Pokemon{
			ID:      1,
			Name:    faker.Name(),
			Ability: faker.Word(),
		}},
	}

	for _, test := range tests {

		repo := &mocks.PokemonRepository{}
		repo.On("CreateOne", test.pokemon).Return(test.pokemon, test.err)

		result, err := repo.CreateOne(test.pokemon)

		assert.Equal(t, test.pokemon, result)
		assert.Equal(t, test.err, err)

		repo.AssertExpectations(t)
	}
}
