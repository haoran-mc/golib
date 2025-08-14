package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/haoran-mc/golib/internal/server"
	_ "github.com/haoran-mc/golib/pkg/log"
	pkghttp "github.com/haoran-mc/golib/pkg/server/http"
	"github.com/haoran-mc/golib/pkg/timeutil"
)

func main() {
	nowSolarDate := time.Now().Format("20060102")
	fmt.Printf("==> today solar date: %s, today lunar date: %s\n", nowSolarDate, timeutil.Lunar(nowSolarDate))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		pkghttp.Run(ctx, server.NewServerHTTP(), ":9080")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		pkghttp.Proxy(ctx, ":9088", server.ProxyHandler)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// curl --cacert certs/ec_cert.crt https://localhost:9443
		pkghttp.RunWithTls(ctx, ":9443", http.HandlerFunc(server.HTTPHandler), "certs/ec_cert.crt", "certs/ec_private.key")
	}()

	<-ctx.Done()
	wg.Wait() // wait for all servers to shutdown
}
