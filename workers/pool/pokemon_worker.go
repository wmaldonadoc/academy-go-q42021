package pool

import (
	"io"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/datastore"

	"go.uber.org/zap"
)

type Job struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobChannel chan Job
type JobQueue chan chan Job
type PokemonChan chan [][]string
type End chan bool

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
	wr.OutputChannel <- result[1:] // Skip header line
	zap.S().Infof("Slice length", len(result))
	zap.S().Infof("Slice content", result)
}
