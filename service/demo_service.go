package service

import (
	"QY_Homework/model"
	"github.com/jinzhu/gorm"
	"errors"
	"time"
	"strconv"
	"path/filepath"
	"os"
	"strings"
	"QY_Homework/tools"
	"fmt"
)

const (
	PublicPath = /*"F:/Homework_files/"*/ "/home/qydev/var/www/"
	UploadPath = PublicPath + "uploads/"
	XlsxPath   = PublicPath + "xlsx/"
	PublicURL  = "/public/"
	XlsxURL    = PublicURL + "xlsx/"
)

//向db中添加数据
func Create_demo(db *gorm.DB, demo *model.Demo_order) (err error) {
	//创建时间
	demo.CreatedAt = time.Now()
	//事务开始
	TransactionCreateDemo := db.Begin()
	if err = db.Create(demo).Error; err != nil {
		TransactionCreateDemo.Rollback()
		return
	}
	if err = SyncFiles(db, demo); err != nil {
		TransactionCreateDemo.Rollback()
		return
	}
	TransactionCreateDemo.Commit()
	//事务结束
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//因为time.now 的时间精度太高，数据库取出的数据精度只到秒，为方便统一，直接把数据库的值取出来发送到前端（暂时没找到更方便的直接削减精度的方法）
	if err = db.Last(demo).Error; err != nil {
		return
	}
	fmt.Println(demo)
	return
}

//更新demo
func Update_demo(db *gorm.DB, demo *model.Demo_order) (err error) {
	if demo.Id == 0 {
		return errors.New("id不能为空")
	}
	//事务开始
	TransactionUpdateDemo := db.Begin()
	if err = db.Model(demo).Where("id=?", demo.Id).Update(demo).Error; err != nil {
		TransactionUpdateDemo.Rollback()
		return
	}
	if err = SyncFiles(db, demo); err != nil {
		TransactionUpdateDemo.Rollback()
		return
	}
	//更新时，如果files里的id不存在，就会报错。正好用于测试事务。
	if err = Id_exist(db,demo.Id,&model.Files{}); err != nil {
		fmt.Println("FFFFFFFFFFFFFFFFFFF", err)
		TransactionUpdateDemo.Rollback()
		return
	}
	TransactionUpdateDemo.Commit()
	//事务结束
	//
	err = db.Model(demo).Where("id=?", demo.Id).First(demo).Error
	return
}

//判断ID是否存在, 需要先连接数据库
func Id_exist(db *gorm.DB, id uint64, demo interface{}) error {
	if err := db.Where("id = ?", id).First(demo).Error; err != nil {
		return errors.New("数据库中没有该记录，请确认该数据已经建立---"+err.Error())
	}
	return nil
}

//同步数据到files表
func SyncFiles(db *gorm.DB, demo *model.Demo_order) error {
	files := &model.Files{}
	//files.Id = demo.Id
	var i = 0
	temparry := make([]string, 10)
	files.File_path = UploadPath + "f" + strconv.FormatUint(demo.Id, 10) + "/"
	err := filepath.Walk(files.File_path, func(files_path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		_, filename := filepath.Split(files_path)
		if strings.Contains(filename, ".") {
			temparry[i] = filename
			i++
		}
		return nil
	})
	if err != nil {
		return err
	}
	if err = files.FilesIsValid(); err != nil{
		return err
	}
	files.File_path += tools.StringJoin(temparry, ",")
	err = db.Create(files).Error
	return err
}
