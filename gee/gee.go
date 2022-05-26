package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
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
	/*
		key := req.Method + "-" + req.URL.Path
		if handler, ok := engine.router[key]; ok { //?
			handler(w, req)
		} else {
			fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		}
	*/
	c := newContext(w, req)
	engine.router.handle(c)
}
