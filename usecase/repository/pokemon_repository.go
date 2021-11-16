package repository

import (
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/exceptions"
)

type PokemonRepository interface {
	FindById(id int) (*model.Pokemon, *exceptions.RepositoryError)
}
