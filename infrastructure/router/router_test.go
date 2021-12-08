package router

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/registry"
)

type response struct {
	name string
	path string
}

var tests = []struct {
	TestCase response
}{
	{TestCase: response{name: "Success response in getting a pokemon by id", path: "/api/v1/pokemons/id/:id"}},
	{TestCase: response{name: "Success response in getting a pokemon by name", path: "/api/v1/pokemons/name/:name"}},
	{TestCase: response{name: "Success response in getting a pokemon by name", path: "/api/v1/pokemons/filter"}},
	{TestCase: response{name: "Success response in getting service health", path: "/api/v1/health/"}},
}

func generatePokemons() []*model.Pokemon {
	var pokemons []*model.Pokemon
	p := model.Pokemon{}
	for i := 0; i < 100; i++ {
		p.ID = i
		p.Name = faker.Word()
		p.Ability = faker.Word()
		pokemons = append(pokemons, &p)
	}

	return pokemons
}
func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pokemons := generatePokemons()

	for i, test := range tests {
		r := registry.NewRegistry(pokemons)
		rtr := NewRouter(r.NewAppController())

		routerPath := rtr.Routes()[i].Path
		assert.Equal(t, test.TestCase.path, routerPath)

	}
}
