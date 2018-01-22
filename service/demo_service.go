package service

import (
	"QY_Homework/model"
	"github.com/jinzhu/gorm"
	"errors"
	"time"
)

const (
	PublicPath = /*"F:/Homework_files/"*/"/home/qydev/var/www/"
	UploadPath = PublicPath + "uploads/"
	XlsxPath = PublicPath + "xlsx/"
	PublicURL = "/public/"
	XlsxURL = PublicURL + "xlsx/"
)


//向db中添加数据
func Create_demo(db *gorm.DB, demo *model.Demo_order) (err error){
	//创建时间
	demo.CreatedAt = time.Now()
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

//判断ID是否存在, 需要先连接数据库
func Id_exist(db *gorm.DB, demo *model.Demo_order) error {
	if db.Where("id = ?",demo.Id).First(demo) == nil {
		return errors.New("数据库中没有该记录，请确认该数据已经建立")
	}
	return nil
}