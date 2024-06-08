package main

import (
	"gofw/Web/5-middleware/gee"
	"net/http"
)

// 通过不同的分组控制
// 实现了 根路由使用了日志中间件 子路由v2使用了测试（onlyForV2）中间件
func main() {
	r := gee.New()

	// 全局中间件
	r.Use(gee.Logger())

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	// v2分组 子中间件
	v2.Use(gee.OnlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
