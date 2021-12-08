package api_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/mocks"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"
)

// TestGet - Test the wrapper of HTTP handler.
// Should return a success HTTPStatus when it's a valid request and a specific error when somethin goes wrong
func TestGet(t *testing.T) {
	// Arrange
	sampleErr := errors.New("failed request")
	sampleJSON := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	sampleBody := io.NopCloser(bytes.NewReader([]byte(sampleJSON)))

	// Test table
	tests := []struct {
		URL         string
		HTTPStatus  int
		name        string
		ErrorCustom *pokerrors.APIClientError
		Response    *http.Response
		APIResponse *api.APIResponse
		Error       error
	}{
		{
			URL:         "https://jsonplaceholder.typicode.com/todos/1",
			HTTPStatus:  http.StatusOK,
			name:        "Making a valid GET request",
			ErrorCustom: nil,
			Response: &http.Response{
				Body:       sampleBody,
				StatusCode: http.StatusOK,
			},
			APIResponse: &api.APIResponse{
				Body:       sampleJSON,
				HTTPStatus: http.StatusOK,
			},
			Error: nil,
		},
		{
			URL:        "http://fakeAPI.com",
			HTTPStatus: http.StatusOK,
			name:       "Making an invalid GET request",
			ErrorCustom: &pokerrors.APIClientError{
				Message:    "Error requesting third part API",
				HTTPStatus: http.StatusBadRequest,
				Code:       constants.ThirdPartAPIExceptionCode,
				Err:        errors.New("Error requesting third part API"),
			},
			Response: nil,
			Error:    sampleErr,
		},
	}

	for _, test := range tests {
		httpService := &mocks.HTTPHandler{}
		httpService.
			On("Get", test.URL).
			Return(
				test.Response,
				test.Error,
			).
			Once()

		h := &api.HTTPClient{
			Client: httpService,
		}

		response, err := h.Get(test.URL)
		assert.Equal(t, test.ErrorCustom, err)
		assert.Equal(t, test.APIResponse, response)

		httpService.AssertExpectations(t)
	}
}
