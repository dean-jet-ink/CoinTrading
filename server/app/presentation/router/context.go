package router

import (
	"mime/multipart"
	"net/http"
)

type Context interface {
	JSON(code int, obj any)
	String(code int, str string)
	NoContent(code int)
	BindQuery(obj any) error
	BindJSON(obj any) error
	Param(key string) string
	FormFile(name string) (*multipart.FileHeader, error)
	Next()
	Writer() http.ResponseWriter
	Request() *http.Request
}
