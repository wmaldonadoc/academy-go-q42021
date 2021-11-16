package exceptions

import (
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
)

type RequestError struct {
	Message    string
	HTTPStatus int
	Code       int
}

func PokemonNotFoundException() *RequestError {
	return &RequestError{
		Message:    "Pokemon not found.",
		HTTPStatus: http.StatusNotFound,
		Code:       constants.NotFoundExceptionCode,
	}
}

func GenericException(message string, httpStatus int, code int) *RequestError {
	return &RequestError{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
	}
}

func UnprocessableEntityException(message string) *RequestError {
	return &RequestError{
		Message:    message,
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       constants.UnprocessableEntityExceptionCode,
	}
}

func ParseTypesException(source string, target string) *RequestError {
	return &RequestError{
		Message:    "Error parsing data from " + source + " to " + target,
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       constants.UnprocessableEntityExceptionCode,
	}
}
