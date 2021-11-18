package vendors

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/exceptions"
)

// HTTPClient - Holds the HTTPClient method.
type HTTPClient interface {
	// Get - Make a HTTP request and returns the response mapped to ApiResponse.
	Get(url string) (*api.ApiResponse, *exceptions.APIClientError)
}
