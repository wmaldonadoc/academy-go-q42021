package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/exceptions"

	"go.uber.org/zap"
)

/*
	ApiResponse - Represents an abstraction of HTTP response.
	- Headers: HTTP response headers
	- Body: HTTP response body as string
	- HTTPStatus: HTTP response status
*/
type ApiResponse struct {
	Headers    interface{}
	Body       string
	HTTPStatus int
}

// NewApiClient - Returns an instance of ApiResponse
func NewApiClient() *ApiResponse {
	return &ApiResponse{}
}

// Get - Make a HTTP request and returns the response mapped to ApiResponse
func (a *ApiResponse) Get(url string) (*ApiResponse, *exceptions.APIClientError) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			readingError := exceptions.NewErrorWrapper(
				constants.ThirdPartAPIExceptionCode,
				err,
				"Error reading response body",
				http.StatusBadRequest,
			)
			zap.S().Errorf("Error reading response body %s", readingError)
			return nil, &readingError
		}
		response := ApiResponse{
			Headers:    resp.Header,
			Body:       string(body),
			HTTPStatus: resp.StatusCode,
		}
		// zap.S().Info("Api response %s", response)

		return &response, nil
	}
	zap.S().Error("Error to request GET to " + url)
	requestError := exceptions.NewErrorWrapper(
		constants.ThirdPartAPIExceptionCode,
		errors.New("request error"),
		"Error in GET request",
		http.StatusBadRequest,
	)

	return nil, &requestError
}
