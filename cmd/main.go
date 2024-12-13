package main

import (
	"gin-cli/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	// 加载跨域中间件
	r.Use(middleware.CORSMiddleware())
	//r.Use(middleware.JWTMiddleware())
	r.Static("/static", "template/static")
	r.StaticFS("/pages/", http.Dir("template/pages"))
	r.StaticFile("/", "template/index.html")

	// 启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
