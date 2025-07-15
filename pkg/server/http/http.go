package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Run(e *echo.Echo, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}
	go func() {
		fmt.Println("==> server addr: " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("server start failed: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("server shutdown failed: %s\n", err))
	}
	fmt.Println("server exiting")
}

func Proxy(addr string, handler fasthttp.RequestHandler) {
	srv := &fasthttp.Server{
		Handler: handler,
	}

	go func() {
		fmt.Println("==> proxy addr: " + addr)
		if err := srv.ListenAndServe(addr); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("proxy start failed: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down proxy...")

	// fasthttp.Server.Shutdown() does not support context with timeout
	if err := srv.Shutdown(); err != nil {
		panic(fmt.Sprintf("proxy shutdown failed: %s\n", err))
	}
	fmt.Println("proxy exiting")
}

func RunWithTls(addr string, handler http.Handler, certFile, keyFile string) {
	tlsSrv := &http.Server{
		Addr:    addr,
		Handler: handler,
		// TLSConfig: &tls.Config{
		// MinVersion: tls.VersionTLS12, // 强制禁用 TLS 1.0/1.1
		// },
	}

	go func() {
		fmt.Println("==> server(tls) addr: " + tlsSrv.Addr)
		if err := tlsSrv.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("server start failed: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := tlsSrv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("server shutdown failed: %s\n", err))
	}
	fmt.Println("server exiting")
}
