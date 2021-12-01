package datastore

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/wmaldonadoc/academy-go-q42021/config"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"go.uber.org/zap"
)

type FS struct {
	Open func(name string) (*os.File, error)
}

func openFile(fileLocation string) ([][]string, *pokerrors.DefaultError) {
	file, err := os.Open(fileLocation)
	zap.S().Debugf("File located at: ", fileLocation)
	if err != nil {
		zap.S().Errorf("Error opening the CSV file %s", err)
		conError := pokerrors.GenerateDefaultError(err.Error())
		return nil, &conError
	}
	zap.S().Info("Successfully openned csv file")

	defer file.Close()

	row1, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		conError := pokerrors.GenerateDefaultError(err.Error())
		return nil, &conError
	}
	_, err = file.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		conError := pokerrors.GenerateDefaultError(err.Error())
		return nil, &conError
	}

	chunks, err := csv.NewReader(file).ReadAll()
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		conError := pokerrors.GenerateDefaultError(err.Error())
		return nil, &conError
	}

	return chunks, nil
}

// OpenFileConcurrently - Open a CSV file given an env var, and return the reader.
func OpenFileConcurrently() *csv.Reader {
	fileLocation := config.GetEnvVariable("FILE_LOCATION")
	csvfile, err := os.Open(fileLocation)
	if err != nil {
		zap.S().Errorf("Error reading the CSV file concurrently %s", err)
		log.Fatal(err)
	}
	// defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	return reader
}

// GeneratePokemonsFromCSV - Receive the id of pokemon and return a instance of Pokemon model.
func GeneratePokemonsFromCSV(id int, data []string) *model.Pokemon {
	return &model.Pokemon{
		ID:      id,
		Name:    data[1],
		Ability: data[2],
	}
}

// NewCSV - Open and reads a CSV file and return it as a slice of pokemons.
func NewCSV(fileLocation string) ([]*model.Pokemon, *pokerrors.DefaultError) {
	pokemones := []*model.Pokemon{}
	chunks, err := openFile(fileLocation)
	if err != nil {
		zap.S().Error("Error with datastore connection", err)
		connError := pokerrors.GenerateDefaultError(err.Message)
		return nil, &connError
	}
	zap.S().Debug("-------- START READING CSV --------")
	for _, line := range chunks {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			zap.S().Error("Error parsing integer -> string")
			connError := pokerrors.GenerateDefaultError("Error reading CSV")

			return nil, &connError
		}
		zap.S().Debugf("ID %s", line[0])
		zap.S().Debugf("Name %s", line[1])
		zap.S().Debugf("Ability %s", line[2])
		pokemon := GeneratePokemonsFromCSV(id, line)
		pokemones = append(pokemones, pokemon)
	}
	zap.S().Debug("-------- END READING CSV --------")
	zap.S().Debugf("Pokemons availables: %s", pokemones)

	return pokemones, nil
}
