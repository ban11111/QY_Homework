package main

import (
	"QY_Homework/db/config"
	. "QY_Homework/db/connection"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func CreateDB(config *config.DbConfig) string {
	openedDb := ConnetDBtoCreate(config)
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + config.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
	if err := openedDb.Exec(createDbSQL).Error; err != nil {
		return fmt.Sprintf("创建%s数据库失败,%s", config.DbName, err.Error())
	}
	return fmt.Sprintf("数据库“%s”创建成功！！", config.DbName)
}

//func CreateTable()

func main() {
	//根据配置，创建数据库
	message := CreateDB(config.NewDbConfig())
	fmt.Println(message)
	//message = DropDB(config.NewDbConfig())
	//fmt.Println(message)
}




func DropDB(config *config.DbConfig) string {
	openedDb := ConnetDB(config)
	dropDbSQL := "DROP DATABASE IF EXISTS " + config.DbName + ";"
	if err := openedDb.Exec(dropDbSQL).Error; err != nil {
		return fmt.Sprintf("删除“%s”数据库失败,%s", config.DbName, err.Error())
	}
	return fmt.Sprintf("数据库“%s”已被删除！！", config.DbName)
}
