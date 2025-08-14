package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/haoran-mc/golib/internal/handler"
	"github.com/haoran-mc/golib/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valyala/fasthttp"
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

	echoRouter := e.Group("/echo")
	{
		echoRouter.GET("/*", func(c echo.Context) error {
			p := c.Param("*")
			return c.String(http.StatusOK, fmt.Sprintf("[GET] Path: %s\n", p))
		})
		echoRouter.POST("/*", func(c echo.Context) error {
			p := c.Param("*")
			body, _ := io.ReadAll(c.Request().Body)
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

	param := e.Group("/param")
	{
		// 1. 路径参数
		param.GET("/params/:foobar", func(c echo.Context) error {
			foobar := c.Param("foobar")
			return c.String(http.StatusOK, fmt.Sprintf("param(foobar): %s\n", foobar))
		})
		// 2. querymap, http://127.0.0.1:8080/querymap?foobar=xxxx
		param.GET("/querymap", func(c echo.Context) error {
			foobar := c.QueryParam("foobar")
			return c.String(http.StatusOK, fmt.Sprintf("querymap(foobar): %s\n", foobar))
		})
		// 3. 请求体
		param.POST("/body", func(c echo.Context) error {
			type User struct {
				Name string `json:"name"`
				Pass string `json:"pass"`
			}
			var u User
			err := c.Bind(&u)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, u)
		})
		// 4. form 表单
		param.POST("/form", func(c echo.Context) error {
			foobar := c.FormValue("foobar")
			return c.String(http.StatusOK, foobar)
		})
	}

	// 重定向
	redirect := e.Group("/redirect")
	{
		redirect.GET("/hello1", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/redirect/hello2")
		})
		redirect.GET("/hello2", func(c echo.Context) error {
			return c.String(http.StatusOK, "OK")
		})
	}

	// 文件上传
	file := e.Group("/file")
	{
		file.POST("/upload", handler.SingleFileUpload)
		file.POST("/multi", handler.MultiFilesUpload)
	}
	return e
}

func ProxyHandler(ctx *fasthttp.RequestCtx) {
	client := &fasthttp.HostClient{
		Addr: "127.0.0.1:9080",
	}
	req := &ctx.Request

	req.SetHost("127.0.0.1:9080")
	req.URI().SetScheme("http")
	req.URI().SetPath("/echo" + string(req.URI().Path()))

	// 代理请求并接收响应
	var resp fasthttp.Response
	if upstreamErr := client.Do(req, &resp); upstreamErr != nil {
		resp.SetStatusCode(http.StatusBadGateway)
		resp.Header.Set("X-Fake", "true")
		resp.Header.SetContentLength(0)
		resp.ResetBody()
		log.Error("upstream request failed", "url", req.URI().String(), "error", upstreamErr.Error())
		return
	}

	// 响应复制回客户端
	ctx.SetStatusCode(resp.StatusCode())
	ctx.Response.SetBodyRaw(resp.Body())
	resp.Header.CopyTo(&ctx.Response.Header)
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS with ECC cert!\n")
}
