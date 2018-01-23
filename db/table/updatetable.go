package table

import (
	"github.com/jinzhu/gorm"
	"QY_Homework/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//新建或更新表字段
func DemoTableUpdate(db *gorm.DB,) {
	if db.HasTable(&model.Demo_order{}) == false {
		db.CreateTable(&model.Demo_order{})
	} else {
		db.AutoMigrate(&model.Demo_order{})
	}
}

func FilesTableUpdate(db *gorm.DB) {
	if db.HasTable(&model.Files{}) == false {
		db.CreateTable(&model.Files{})
	} else {
		db.AutoMigrate(&model.Files{})
	}
}