package exceptions

/*
	APIClientError - Represents exceptions with the following fields:
	- Message
	- HTTPStatus
	- Code: Represents a custom error code
	- Err: An Error from errors package
*/
type APIClientError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}

// NewErrorWrapper - Create and returns an instance of APIClientError
func NewErrorWrapper(code int, err error, message string, httpStatus int) APIClientError {
	return APIClientError{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
		Err:        err,
	}
}

// Error - Returns the built-in error/message passend in the wrapped function.
func (err APIClientError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}
