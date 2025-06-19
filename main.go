package main

import (
	"github.com/haoran-mc/golib/internal/server"
	_ "github.com/haoran-mc/golib/pkg/log"
	"github.com/haoran-mc/golib/pkg/server/http"
)

func main() {
	http.Run(server.NewServerHTTP(), ":9520")
}
