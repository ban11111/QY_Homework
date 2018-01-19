package service

import (
	"QY_Homework/model"
	"github.com/jinzhu/gorm"
	"errors"
)

//向db中添加数据
func Create_demo(db *gorm.DB, demo *model.Demo_order) (err error){
	err = db.Create(demo).Error
	return
}

//更新demo
func Update_demo(db *gorm.DB, demo *model.Demo_order) error{
	if demo.Id == 0 {
		return errors.New("id不能为空")
	}
	err := db.Model(demo).Where("id=?", demo.Id).Update(demo).Error
	return err
}