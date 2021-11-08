package controller

import (
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
	"go.uber.org/zap"
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
	pokemonId := c.Param("id")
	zap.S().Infof("Pokemon id %s", pokemonId)
	p, _ := pc.pokemonInteractor.GetById(id)
	// if err != nil {
	// 	return err
	// }
	c.JSON(http.StatusOK, p)
}
