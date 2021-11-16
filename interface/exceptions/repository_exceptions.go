package exceptions

type RepositoryError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}

func NewErrorWrapper(code int, err error, message string, httpStatus int) RepositoryError {
	return RepositoryError{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
		Err:        err,
	}
}

func (err RepositoryError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}
