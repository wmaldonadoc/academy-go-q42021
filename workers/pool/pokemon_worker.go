package pool

import (
	"io"
	"strconv"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/datastore"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"go.uber.org/zap"
)

// Job - Represent a job to dispatch. Have the fields:
// ID - Job unique ID
// Name - Job unique name
// CreateAt
// UpdatedAt
type Job struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// JobChannel - A channel fot Jobs
type JobChannel chan Job

// JobQueue - Shared JobPool between the workers
type JobQueue chan *Job

// PokemonChan - Results output chan
type PokemonChan chan []*model.Pokemon

// End - Workers finished flag's channel
type End chan bool

// Worker
type Worker struct {
	ID            int         // Unique ID
	JobChan       JobChannel  // Client submits job to this channel
	Queue         JobQueue    // shared between all workers and dispatchers.
	ItemsLimit    int         // Represents the items per worker to collect
	OutputChannel PokemonChan // Holds the result of jobs
	End           End         // Worker end flag
	WorkersSize   int         // Number of workers available
}

// NewPokemonWorker - Create a instance of Worker
func NewPokemonWorker(
	ID int,
	JobChan JobChannel,
	Queue JobQueue,
	itemsPerWorker int,
	out PokemonChan,
	end End,
	workerSize int,
) *Worker {
	zap.S().Infof("Worker ID::%d created", ID)
	return &Worker{
		ID:            ID,
		JobChan:       JobChan,
		Queue:         Queue,
		ItemsLimit:    itemsPerWorker,
		OutputChannel: out,
		End:           end,
		WorkersSize:   workerSize,
	}
}

// Start - Runs the listening of jobs and distributed to workers
func (wr *Worker) Start() {
	go func() {
		for {
			job := <-wr.Queue
			if job != nil {
				wr.readPokemons(wr.ItemsLimit, *job)
			} else {
				zap.S().Debugf("Worker %d stopped", wr.ID)
				wr.End <- true
				break
			}
		}
	}()
}

// parsePokemon - Map a CSV row to slice of pokemon model
func parsePokemon(data [][]string) ([]*model.Pokemon, *pokerrors.DefaultError) {
	var pokemons []*model.Pokemon
	for _, line := range data {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			zap.S().Error("Error parsing integer -> string")
			connError := pokerrors.GenerateDefaultError("Error reading CSV")

			return nil, &connError
		}
		pokemon := datastore.GeneratePokemonsFromCSV(id, line)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

// readPokemons - Open the CSV file and each worker attach their items recollect to output channel

func (wr *Worker) readPokemons(itemsPerWorker int, job Job) {
	var result [][]string
	reader := datastore.OpenFileConcurrently()
	zap.S().Debugf("Worker ID [%d] working, need to reach %d items", wr.ID, itemsPerWorker)

	items := (itemsPerWorker + wr.WorkersSize) + constants.FIXCSVHEADER
	for j := 0; j < items; j++ {
		record, err := reader.Read()
		if err == io.EOF {
			zap.S().Infof("End of file")
			break
		} else if err != nil {
			zap.S().Errorf("Worker error", err)
			break
		}
		zap.S().Infof("Pokemon worker: ", record)
		result = append(result, record)
		zap.S().Debugf("Worker %d collect %d pokemons", wr.ID, len(result))
	}
	pokemon, pError := parsePokemon(result[1:])
	if pError != nil {
		zap.S().Error("Error parsing pokemons", pError)
	}
	wr.OutputChannel <- pokemon
}
