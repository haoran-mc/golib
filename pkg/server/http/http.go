package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Run(ctx context.Context, e *echo.Echo, addr string) {
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

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		panic(fmt.Sprintf("server shutdown failed: %s\n", err))
	}
	fmt.Println("server exiting...")
}

func Proxy(ctx context.Context, addr string, handler fasthttp.RequestHandler) {
	srv := &fasthttp.Server{
		Handler: handler,
	}

	go func() {
		fmt.Println("==> proxy addr: " + addr)
		if err := srv.ListenAndServe(addr); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("proxy start failed: %s\n", err))
		}
	}()

	<-ctx.Done()

	// fasthttp.Server.Shutdown() does not support context with timeout
	if err := srv.Shutdown(); err != nil {
		panic(fmt.Sprintf("proxy shutdown failed: %s\n", err))
	}
	fmt.Println("proxy exiting...")
}

func RunWithTls(ctx context.Context, addr string, handler http.Handler, certFile, keyFile string) {
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

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := tlsSrv.Shutdown(shutdownCtx); err != nil {
		panic(fmt.Sprintf("server shutdown failed: %s\n", err))
	}
	fmt.Println("server(tls) exiting...")
}
