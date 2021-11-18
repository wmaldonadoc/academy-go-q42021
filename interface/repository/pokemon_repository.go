package repository

import (
	"encoding/csv"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/exceptions"

	"go.uber.org/zap"
)

type pokemonRepository struct {
	db []*model.Pokemon
}

// PokemonRepository - Holds the abstraction of the repository methods.
type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, *exceptions.RepositoryError)
	CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *exceptions.RepositoryError)
}

// NewPokemonRepository - Receive a slice of pokemons and return a concrete instance od pokemonRepository.
func NewPokemonRepository(db []*model.Pokemon) *pokemonRepository {
	return &pokemonRepository{db}
}

// FindById - Find and return a pokemon given an ID.
// It will return a error if the pokemon doesn't exists.
func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, *exceptions.RepositoryError) {
	for _, poke := range pr.db {
		if poke.ID == id {
			return poke, nil
		}
	}
	repositoryError := exceptions.NewErrorWrapper(
		constants.NotFoundExceptionCode,
		errors.New("pokemon not found"),
		"Pokemon not found",
		http.StatusNotFound,
	)
	return nil, &repositoryError
}

// CreateOne - Append a new row in the CSV file.
// It will return an error if something with the CSV fail.
func (pr *pokemonRepository) CreateOne(pokemon *model.Pokemon) (*model.Pokemon, *exceptions.RepositoryError) {
	fileLocation := config.GetEnvVariable("FILE_LOCATION")
	if file, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_APPEND, 0644); err == nil {
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
	repositoryError := exceptions.NewErrorWrapper(
		constants.WritingCSVFileExceptionCode,
		errors.New("error writing csv file"),
		"Error storing the record",
		http.StatusInternalServerError,
	)
	return nil, &repositoryError
}
