package repository

import "github.com/wmaldonadoc/academy-go-q42021/domain/model"

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, error)
}
