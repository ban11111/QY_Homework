package test_configs

import (
	"QY_Homework/db/configs"
)

const(
	TestServer = "http://localhost:8080"
	Version = "/v1"
	Createdemo = Version + "/demo"
	Updatedemo = Version + "/demo/1"
	Getdemoinfo = Updatedemo
	Postdemolist = Version + "/demo-all"
	Getdemolist = Postdemolist
	Getdemoxlsx = Version + "/demo-xlsx"
	Transction = Version + "/demo/2"
)

//正确的配置
func DbConfigforTest() *configs.DbConfig {
	pwd := "root"
	uname := "root"
	dbname := "testforfun"
	conf := &configs.DbConfig{Username: uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}

//错误的配置
func DbConfigforTestFail() *configs.DbConfig {
	pwd := "wrongpwd"
	uname := "root"
	dbname := "testfail"
	conf := &configs.DbConfig{Username: uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}