package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/controller"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
)

func TestGetServiceHealth(t *testing.T) {
	startTime := time.Now()
	uptime := time.Since(startTime)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name       string
		HTTPStatus int
		Response   *controller.ControllerResponse
	}{
		{
			name:       "Getting a success response",
			HTTPStatus: http.StatusOK,
			Response: &controller.ControllerResponse{
				HTTPStatus: http.StatusOK,
				Data: model.Health{
					Uptime:     uptime,
					StatusCode: http.StatusOK,
				},
			},
		},
	}

	for _, test := range tests {
		mockHealthController := &mocks.HealthController{}

		mockHealthController.
			On("GetServiceHealth", ctx).
			Return(test.Response)

		response := mockHealthController.GetServiceHealth(ctx)

		assert.Equal(t, test.Response.HTTPStatus, response.HTTPStatus)
		mockHealthController.AssertExpectations(t)
	}
}
