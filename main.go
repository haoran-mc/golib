package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/haoran-mc/golib/internal/server"
	_ "github.com/haoran-mc/golib/pkg/log"
	pkghttp "github.com/haoran-mc/golib/pkg/server/http"
	"github.com/haoran-mc/golib/pkg/timeutil"
)

func main() {
	nowSolarDate := time.Now().Format("20060102")
	fmt.Printf("==> today solar date: %s, today lunar date: %s\n", nowSolarDate, timeutil.Lunar(nowSolarDate))

	go pkghttp.Run(server.NewServerHTTP(), ":9520")

	go pkghttp.Proxy(":9088", server.ProxyHandler)

	// curl --cacert certs/ec_cert.crt https://localhost:9443
	go pkghttp.RunWithTls(":9443", http.HandlerFunc(server.HttpHandler), "certs/ec_cert.crt", "certs/ec_private.key")

	select {}
}
