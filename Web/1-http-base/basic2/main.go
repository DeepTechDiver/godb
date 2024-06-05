package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}

/**

解读:
定义了一个空结构体: Engine 实现了一个方法: ServeHTTP
ServeHTTP () 第一个参数 可以利用ResponseWriter来构造针对该请求的响应  第二个参数 Request 包含了该HTTP请求的所有的信息（请求地址、Header和Body等信息）


实现Web框架的第一步，即，将所有的HTTP请求转向了我们自己的处理逻辑

*/
