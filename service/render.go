package service

import (
	"github.com/gin-gonic/gin"
	"QY_Homework/model"
)

type BaseResp struct {
	Success bool			`json:"success"`
	Info string				`json:"info"`
}

type SuccessResp struct {
	BaseResp
	*model.Demo_order
}

func Render200(data interface{}, c *gin.Context) {
	c.JSON(200, data)
}

func Render400(info string, c *gin.Context) {
	c.JSON(400, BaseResp{Success: false, Info: info})
}