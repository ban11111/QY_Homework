package main

import (
	"./config"
	"fmt"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func check_config(config *config.DbConfig) error{
	var err error
	if len(config.DbName) == 0 {
		err = errors.New("db_name 不能为空")
		return err
	}
	if len(config.Username) == 0 {
		err = errors.New("username 不能为空")
		return err
	}
	if len(config.Password) == 0 {
		err = errors.New("password 不能为空")
		return err
	}
	if len(config.Host) == 0 {
		err = errors.New("host 不能为空")
		return err
	}
	if len(config.Port) == 0 {
		err = errors.New("port 不能为空")
		return err
	}
	return nil
}

func CreateDB(config *config.DbConfig) {
	if err := check_config(config) ;err != nil {
		panic("数据库连接配置不正确" + err.Error())
	}
	openedDb, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, "information_schema"))
	if err == nil {
		panic("数据库连接出错" + err.Error())
	}
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + config.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
	if err := openedDb.Exec(createDbSQL).Error;err != nil {
		fmt.Printf("创建%s数据库失败,%s",config.DbName,err.Error())
	}
}