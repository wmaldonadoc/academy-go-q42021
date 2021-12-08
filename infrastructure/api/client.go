package api

import (
	"io/ioutil"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/pokerrors"

	"go.uber.org/zap"
)

// HTTPHandler abstract the representation of a HTTPClient and their methods.
// In this case
type HTTPHandler interface {
	Get(url string) (*http.Response, error)
}

//	APIResponse - Represents an abstraction of HTTP response.
//	- Headers: HTTP response headers
//	- Body: HTTP response body as string
//	- HTTPStatus: HTTP response status
type APIResponse struct {
	Headers    http.Header
	Body       string
	HTTPStatus int
}

// HTTPClient it's the concrete object of HTTPHandler
type HTTPClient struct {
	Client HTTPHandler
}

// NewApiClient - Returns an instance of APIResponse
func NewApiClient(clnt HTTPHandler) *HTTPClient {
	return &HTTPClient{clnt}
}

// Get - Make a HTTP request and returns the response mapped to APIResponse
func (handler *HTTPClient) Get(url string) (*APIResponse, *pokerrors.APIClientError) {
	resp, err := handler.Client.Get(url)
	if err != nil {
		zap.S().Error("Error to request GET to " + url)
		requestError := pokerrors.GenerateAPIError("Error requesting third part API")

		return nil, &requestError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		readingError := pokerrors.GenerateAPIError("Error reading response body")
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
