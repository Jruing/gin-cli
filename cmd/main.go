package main

import (
	"gin-cli/app/controller"
	"gin-cli/app/middleware"
	_ "gin-cli/cmd/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	// 加载跨域中间件
	r.Use(middleware.CORSMiddleware())
	//r.Use(middleware.JWTMiddleware())
	r.Static("/static", "template/static")
	r.StaticFS("/pages/", http.Dir("template/pages"))
	r.StaticFile("/", "template/index.html")

	user := r.Group("/user")
	{
		user.GET("/user/all", controller.GetUserDetail)
		user.GET("/user/:id", controller.GetUserDetail)
		user.PUT("/user/:id", controller.UpdateUser)
		user.POST("/user/", controller.CreateUser)
		user.DELETE("/user/:id", controller.DeleteUser)
	}
	role := r.Group("/user")
	{
		role.GET("/role/all", controller.GetRoleDetail)
		role.GET("/role/:id", controller.GetRoleDetail)
		role.PUT("/role/:id", controller.UpdateRole)
		role.POST("/role/", controller.CreateUser)
		role.DELETE("/role/:id", controller.DeleteRole)
	}
	domain := r.Group("/user")
	{
		domain.GET("/domain/all", controller.GetDomainDetail)
		domain.GET("/domain/:id", controller.GetDomainDetail)
		domain.PUT("/domain/:id", controller.UpdateDomain)
		domain.POST("/domain/", controller.CreateDomain)
		domain.DELETE("/domain/:id", controller.DeleteDomain)
	}

	// 启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
