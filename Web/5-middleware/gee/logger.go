package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Fail(500, "Internal Server Error")
		c.Next()
		// Calculate resolution time
		log.Printf("全局日志:  [%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func OnlyForV2() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("V2日志: [%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
