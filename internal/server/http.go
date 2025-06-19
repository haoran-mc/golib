package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/haoran-mc/golib/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServerHTTP() *echo.Echo {
	e := echo.New()
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 90 * time.Second
	e.Use(middleware.Decompress())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${remote_ip} ${method} ${uri} ${status} ${latency_human}\n",
	}))

	e.GET("/", handler.Hello)

	echoRouter := e.Group("/")
	{
		echoRouter.GET("/*", func(c echo.Context) error {
			p := c.Param("*")
			return c.String(http.StatusOK, fmt.Sprintf("[GET] Path: %s\n", p))
		})

		echoRouter.POST("/*", func(c echo.Context) error {
			p := c.Param("*")
			body, _ := io.ReadAll(c.Request().Body)
			fmt.Println(string(body))
			return c.String(http.StatusOK, fmt.Sprintf("[POST] Path: %s, Data: %s\n", p, string(body)))
		})

		echoRouter.PUT("/*", func(c echo.Context) error {
			p := c.Param("*")
			body, _ := io.ReadAll(c.Request().Body)
			return c.String(http.StatusOK, fmt.Sprintf("[PUT] Path: %s, Data: %s\n", p, string(body)))
		})

		echoRouter.DELETE("/*", func(c echo.Context) error {
			p := c.Param("*")
			body, _ := io.ReadAll(c.Request().Body)
			return c.String(http.StatusOK, fmt.Sprintf("[DELETE] Path: %s, Data: %s\n", p, string(body)))
		})
	}

	redirect := e.Group("/redirect")
	{
		redirect.GET("/hello1", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/redirect/hello2")
		})

		redirect.GET("/hello2", func(c echo.Context) error {
			return c.String(http.StatusOK, "OK")
		})
	}
	return e
}
