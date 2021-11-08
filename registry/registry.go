package registry

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
)

type registry struct {
	db []*model.Pokemon
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db []*model.Pokemon) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
		Health:  r.NewHealthController(),
	}
}
