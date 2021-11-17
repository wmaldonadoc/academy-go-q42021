package vendors

import (
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/exceptions"
)

type HTTPClient interface {
	Get(url string) (*api.ApiResponse, *exceptions.APIClientException)
}
