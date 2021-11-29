package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"go.uber.org/zap"
)

//	APIResponse - Represents an abstraction of HTTP response.
//	- Headers: HTTP response headers
//	- Body: HTTP response body as string
//	- HTTPStatus: HTTP response status
type APIResponse struct {
	Headers    interface{}
	Body       string
	HTTPStatus int
}

// NewApiClient - Returns an instance of APIResponse
func NewApiClient() *APIResponse {
	return &APIResponse{}
}

// Get - Make a HTTP request and returns the response mapped to APIResponse
func (a *APIResponse) Get(url string) (*APIResponse, *pokerrors.APIClientError) {
	resp, err := http.Get(url)
	if err != nil {
		zap.S().Error("Error to request GET to " + url)
		requestError := pokerrors.APIClientError{
			Message:    "Error requesting third part API.",
			HTTPStatus: http.StatusBadRequest,
			Code:       constants.ThirdPartAPIExceptionCode,
			Err:        errors.New("error requesting third part API"),
		}

		return nil, &requestError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		readingError := pokerrors.APIClientError{
			Message:    "Error reading response body.",
			HTTPStatus: http.StatusBadRequest,
			Code:       constants.ThirdPartAPIExceptionCode,
			Err:        errors.New("error reading response body"),
		}
		zap.S().Errorf("Error reading response body %s", readingError)
		return nil, &readingError
	}
	response := APIResponse{
		Headers:    resp.Header,
		Body:       string(body),
		HTTPStatus: resp.StatusCode,
	}
	zap.S().Debugf("Api response %s", response)

	return &response, nil
}
