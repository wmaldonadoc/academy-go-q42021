package registry

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
)

type registry struct {
	db []*model.Pokemon
}

// Registry - Holds an abstraction of the AppController.
type Registry interface {
	NewAppController() controller.AppController
}

// Newregistry - Receives a slice of pokemons and return a concret instance of registry.
func NewRegistry(db []*model.Pokemon) *registry {
	return &registry{db}
}

// NewAppController - Implements the applications controllers.
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
		Health:  r.NewHealthController(),
	}
}
