package C_D

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"QY_Homework/db/connection"
	"QY_Homework/db/configs"
)

func CreateDB(config *configs.DbConfig) string {
	openedDb := connection.ConnetDBtoCreate()
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + config.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
	if err := openedDb.Exec(createDbSQL).Error; err != nil {
		return fmt.Sprintf("创建%s数据库失败,%s", config.DbName, err.Error())
	}
	return fmt.Sprintf("数据库“%s”创建成功！！", config.DbName)
}

func DropDB(config *configs.DbConfig) string {
	openedDb := connection.ConnetDB()
	dropDbSQL := "DROP DATABASE IF EXISTS " + config.DbName + ";"
	if err := openedDb.Exec(dropDbSQL).Error; err != nil {
		return fmt.Sprintf("删除“%s”数据库失败,%s", config.DbName, err.Error())
	}
	return fmt.Sprintf("数据库“%s”已被删除！！", config.DbName)
}