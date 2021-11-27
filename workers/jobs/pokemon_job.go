package jobs

import (
	"encoding/csv"
	"io"

	"go.uber.org/zap"
)

type PokemonJob interface {
	CreateJobs(jobCount int, workerLimit int) int
	ReadPokemonsJob(reader *csv.Reader)
}

type pokemonJob struct{}

func NewPokemonJob() *pokemonJob {
	return &pokemonJob{}
}

func (pj *pokemonJob) CreateJobs(jobCount int) int {
	return jobCount
}

func (pj *pokemonJob) ReadPokemonsJob(reader *csv.Reader) {
	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				zap.S().Infof("End of file ", err)
				break
			} else if err != nil {
				zap.S().Errorf("Error running job: ", err)
			}
			zap.S().Infof("Record: ", record)
		}
	}()
}
