package router

import (
	"github.com/gin-gonic/gin"
	. "../handler"
)

const PublicPath = "/home/qydev/var/www"

func Start_Server() {
	router := gin.Default()   // routes
	// 将网页路径/public,映射到文件路径/var/www
	router.Static("/public/", PublicPath)
	// 请求routes
	{
		v1 := router.Group("/v1")
		// 创建demo oder
		v1.POST("/demo", Create_demo_Handler)
		// 更新
		v1.PUT("/demo/:id", Update_demo_Handler)
		// 获取详情
		v1.GET("/users", Get_demoinfo_Handler)
		// 获取列表
		v1.GET("/users", Get_demo_Handler)
	}
	// 默认启动在8080
	router.Run(":8080")
}