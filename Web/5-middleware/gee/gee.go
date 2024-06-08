package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
// 嵌套分组
type (
	RouterGroup struct {
		prefix      string        // 前缀
		middlewares []HandlerFunc // 中间件
		parent      *RouterGroup  // 分组的父亲(parent) 实现嵌套功能
		engine      *Engine       // 所有Group共享一个Engine实例
	}

	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 被定义用来创建一个新的 RouterGroup
// 所有的组共享同一个 Engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

/**
** 调用了group.engine.router.addRoute来实现了路由的映射。
** 由于Engine从某种意义上继承了RouterGroup的所有属性和方法，因为 (*Engine).engine 是指向自己的。
** 这样实现，我们既可以像原来一样添加路由，也可以通过分组添加路由。
**/
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 当接收到一个具体请求时
	// 要判断该请求适用于哪些中间件，通过 URL 的前缀来判断
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			// 中间件列表 [A, B, Handler]
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	// 得到中间件列表后，赋值给 c.handlers
	c.handlers = middlewares
	engine.router.handle(c)
}
