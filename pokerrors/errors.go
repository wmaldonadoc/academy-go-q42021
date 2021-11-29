package pokerrors

//	APIClientError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type APIClientError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}

//	BuiltinFunctionError - Represents exceptions with the following fields:
//	- Message
//	- HTTPStatus
//	- Code: Represents a custom error code
//	- Err: An Error from errors package
type BuiltinFunctionError struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
	Code       int    `json:"-"`
	Err        error  `json:"-"`
}
