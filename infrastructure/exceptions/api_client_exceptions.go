package exceptions

type APIClientException struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}

func NewErrorWrapper(code int, err error, message string, httpStatus int) APIClientException {
	return APIClientException{
		Message:    message,
		HTTPStatus: httpStatus,
		Code:       code,
		Err:        err,
	}
}

func (err APIClientException) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}
