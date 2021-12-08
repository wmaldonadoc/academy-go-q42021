package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
)

func TestGetPokemonById(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tests := []struct {
		name       string
		keyParam   string
		valueParam string
		Response   *controller.ControllerResponse
	}{
		{name: "Return a successfull response", keyParam: "id", valueParam: "10", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusOK,
		}},
		{name: "Return a bad request response", keyParam: "test", valueParam: "4", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusBadRequest,
		}},
		{name: "Return a not found response", keyParam: "id", valueParam: "200", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusNotFound,
		}},
	}

	for _, test := range tests {
		ctx.Params = []gin.Param{
			{
				Key:   test.keyParam,
				Value: test.valueParam,
			},
		}

		mockPokemonController := &mocks.PokemonController{}

		mockPokemonController.On("GetByID", ctx).Return(test.Response)

		resp := mockPokemonController.GetByID(ctx)

		assert.Equal(t, test.Response.HTTPStatus, resp.HTTPStatus)
		mockPokemonController.AssertExpectations(t)

	}
}

func TestGetByName(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tests := []struct {
		name       string
		keyParam   string
		valueParam string
		Response   *controller.ControllerResponse
	}{
		{name: "Return a successfull response", keyParam: "name", valueParam: "ditto", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusOK,
		}},
		{name: "Return a bad request response", keyParam: "", valueParam: "4", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
		}},
		{name: "Return a not found response", keyParam: "name", valueParam: "benito", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusNotFound,
		}},
	}

	for _, test := range tests {
		ctx.Params = []gin.Param{
			{
				Key:   test.keyParam,
				Value: test.valueParam,
			},
		}

		mockPokemonController := &mocks.PokemonController{}

		mockPokemonController.On("GetByName", ctx).Return(test.Response)

		resp := mockPokemonController.GetByName(ctx)

		assert.Equal(t, test.Response.HTTPStatus, resp.HTTPStatus)
		mockPokemonController.AssertExpectations(t)
	}
}

func TestFilterSearching(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tests := []struct {
		name string

		Response *controller.ControllerResponse
	}{
		{name: "Return a successfull response", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusOK,
		}},
		{name: "Return a unprocessable entity response", Response: &controller.ControllerResponse{
			HTTPStatus: http.StatusUnprocessableEntity,
		}},
	}

	for _, test := range tests {

		mockPokemonController := &mocks.PokemonController{}

		mockPokemonController.On("FilterSearching", ctx).Return(test.Response)

		resp := mockPokemonController.FilterSearching(ctx)

		assert.Equal(t, test.Response.HTTPStatus, resp.HTTPStatus)
		mockPokemonController.AssertExpectations(t)
	}
}
