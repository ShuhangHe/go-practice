package gee

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]MyHandler
}

func NewRouter() *Router {
	return &Router{make(map[string]MyHandler)}
}

// 添加路由
func (router *Router) addRouter(method string, pattern string, handler MyHandler) bool {
	// 异常处理
	success := true
	defer func() {
		if err := recover(); err != nil {
			success = false
		}
	}()
	router.handlers[method+"-"+pattern] = handler
	return success
}

func (router *Router) Handler(context *Context) {
	key := context.Method + "-" + context.Path
	log.Printf("router: %v", key)
	if handler, ok := router.handlers[key]; ok {
		handler(context)
	} else {
		err := context.String(http.StatusNotFound, "404 Not Found %s", key)
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	}
}
