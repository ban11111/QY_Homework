package config

import (
	"github.com/jinzhu/gorm"
	"QY_Homework/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type DbConfig struct {
	// 数据库用户名
	Username string
	// 密码
	Password string
	Host     string
	Port     string
	// 数据库名
	DbName   string
	// 最大闲置连接
	MaxIdleConns int
	// 最大打开连接
	MaxOpenConns int
}

//数据库配置
func NewDbConfig() *DbConfig {
	pwd := "root"
	uname := "root"
	dbname := "homework"
	conf := &DbConfig{Username: uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}

func SelfIncrease(db *gorm.DB, demo *model.Demo_order) {
	var id_value model.Demo_order
	db.Last(&id_value)
	demo.Id = id_value.Id + 1
}