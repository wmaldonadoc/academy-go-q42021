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

func NewDb() []*model.Pokemon {
	pokemones := []*model.Pokemon{}
	chunks := openFile()
	zap.S().Debug("-------- START READING CSV --------")
	for _, line := range chunks {
		if id, err := strconv.Atoi(line[0]); err == nil {

			zap.S().Debugf("ID %s", line[0])
			zap.S().Debugf("Name %s", line[1])
			zap.S().Debugf("Ability %s", line[2])
			pokemon := model.Pokemon{
				ID:      id,
				Name:    line[1],
				Ability: line[2],
			}
			pokemones = append(pokemones, &pokemon)
		} else {
			zap.S().Error("Error parsing integer -> string")
			log.Fatal(err)
		}
	}
	zap.S().Debug("-------- END READING CSV --------")
	zap.S().Debugf("Pokemons availables: %q", pokemones)

	return pokemones
}
