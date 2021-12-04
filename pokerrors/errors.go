package pokerrors

import (
	"errors"
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
)

//	APIClientError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type APIClientError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

//	RepositoryError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type RepositoryError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

//	DefaultError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type DefaultError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

//	UnprocessableEntityError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type UnprocessableEntityError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

//	UseCaseError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type UseCaseError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

// GenerateAPIError - Generate an instance of APIClientError
func GenerateAPIError(msg string) APIClientError {
	return APIClientError{
		Message:    msg,
		Code:       constants.ThirdPartAPIExceptionCode,
		HTTPStatus: http.StatusBadRequest,
		Err:        errors.New(msg),
	}
}

// GenerateDefaultError - Generate an instance of DefaultError
func GenerateDefaultError(msg string) DefaultError {
	return DefaultError{
		Message:    msg,
		Code:       constants.DefaultExceptionCode,
		HTTPStatus: http.StatusInternalServerError,
		Err:        errors.New(msg),
	}
}

// GenerateUnprocessableEntityError - Generate an instance of UnprocessableEntityError
func GenerateUnprocessableEntityError(msg string) UnprocessableEntityError {
	return UnprocessableEntityError{
		Message:    msg,
		Code:       constants.UnprocessableEntityExceptionCode,
		HTTPStatus: http.StatusUnprocessableEntity,
		Err:        errors.New(msg),
	}
}

// GenerateRepositoryError - Generate an instance of RepositoryError
func GenerateRepositoryError(msg string) RepositoryError {
	return RepositoryError{
		Message:    msg,
		Code:       constants.DefaultExceptionCode,
		HTTPStatus: http.StatusInternalServerError,
		Err:        errors.New(msg),
	}
}

// GenerateNotFoundError - Generate an instance of RepositoryError
func GenerateNotFoundError(msg string) RepositoryError {
	return RepositoryError{
		Message:    msg,
		Code:       constants.NotFoundExceptionCode,
		HTTPStatus: http.StatusNotFound,
		Err:        errors.New(msg),
	}
}

// GenerateUseCaseError - Generate an instance of UseCaseError
func GenerateUseCaseError(msg string) UseCaseError {
	return UseCaseError{
		Message:    msg,
		Code:       constants.DefaultExceptionCode,
		HTTPStatus: http.StatusNotFound,
		Err:        errors.New(msg),
	}
}
