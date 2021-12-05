package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/presenter"
	"github.com/wmaldonadoc/academy-go-q42021/interface/repository"
	"github.com/wmaldonadoc/academy-go-q42021/usecase/interactor"
	"github.com/wmaldonadoc/academy-go-q42021/workers"
)

func TestGetPokemonById(t *testing.T) {
	var pokemons []*model.Pokemon
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	pokemon := model.Pokemon{}
	for i := 1; i <= 10; i++ {
		pokemon.ID = i
		pokemon.Ability = faker.Word()
		pokemon.Name = faker.Name()
		pokemons = append(pokemons, &pokemon)
	}
	rp := repository.NewPokemonRepository(pokemons)
	pr := presenter.NewPokemonPresenter()
	api := api.NewApiClient(&http.Client{})
	dispatcher := workers.NewDispatcher()

	tests := []struct {
		name       string
		HTTPStatus int
		keyParam   string
		valueParam string
	}{
		{name: "Return a successfull response", HTTPStatus: 200, keyParam: "id", valueParam: "10"},
		{name: "Return a bad request response", HTTPStatus: 422, keyParam: "test", valueParam: "4"},
		{name: "Return a not found response", HTTPStatus: 404, keyParam: "id", valueParam: "200"},
	}

	for _, test := range tests {
		ctx.Params = []gin.Param{
			{
				Key:   test.keyParam,
				Value: test.valueParam,
			},
		}

		pi := interactor.NewPokemonInteractor(
			rp,
			pr,
			api,
			dispatcher,
		)
		ctrl := NewPokemonController(pi)
		resp := ctrl.GetByID(ctx)

		if !reflect.DeepEqual(resp.HTTPStatus, test.HTTPStatus) {
			t.Errorf("%s: Expected %d but got %d", test.name, test.HTTPStatus, resp.HTTPStatus)
		}

	}
}

func TestGetByName(t *testing.T) {
	var pokemons []*model.Pokemon
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	pokemon := model.Pokemon{}
	for i := 1; i <= 10; i++ {
		pokemon.ID = i
		pokemon.Ability = faker.Word()
		pokemon.Name = faker.Name()
		pokemons = append(pokemons, &pokemon)
	}
	rp := repository.NewPokemonRepository(pokemons)
	pr := presenter.NewPokemonPresenter()
	api := api.NewApiClient(&http.Client{})
	dispatcher := workers.NewDispatcher()

	tests := []struct {
		name       string
		HTTPStatus int
		keyParam   string
		valueParam string
	}{
		{name: "Return a successfull response", HTTPStatus: 200, keyParam: "name", valueParam: "ditto"},
		{name: "Return a bad request response", HTTPStatus: 422, keyParam: "", valueParam: "4"},
		{name: "Return a not found response", HTTPStatus: 404, keyParam: "name", valueParam: "benito"},
	}

	for _, test := range tests {
		ctx.Params = []gin.Param{
			{
				Key:   test.keyParam,
				Value: test.valueParam,
			},
		}

		pi := interactor.NewPokemonInteractor(
			rp,
			pr,
			api,
			dispatcher,
		)

		ctrl := NewPokemonController(pi)
		resp := ctrl.GetByName(ctx)

		t.Log(resp)
		if !reflect.DeepEqual(test.HTTPStatus, resp.HTTPStatus) {
			t.Errorf("%s: Expected %d but got %d", test.name, test.HTTPStatus, resp.HTTPStatus)
		}
	}
}

func TestFilterSearching(t *testing.T) {
	var pokemons []*model.Pokemon
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	pokemon := model.Pokemon{}
	for i := 1; i <= 10; i++ {
		pokemon.ID = i
		pokemon.Ability = faker.Word()
		pokemon.Name = faker.Name()
		pokemons = append(pokemons, &pokemon)
	}
	rp := repository.NewPokemonRepository(pokemons)
	pr := presenter.NewPokemonPresenter()
	api := api.NewApiClient(&http.Client{})
	dispatcher := workers.NewDispatcher()

	tests := []struct {
		name       string
		HTTPStatus int
		keyParam   string
		valueParam string
	}{
		{name: "Return a successfull response", HTTPStatus: 200, keyParam: "name", valueParam: "ditto"},
		{name: "Return a bad request response", HTTPStatus: 422, keyParam: "", valueParam: "4"},
		{name: "Return a not found response", HTTPStatus: 404, keyParam: "name", valueParam: "benito"},
	}

	for _, test := range tests {
		ctx.Params = []gin.Param{
			{
				Key:   test.keyParam,
				Value: test.valueParam,
			},
		}

		pi := interactor.NewPokemonInteractor(
			rp,
			pr,
			api,
			dispatcher,
		)

		ctrl := NewPokemonController(pi)
		resp := ctrl.FilterSearching(ctx)

		t.Log(resp)
		if !reflect.DeepEqual(test.HTTPStatus, resp.HTTPStatus) {
			t.Errorf("%s: Expected %d but got %d", test.name, test.HTTPStatus, resp.HTTPStatus)
		}
	}
}
