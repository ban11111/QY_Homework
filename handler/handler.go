package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "QY_Homework/db/connection"
	. "QY_Homework/db/config"
	"QY_Homework/model"
	//"QY_Homework/db/table"
	"QY_Homework/service"
	"fmt"
	"QY_Homework/db/table"
)

func Create_demo_Handler(c *gin.Context) {
	var newdemo *model.Demo_order
	//if err := c.BindJSON(demo); err != nil {
	//	service.Render400(err.Error(), c)
	//	return
	//}
	openedDb := ConnetDB(NewDbConfig())
	//新建或更新表字段
	table.TableUpdate(openedDb)
	if form, err := c.MultipartForm(); err != nil {
		service.Render400(err.Error(), c)
		return
	} else {
		fmt.Println("表单数据",form.Value)
		newdemo, err = Transfer_form_to_model(form.Value)
		if err != nil {
			service.Render400(err.Error(), c)
		}
		//id自増
		SelfIncrease(openedDb, newdemo)
		fmt.Println("转换成模型",newdemo)
		if err = newdemo.IsValid(); err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//添加数据
		if err = service.Create_demo(openedDb, newdemo); err != nil {
			service.Render400(err.Error(), c)
			return
		}
		tmpresp := service.BaseResp{true, "添加成功！"}
		service.Render200(&service.SuccessResp{
			BaseResp:   tmpresp,
			Demo_order: newdemo,
		}, c)
	}
}

func Update_demo_Handler(c *gin.Context) {
	id := c.Param("id")


}

func Get_demoinfo_Handler(c *gin.Context) {

}

func Get_demo_Handler(c *gin.Context) {

}

// 从表单中获取用户
func ValueFromMultipartForm(key string, f map[string][]string) string {
	if len(f[key]) > 0 {
		return f[key][0]
	} else {
		return ""
	}
}
// 将表单转换成模型对象
func Transfer_form_to_model(f map[string][]string) (*model.Demo_order, error) {
	str := ValueFromMultipartForm("amount", f)
	//fmt.Println(str)
	if amount, err := strconv.ParseFloat(str, 64); err == nil {
		var tmp *model.Demo_order
		tmp = &model.Demo_order{
			Order_id:  ValueFromMultipartForm("order_id", f),
			User_name: ValueFromMultipartForm("user_name", f),
			Amount:    amount,
			Status:    ValueFromMultipartForm("status", f),
		}
		//fmt.Println(tmp.Order_id)
		return tmp, nil
	} else {
		//fmt.Println(err)
		return nil, err
	}
}
