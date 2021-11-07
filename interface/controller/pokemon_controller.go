package controller

import (
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetById(c Context)
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetById(c Context) {
	id := 0
	p, _ := pc.pokemonInteractor.GetById(id)
	// if err != nil {
	// 	return err
	// }
	c.JSON(http.StatusOK, p)
}
