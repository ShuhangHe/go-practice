package gee

import (
	"net/http"
)

// MyHandler 定义方法类型
type MyHandler func(context *Context)

// Engine 定义结构体
type Engine struct {
	// 路由表，key为请求路径，value为具体的处理方法
	router *Router
}

// New 初始化方法
func New() *Engine {
	return &Engine{router: NewRouter()}
}

// 添加路由
func (engine *Engine) addRouter(method string, pattern string, handler MyHandler) bool {
	// 异常处理
	return engine.router.addRouter(method, pattern, handler)
}

// GET 类型 添加路由规则
func (engine *Engine) GET(pattern string, handler MyHandler) bool {
	return engine.addRouter("GET", pattern, handler)
}

// POST 类型添加路由规则
func (engine *Engine) POST(pattern string, handler MyHandler) bool {
	return engine.addRouter("POST", pattern, handler)
}

// 启动方法
// addr 填写为   :+端口号
// eg ":8080"
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.Handler(NewContext(w, req))
}
