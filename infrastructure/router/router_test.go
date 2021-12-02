package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/registry"
)

type response struct {
	StatusCode int
	name       string
	path       string
	method     string
}

var tests = []struct {
	TestCase response
}{
	{TestCase: response{StatusCode: 200, name: "Success response in getting service health", path: "/api/v1/health/", method: "GET"}},
	{TestCase: response{StatusCode: 200, name: "Success response in getting a pokemon by id", path: "/api/v1/pokemons/id/2", method: "GET"}},
	{TestCase: response{StatusCode: 200, name: "Success response in getting a pokemon by name", path: "/api/v1/pokemons/name/ditto", method: "GET"}},
}

func doRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func generatePokemons() []*model.Pokemon {
	var pokemons []*model.Pokemon
	p := model.Pokemon{}
	for i := 0; i < 100; i++ {
		p.ID = i
		p.Name = faker.Word()
		p.Ability = faker.Word()
		fmt.Printf("Pokemons generated %v", p)
		pokemons = append(pokemons, &p)
	}

	return pokemons
}
func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pokemons := generatePokemons()

	for _, e := range tests {
		r := registry.NewRegistry(pokemons)
		rtr := NewRouter(r.NewAppController())

		w := doRequest(rtr, e.TestCase.method, e.TestCase.path)
		if !reflect.DeepEqual(e.TestCase.StatusCode, w.Code) {
			t.Fatalf("%s: Failed expect %d but got %d ", e.TestCase.name, e.TestCase.StatusCode, w.Code)
		}

	}
}
