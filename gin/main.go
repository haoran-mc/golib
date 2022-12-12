package main

import (
	"net/http"

	"github.com/haoran-mc/go_pkgs/gin/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin</h1>")
	})

	r.GET("/hello", func(c *gin.Context) {
		// expect /hello?name=haoran
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		// expect /hello/haoran
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"filepath": c.Param("filepath"),
		})
	})

	_ = r.Run(":9999")
}
