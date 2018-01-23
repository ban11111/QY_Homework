package service

import (
	"github.com/tealeg/xlsx"
	"QY_Homework/model"
	"QY_Homework/db/connection"
	"QY_Homework/db/config"
	"reflect"
)

//生成excel文档
func GetXLSX() (err error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//创建新表
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Demo Order")
	if err != nil {
		return
	}
	//创建Title
	var Title model.Demo_order
	title := MakeTitle(Title)
	row = sheet.AddRow()
	row.WriteSlice(&title, -1)
	//写入数据
	var DemoToWrite []model.Demo_order
	openedDB := connection.ConnetDB(config.NewDbConfig())
	openedDB.Find(&DemoToWrite)
	for i := range DemoToWrite {
		row = sheet.AddRow()
		row.WriteStruct(&DemoToWrite[i], -1)
	}
	//保存文档
	err = file.Save(XlsxPath + "Demo_Order.xlsx")
	if err != nil {
		return
	}
	return nil
}

//创建标题
func MakeTitle(Title model.Demo_order) []string {
	title := make([]string, 7, 10)
	type_ := reflect.ValueOf(&Title).Elem().Type()
	for i := 0; i < type_.NumField(); i++ {
		title[i] = type_.Field(i).Name
	}
	return title
}
