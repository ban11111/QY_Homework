package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "QY_Homework/db/connection"
	"QY_Homework/model"
	"QY_Homework/service"
	"QY_Homework/db/table"
	"fmt"
	"errors"
	"QY_Homework/db/configs"
	"io/ioutil"
)

func CreateDemoHandler(c *gin.Context) {
	var newdemo *model.Demo_order
	//if err := c.BindJSON(demo); err != nil {
	//	service.Render400(err.Error(), c)
	//	return
	//}
	if form, err := c.MultipartForm(); err != nil {
		service.Render400(err.Error(), c)
		return
	} else {
		openedDb := ConnetDB()
		//新建或更新表字段
		table.DemoTableUpdate(openedDb)
		table.FilesTableUpdate(openedDb)
		//fmt.Println("表单数据", form.Value)
		newdemo, err = Transfer_form_to_model(form.Value)
		if err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//id自増
		configs.SelfIncrease(openedDb, newdemo)
		//上传文件（可以上传多个文件）,需要放在id生成之后
		newdemo.File_url, err = service.Uploadfiles(form.File["file_url"], newdemo.Id, c)
		if err != nil {
			service.Render400(err.Error(), c)
			return
		}
		if err = newdemo.IsValid(); err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//添加数据
		if err = service.Create_demo(openedDb, newdemo); err != nil {
			service.Render400(err.Error(), c)
			return
		}
		respinfo := service.BaseResp{true, "添加成功！"}
		service.Render200(service.SuccessResp{
		BaseResp: respinfo,
		Demo_order: *newdemo,
		}, c)
	}
}

func UpdatedemoHandler(c *gin.Context) {
	var updatedemo *model.Demo_order
	if form, err := c.MultipartForm(); err != nil {
		service.Render400(err.Error(), c)
		return
	} else {
		openedDb := ConnetDB()
		updatedemo, err = Transfer_form_to_model(form.Value)
		//fmt.Println("更新前",updatedemo)
		//获取ID
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//fmt.Println("更新前中",updatedemo)
		updatedemo.Id = id
		//判断ID 是否已经存在，若不存在就报错,注意一定要传空值的模型进去
		if openedDb.Where("id = ?", id).First(&model.Demo_order{}).RecordNotFound() {
			service.Render400(errors.New("该条记录不存在!不能更新!").Error(), c)
			return
		}
		//fmt.Println("更新中",updatedemo)
		//更新文件
		updatedemo.File_url, err = service.Updatefiles(form.File["file_url"], updatedemo.Id, c)
		if err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//更新demo
		if err := service.Update_demo(openedDb, updatedemo); err != nil {
			service.Render400(err.Error(), c)
			return
		}
		//fmt.Println("更新后",updatedemo)
		respinfo := service.BaseResp{true, "更新成功！"}
		service.Render200(service.SuccessResp{
			BaseResp:   	respinfo,
			Demo_order: 	*updatedemo,
		}, c)
	}
}

//查询单条记录
func GetDemoInfoHandler(c *gin.Context) {
	var DBdata model.Demo_order
	id := c.Param("id")
	idu,_ := strconv.ParseUint(id, 10, 64)
	openedDb := ConnetDB()
	//如果id不存在,则返回错误
	if openedDb.Where("id = ?", idu).First(&DBdata).RecordNotFound() {
		service.Render400(fmt.Sprintf("查询不到id为'%d'的数据", id), c)
		return
	}
	respinfo := service.BaseResp{true, "查询成功！"}
	service.Render200(&service.SuccessResp{
		BaseResp:   respinfo,
		Demo_order: DBdata,
	}, c)
}

//GET 查询所有记录 按id主键排序
func GetDemoListHandler(c *gin.Context) {
	var DBdata []model.Demo_order
	openedDb := ConnetDB()
	openedDb.Find(&DBdata)
	respinfo := service.BaseResp{true, "查询成功！"}
	service.Render200(&service.ListResp{
		BaseResp: respinfo,
		List:     DBdata,
	}, c)
}

//POST 带条件 查询所有记录
func PostDemoListHandler(c *gin.Context) {
	var err error
	var search, sortby string
	//json
	var jsonmap map[string]string
	if err := c.BindJSON(&jsonmap); err != nil {
		service.Render400(err.Error(), c)
		return
	}
	search = jsonmap["search"]
	sortby = jsonmap["sortby"]
	var SORT string
	switch sortby {
	case "time":
		SORT = "created_at"
	case "timedesc":
		SORT = "created_at desc"
	case "amount":
		SORT = "amount"
	case "amountdesc":
		SORT = "amount desc"
	default:
		service.Render400("前端post数据有误", c)
	}
	var DBdata []model.Demo_order
	openedDb := ConnetDB()
	if search != "" {
		search = "%" + search + "%"
		err = openedDb.Where("user_name LIKE ?", search).Order(SORT).Find(&DBdata).Error
	}else{
		err = openedDb.Order(sortby).Find(&DBdata).Error
	}
	if err != nil {
		service.Render400(err.Error(), c)
	}
	respinfo := service.BaseResp{true, "查询成功！"}
	service.Render200(&service.ListResp{
		BaseResp: respinfo,
		List:     DBdata,
	}, c)
}

//生成xlsx，并redirect到下载地址
func GetDemoXlsxHandler(c *gin.Context) {
	if err := service.GetXLSX(); err != nil {
		service.Render400(err.Error(), c)
		return
	}
	//respinfo := service.BaseResp{true, "xlsx成功生成！"}
	//service.Render200(&service.SuccessResp{
	//	BaseResp:   respinfo,
	//}, c)
	//c.File()
	//c.File(service.XlsxPath + "Demo_Order.xlsx")
	c.Writer.Header().Set("content-type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=123.xlsx")
	data, err := ioutil.ReadFile(service.XlsxPath + "Demo_Order.xlsx")
	if err != nil {
		service.Render400(err.Error(), c)
		return
	}
	c.Writer.Write(data)
}

// 从Multipart表单中获取用户
func ValueFromMultipartForm(key string, f map[string][]string) string {
	if len(f[key]) > 0 {
		return f[key][0]
	} else {
		return ""
	}
}

// 将Multipart表单转换成模型对象
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
