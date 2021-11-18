package exceptions

import (
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
)

/*
	RequestError - Represents exceptions with the following fields:
	- Message
	- HTTPStatus
	- Code: Represents a custom error code
*/
type RequestError struct {
	Message    string
	HTTPStatus int
	Code       int
}

// PokemonNotFoundException - Creates and return an instance of RequestError with default values.
func PokemonNotFoundException() *RequestError {
	return &RequestError{
		Message:    "Pokemon not found.",
		HTTPStatus: http.StatusNotFound,
		Code:       constants.NotFoundExceptionCode,
	}
}

// GenericException - Creates and return an instance of RequestError given a message, httpStatus and code.
func GenericException(message string, httpStatus int, code int) *RequestError {
	return &RequestError{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
	}
}

// UnprocessableEntityException - Creates and return an instance of RequestError given a message.
func UnprocessableEntityException(message string) *RequestError {
	return &RequestError{
		Message:    message,
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       constants.UnprocessableEntityExceptionCode,
	}
}

// ParseTypesException - Creates and return an instance of RequestError given a source & target types.
func ParseTypesException(source string, target string) *RequestError {
	return &RequestError{
		Message:    "Error parsing data from " + source + " to " + target,
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       constants.UnprocessableEntityExceptionCode,
	}
}
