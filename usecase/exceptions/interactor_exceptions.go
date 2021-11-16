package exceptions

type UseCaseError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}

func NewErrorWrapper(code int, httpStatus int, err error, message string) UseCaseError {
	return UseCaseError{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
		Err:        err,
	}
}

func (err UseCaseError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}
