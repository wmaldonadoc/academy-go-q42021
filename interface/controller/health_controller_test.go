package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetServiceHealth(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name       string
		HTTPStatus int
	}{
		{name: "Getting a success response", HTTPStatus: http.StatusOK},
	}

	for _, test := range tests {
		healthController := NewHealthController()
		resp := healthController.GetServiceHealth(ctx)

		if !reflect.DeepEqual(resp.HTTPStatus, test.HTTPStatus) {
			t.Errorf("%s: Expected %d but got %d", test.name, test.HTTPStatus, resp.HTTPStatus)
		}

		if resp.HTTPStatus == 0 {
			t.Errorf("%s: Expected the field 'HTTPStatus'", test.name)
		}
		if resp.Data == nil {
			t.Errorf("%s: Expected the field 'Data'", test.name)
		}
	}
}
