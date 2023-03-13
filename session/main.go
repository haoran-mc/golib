package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haoran-mc/go_pkgs/session/session_store"
)

type User struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

const (
	name string = "name"
	pass string = "pass"
)

func main() {
	session_store.InitCookieStore()
	session_store.InitFilesystemStore()

	engine := gin.Default()

	engine.POST("/user", func(ctx *gin.Context) {
		user := User{}
		ctx.ShouldBindJSON(&user)

		session, _ := session_store.GetFileSystemSession(ctx.Request)
		session.Values[name] = user.Name
		session.Values[pass] = user.Pass
		session.Save(ctx.Request, ctx.Writer)

		ctx.JSON(http.StatusOK, gin.H{
			"message":    "ok",
			"session.ID": session.ID, // 仅在使用 FilesystemStore 时生成
		})
	})

	engine.GET("/user", func(ctx *gin.Context) {
		session, _ := session_store.GetFileSystemSession(ctx.Request)
		name := session.Values[name]
		pass := session.Values[pass]
		data := session.Values["data"]

		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"name":    name,
			"pass":    pass,
			"data":    data,
			"id":      session.ID,
		})
	})

	engine.Run(":8080")
}
