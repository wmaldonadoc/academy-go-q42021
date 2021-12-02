package workers

import (
	"math"
	"sync"

	worker "github.com/wmaldonadoc/academy-go-q42021/workers/pool"

	"go.uber.org/zap"
)

type Dispatcher interface {
	// Create the worker pool and and return a Dispatch with the channels to handle the jobs.
	// The number of workers are calculated given ⌈ITEMS / ITEMSPERWORKER⌉
	SetPoolSize(num int, itemsPerWorker int) *Disp
	// Start - creates pool of num count of workers.
	// Also dispatch the worker's jobs.
	Start() *Disp
	// Stop - Iterates over every worker and add them into a wainting group
	// and run a gracefully stop when a worker done
	Stop()
	// Submit - Receive a job and add to WorkChan's queue.
	Submit(job worker.Job)
}

// Disp is the link between the client and the workers
type Disp struct {
	Workers       []*worker.Worker  // The list of workers that dispatcher tracks
	WorkChan      worker.JobChannel // Client submits job to this channel
	Queue         worker.JobQueue   // Shared JobPool between the workers
	OutputChannel worker.PokemonChan
	ItemsLimit    int
	End           worker.End
}

//  NewDispatcher - New returns a new dispatcher. A Dispatcher communicates between the client
// and the worker. Its main job is to receive a job and share it on the WorkPool
// WorkPool is the link between the dispatcher and all the workers as
// the WorkPool of the dispatcher is common JobPool for all the workers
func NewDispatcher() *Disp {
	return &Disp{}
}

// SetPoolSize - Create the worker pool and and return a Dispatch with the channels to handle the jobs.
// The number of workers are calculated given ⌈ITEMS / ITEMSPERWORKER⌉
func (d *Disp) SetPoolSize(size int, itemsPerWorker int) *Disp {
	poolSize := int(math.Floor(float64(size) / float64(itemsPerWorker)))
	zap.S().Debug("Pool size obtained: ", poolSize)
	return &Disp{
		Workers:       make([]*worker.Worker, poolSize),
		WorkChan:      make(worker.JobChannel),
		Queue:         make(worker.JobQueue),
		OutputChannel: make(worker.PokemonChan),
		ItemsLimit:    itemsPerWorker,
		End:           make(worker.End),
	}
}

// Start - creates pool of num count of workers.
// Also dispatch the worker's jobs.
func (d *Disp) Start() *Disp {
	l := len(d.Workers)
	for i := 1; i <= l; i++ {
		wrk := worker.NewPokemonWorker(
			i,
			make(worker.JobChannel),
			d.Queue,
			d.ItemsLimit,
			d.OutputChannel,
			d.End,
			len(d.Workers),
		)
		wrk.Start()
		d.Workers = append(d.Workers, wrk)
	}
	go d.process()
	return d
}

// Stop - Iterates over every worker and add them into a wainting group
// and run a gracefully stop when a worker done
func (d *Disp) Stop() {
	var wg sync.WaitGroup
	for _, w := range d.Workers {
		wg.Add(1)
		go func(wr *worker.Worker) {
			defer wg.Done()
		}(w)
	}
	wg.Wait()
	zap.S().Debugf("All workers stopped")
}

// process -  listens to a job submitted on WorkChan and
// relays it to the WorkPool. The WorkPool is shared between
// the workers.
func (d *Disp) process() {
	for data := range d.WorkChan {
		work := <-d.WorkChan

		zap.S().Infof("Job ready: ", work)
		d.Queue <- &data

	}
}

// Submit - Receive a job and add to WorkChan's queue.
func (d *Disp) Submit(job worker.Job) {
	d.WorkChan <- job
}
