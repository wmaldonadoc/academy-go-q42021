package workers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
	"github.com/wmaldonadoc/academy-go-q42021/workers"
	"github.com/wmaldonadoc/academy-go-q42021/workers/pool"
)

func TestSetPoolSize(t *testing.T) {
	tests := []struct {
		name           string
		disc           string
		items          int
		itemsPerWorker int
		workers        int
		disp           *workers.Disp
	}{
		{name: "Getting a 2 workers in the pool", disc: "odd", items: 10, itemsPerWorker: 5, workers: 2, disp: &workers.Disp{}},
	}

	for _, test := range tests {
		for i := 0; i < test.workers; i++ {
			w := &pool.Worker{
				ID: 1,
			}
			test.disp.Workers = append(test.disp.Workers, w)
		}
		d := &mocks.Dispatcher{}

		d.On("SetPoolSize", test.items, test.itemsPerWorker, test.disc).Return(test.disp)

		result := d.SetPoolSize(test.items, test.itemsPerWorker, test.disc)

		assert.Equal(t, test.workers, len(result.Workers))

		d.AssertExpectations(t)
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name string
		disp *workers.Disp
	}{
		{name: "Return a dispatcher", disp: &workers.Disp{}},
	}

	for _, test := range tests {
		d := &mocks.Dispatcher{}

		d.On("Start").Return(test.disp)

		result := d.Start()

		assert.Equal(t, test.disp, result)

		d.AssertExpectations(t)
	}
}
