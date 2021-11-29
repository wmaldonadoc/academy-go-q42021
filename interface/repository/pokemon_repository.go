package repository

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"go.uber.org/zap"
)

type pokemonRepository struct {
	db []*model.Pokemon
}

// PokemonRepository - Holds the abstraction of the repository methods.
type PokemonRepository interface {
	FindByID(id int) (*model.Pokemon, *pokerrors.RepositoryError)
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *pokerrors.RepositoryError)
}

// NewPokemonRepository - Receive a slice of pokemons and return a conretee instance od pokemonRepository.
func NewPokemonRepository(db []*model.Pokemon) *pokemonRepository {
	return &pokemonRepository{db}
}

// FindByID - Find and return a pokemon given an ID.
// It will return a error if the pokemon doesn't exists.
func (pr *pokemonRepository) FindByID(id int) (*model.Pokemon, *pokerrors.RepositoryError) {
	for _, poke := range pr.db {
		if poke.ID == id {
			return poke, nil
		}
	}

	repositoryError := pokerrors.GenerateNotFoundError("Pokemon not found")

	return nil, &repositoryError
}

// CreateOne - Append a new row in the CSV file.
// It will return an error if something with the CSV fail.
func (pr *pokemonRepository) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *pokerrors.RepositoryError) {
	fileLocation := config.GetEnvVariable("FILE_LOCATION")
	file, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		w := csv.NewWriter(file)
		record := []string{strconv.Itoa(pokemon.ID), pokemon.Name, pokemon.Ability}
		if err := w.Write(record); err == nil {
			w.Flush()
			pr.db = append(pr.db, pokemon)
			zap.S().Info("REPOSITORY: Pokemon stored successfully")
			zap.S().Infof("REPOSITORY: Pokemons array updated %s", pr.db)

			return pokemon, nil
		}
	}
	zap.S().Errorf("REPOSITORY: Error storing the record ", pokemon)
	repositoryError := pokerrors.GenerateRepositoryError("Error storing the record")

	return nil, &repositoryError
}
