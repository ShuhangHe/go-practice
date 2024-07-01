package gee

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]MyHandler
}

func NewRouter() *Router {
	return &Router{make(map[string]*node), make(map[string]MyHandler)}
}

func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, i := range split {
		// 不为空则加入parts中
		if i != "" {
			parts = append(parts, i)
			// 如果当前index的首个字符是*说明后续可以全匹配
			if i[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由
func (router *Router) addRouter(method string, pattern string, handler MyHandler) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	// 如果不存在当前method的路由在，则初始化
	if _, ok := router.roots[method]; !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	router.handlers[key] = handler
}

func (router *Router) getRouter(method string, pattern string) (*node, map[string]string) {
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	// 当前路径的分片
	parts := parsePattern(pattern)
	params := make(map[string]string)
	nd := root.search(parts, 0)
	if nd != nil {
		// 匹配到的规则分片
		routerParts := parsePattern(nd.pattern)
		for index, part := range routerParts {
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}
			if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return nd, params
	}
	return nil, nil
}

func (router *Router) Handler(context *Context) {
	nd, params := router.getRouter(context.Method, context.Path)
	key := context.Method + "-" + context.Path
	if nd != nil {
		log.Printf("router: %v", key)
		context.Params = params
		router.handlers[context.Method+"-"+nd.pattern](context)
	} else {
		err := context.String(http.StatusNotFound, "404 Not Found %s", key)
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	}
}
