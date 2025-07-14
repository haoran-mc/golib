package main

import (
	"fmt"
	"time"

	"github.com/haoran-mc/golib/internal/server"
	_ "github.com/haoran-mc/golib/pkg/log"
	"github.com/haoran-mc/golib/pkg/server/http"
	"github.com/haoran-mc/golib/pkg/timeutil"
)

func main() {
	nowSolarDate := time.Now().Format("20060102")
	fmt.Printf("==> today solar date: %s, today lunar date: %s\n", nowSolarDate, timeutil.Lunar(nowSolarDate))

	go http.Run(server.NewServerHTTP(), ":9520")

	http.Proxy(":9088", server.ProxyHandler)
}
