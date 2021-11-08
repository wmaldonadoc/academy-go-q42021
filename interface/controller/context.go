package controller

type Context interface {
	JSON(code int, i interface{})
	Bind(i interface{}) error
	Param(key string) string
}
