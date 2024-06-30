package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyMap map[string]interface{}

type Context struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

// NewContext 初始化方法
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
		Path:     r.URL.Path,
		Method:   r.Method,
	}
}

func (context *Context) PostFrom(key string) string {
	return context.Request.FormValue(key)
}

func (context *Context) Query(key string) string {
	return context.Request.URL.Query().Get(key)
}

func (context *Context) Status(code int) {
	context.StatusCode = code
	context.Response.WriteHeader(code)
}

func (context *Context) SetHeader(key string, value string) {
	context.Response.Header().Set(key, value)
}

func (context *Context) String(code int, format string, values ...interface{}) error {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	_, err := context.Response.Write([]byte(fmt.Sprintf(format, values...)))
	return err
}

func (context *Context) Json(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json; charset=utf-8")
	context.Status(code)
	encoder := json.NewEncoder(context.Response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Response, err.Error(), 500)
	}
}

func (context *Context) Data(code int, data []byte) error {
	context.Status(code)
	_, err := context.Response.Write(data)
	return err
}

func (context *Context) Html(code int, html string) error {
	context.SetHeader("Content-Type", "text/html; charset=utf-8")
	context.Status(code)
	_, err := context.Response.Write([]byte(html))
	return err
}

func (context *Context) Error(code int, err error) {
	http.Error(context.Response, err.Error(), code)
}
