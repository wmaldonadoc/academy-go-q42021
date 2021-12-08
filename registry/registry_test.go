package registry_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
)

func TestRegistry(t *testing.T) {
	pokemonController := &mocks.PokemonController{}
	healthController := &mocks.HealthController{}
	tests := []struct {
		name     string
		response controller.AppController
	}{{name: "Should return a registry with two controllers", response: controller.AppController{
		Pokemon: pokemonController,
		Health:  healthController,
	}}}

	for _, test := range tests {
		r := &mocks.Registry{}

		r.On("NewAppController").Return(test.response)
		t.Logf("Running test: %s", test.name)

		result := r.NewAppController()

		assert.Implements(t, (*controller.PokemonController)(nil), result.Pokemon)
		assert.Implements(t, (*controller.HealthController)(nil), result.Health)

		r.AssertExpectations(t)
	}
}
