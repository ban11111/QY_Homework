package table

import (
	"github.com/jinzhu/gorm"
	"QY_Homework/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//新建或更新表字段
func TableUpdate(openedDb *gorm.DB,) {
	if openedDb.HasTable(&model.Demo_order{}) == false {
		openedDb.CreateTable(&model.Demo_order{})
	} else {
		openedDb.AutoMigrate(&model.Demo_order{})
	}
}