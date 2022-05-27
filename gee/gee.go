package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)
type (
	Engine struct {
		router *router
		*RouterGroup
		groups []*RouterGroup //store all groups
	}

	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc //support middleware
		parent      *RouterGroup  //support nesting
		engine      *Engine       //all groups share a Engine instance
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
} //添加分组

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
} //添加路由

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
} //添加GET请求

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
