package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
} //添加路由

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
} //添加GET请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
} //添加POST请求
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
} //启动一个http服务
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok { //?
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
