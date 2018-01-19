package router

import (
	"github.com/gin-gonic/gin"
	. "QY_Homework/handler"
	"QY_Homework/service"
)

func Start_Server() {
	router := gin.Default()   // routes
	// 将网页路径/public,映射到文件路径/var/www
	router.Static(service.PublicURL, service.PublicPath)
	// 请求routes
	{
		v1 := router.Group("/v1")
		// 创建demo oder
		v1.POST("/demo", Create_demo_Handler)
		// 更新
		v1.PUT("/demo/:id", Update_demo_Handler)
		// 获取详情
		//v1.GET("/users", Get_demoinfo_Handler)
		// 获取列表
		//v1.GET("/users", Get_demo_Handler)
	}
	// 默认启动在8080
	router.Run(":8080")
}