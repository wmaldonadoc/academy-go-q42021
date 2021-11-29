package repository

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
)

// PokemonRepository - Holds the abstraction of repository methods.
type PokemonRepository interface {
	// FindById - Find and return a pokemon given an ID.
	// It will return a error if the pokemon doesn't exists.
	FindById(id int) (*model.Pokemon, *pokerrors.RepositoryError)
	// CreateOne - Append a new row in the CSV file.
	// It will return an error if something with the CSV fail.
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *pokerrors.RepositoryError)
}
