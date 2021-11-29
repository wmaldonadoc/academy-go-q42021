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

	"go.uber.org/zap"
)

func openFile() [][]string {
	fileLocation := config.GetEnvVariable("FILE_LOCATION")
	file, err := os.Open(fileLocation)
	if err != nil {
		zap.S().Errorf("Error opening the CSV file %s", err)
	}
	zap.S().Info("Successfully openned csv file")

	defer file.Close()

	row1, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		log.Fatal(err)
	}
	_, err = file.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		log.Fatal(err)
	}

	chunks, err := csv.NewReader(file).ReadAll()
	if err != nil {
		zap.S().Errorf("Error reading the CSV file %s", err)
		log.Fatal(err)
	}

	return chunks
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

func generatePokemonsFromCSV(id int, data []string) *model.Pokemon {
	return &model.Pokemon{
		ID:      id,
		Name:    data[1],
		Ability: data[2],
	}
}

// NewCSV - Open and reads a CSV file and return it as a slice of pokemons.
func NewCSV() []*model.Pokemon {
	pokemones := []*model.Pokemon{}
	chunks := openFile()
	zap.S().Debug("-------- START READING CSV --------")
	for _, line := range chunks {
		if id, err := strconv.Atoi(line[0]); err == nil {

			zap.S().Debugf("ID %s", line[0])
			zap.S().Debugf("Name %s", line[1])
			zap.S().Debugf("Ability %s", line[2])
			pokemon := generatePokemonsFromCSV(id, line)
			pokemones = append(pokemones, pokemon)
		} else {
			zap.S().Error("Error parsing integer -> string")
			log.Fatal(err)
		}
	}
	zap.S().Debug("-------- END READING CSV --------")
	zap.S().Debugf("Pokemons availables: %s", pokemones)

	return pokemones
}
