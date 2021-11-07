package registry

import (
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
	ip "github.com/wmaldonadoc/academy-go-q42021/interface/presenter"
	ir "github.com/wmaldonadoc/academy-go-q42021/interface/repository"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
	up "github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	ur "github.com/wmaldonadoc/academy-go-q42021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository()
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
