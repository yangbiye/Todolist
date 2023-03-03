package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yby/todo_list/api"
	"github.com/yby/todo_list/middleware"
)

// Setup 启动路由
func Setup() {
	// router
	r := gin.Default()

	// v1版 http请求响应
	v1 := r.Group("/api/v1")
	{
		// 用户模块
		v1.POST("/login", api.LoginUser)
		v1.POST("/register", api.RegisterUser)

		// 事务模块
		v1.Use(middleware.JWT()).POST("/create", api.CreateItem)
		v1.Use(middleware.JWT()).DELETE("/delete", api.DeleteItem)
		v1.Use(middleware.JWT()).POST("/update", api.UpdateItem)
		v1.Use(middleware.JWT()).GET("/get", api.GetItem)
	}

	r.Run(":8080")
}
