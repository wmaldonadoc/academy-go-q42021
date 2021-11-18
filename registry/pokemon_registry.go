package registry

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
	ip "github.com/wmaldonadoc/academy-go-q42021/interface/presenter"
	ir "github.com/wmaldonadoc/academy-go-q42021/interface/repository"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
	up "github.com/wmaldonadoc/academy-go-q42021/usecase/presenter"
	ur "github.com/wmaldonadoc/academy-go-q42021/usecase/repository"
)

// NewPokemonController - Creates and returns an instance of controller.
func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

// NewPokemonInteractor - Creates an instance of interactor.
// Also inject the following dependencies: Repository, Presenter & HTTPClient.
func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter(), api.NewApiClient())
}

// NewPokemonRepository - Creates and returns an instance of repository.
// Also inject the database, in this case a slice of pokemons.
func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.db)
}

// NewPokemonPresenter - Creates and returns an instance of presenter.
func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
