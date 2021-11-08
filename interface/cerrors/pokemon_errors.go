package cerrors

type RequestError struct {
	Message  string
	HttpCode int
	Code     int
}

func PokemonNotFoundException() *RequestError {
	return &RequestError{
		Message:  "Pokemon not found.",
		HttpCode: 404,
		Code:     1000,
	}
}

func ParseTypesException(source string, target string) *RequestError {
	return &RequestError{
		Message:  "Error parsing data from " + source + " to " + target,
		HttpCode: 422,
		Code:     1001,
	}
}
