package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

/**
* 第一、首先定义了类型HandlerFunc，用来定义路由映射的处理方法。
* 第二、添加了一张路由映射表router，key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。

封装了一个 addRoute 方法 用于绑定请求方法是POST GET 等等。当用户调用(*Engine).GET()方法时，会将路由（GET-/xxx）和处理方法（HandlerFunc）注册到映射表 router 中。即 形成映射关系 代码对应：engine.router[key] = handler
封装了一个 Run 方法 (*Engine).Run()方法，是 ListenAndServe 的包装。

Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径（对应代码：key := req.Method + "-" + req.URL.Path） ，查找路由映射表，如果查到，就执行注册的处理方法（通过路由映射表engine.router[key]找到对应的处理函数handler(w, req)）。如果查不到，就返回 404 NOT FOUND 。
**/
