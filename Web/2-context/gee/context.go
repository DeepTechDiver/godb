package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 给map[string]interface{}起了一个别名gee.H
type H map[string]interface{}

// Context Context目前只包含了 请求响应 和 请求内容，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
type Context struct {
	// origin objects
	Writer http.ResponseWriter //构造响应
	Req    *http.Request       //请求内容
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

/**代码解读

第一步、给map[string]interface{}起了一个别名gee.H

第二步、提供了访问Query和PostForm参数的方法。

第三步、提供了快速构造String/Data/JSON/HTML响应的方法。

**/
