package pool

import (
	"io"
	"strconv"
	"time"

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
type JobQueue chan chan Job

// PokemonChan - Results output chan
type PokemonChan chan []*model.Pokemon

// End - Workers finished flag's channel
type End chan bool

// Worker
type Worker struct {
	ID            int
	JobChan       JobChannel
	Queue         JobQueue // shared between all workers and dispatchers.
	Quit          chan struct{}
	ItemsLimit    int
	OutputChannel PokemonChan
	End           End
}

func NewPokemonWorker(ID int, JobChan JobChannel, Queue JobQueue, Quit chan struct{}, itemsPerWorker int, out PokemonChan, end chan bool) *Worker {
	return &Worker{
		ID:            ID,
		JobChan:       JobChan,
		Queue:         Queue,
		Quit:          Quit,
		ItemsLimit:    itemsPerWorker,
		OutputChannel: out,
		End:           end,
	}
}

func (wr *Worker) Start() {
	go func() {
		for {
			wr.Queue <- wr.JobChan
			select {
			case job := <-wr.JobChan:
				zap.S().Infof("Job at worker: ", job)
				wr.readPokemons(wr.ItemsLimit)
			case <-wr.Quit:
				close(wr.JobChan)
				return
			case end := <-wr.End:
				zap.S().Infof("Shutting down worker", end)
				return
			}
		}
	}()
}

func (wr *Worker) Stop() {
	wr.End <- true
}

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

func (wr *Worker) readPokemons(itemsPerWorker int) {
	var result [][]string
	reader := datastore.OpenFileConcurrently()
	zap.S().Debugf("Worker ID [%d] working...", wr.ID)
	itemsPerWorker += 1 // Support the header skipping
	for i := 0; i < itemsPerWorker; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			zap.S().Infof("End of file")
			wr.End <- true
			break
		} else if err != nil {
			zap.S().Errorf("Worker error", err)
		}
		zap.S().Infof("Pokemon worker: ", record)
		result = append(result, record)
	}
	pokemon, err := parsePokemon(result[1:])
	if err != nil {
		zap.S().Error("Error parsing pokemons", err)
	}
	wr.OutputChannel <- pokemon
	zap.S().Infof("Slice length", len(result))
	zap.S().Infof("Slice content", result)
}
