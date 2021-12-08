package controller

// Context - Abstract the methods of a HTTP framework.
// In this case is Gin.
type Context interface {
	JSON(code int, i interface{})
	Bind(i interface{}) error
	Param(key string) string
	AbortWithStatusJSON(code int, i interface{})
	AbortWithStatus(code int)
	ShouldBindQuery(obj interface{}) error
	BindUri(obj interface{}) error
}
