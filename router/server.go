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
		v1.POST("/demo", CreateDemoHandler)
		// 更新
		v1.PUT("/demo/:id", UpdatedemoHandler)
		// 获取详情
		v1.GET("/demo/:id", GetDemoInfoHandler)
		// 获取列表
		v1.GET("/demo-all", GetDemoListHandler)
		v1.POST("/demo-all", PostDemoListHandler)
		//下载xlsx文档
		v1.GET("/demo-xlsx", GetDemoXlsxHandler)
	}
	// 默认启动在8080
	router.Run(":8080")
}