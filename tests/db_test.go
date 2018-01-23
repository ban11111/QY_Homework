package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"QY_Homework/db/C_D"
	"QY_Homework/db/config"
	"QY_Homework/db/connection"
	"strings"
	"fmt"
)

//正确的配置
func DbConfigforTest() *config.DbConfig {
	pwd := "root"
	uname := "root"
	dbname := "testforfun"
	conf := &config.DbConfig{Username: uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}
//错误的配置
func DbConfigforTestFail() *config.DbConfig {
	pwd := "wrongpwd"
	uname := "root"
	dbname := "testfail"
	conf := &config.DbConfig{Username: uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}

func TestDB(t *testing.T) {
	//defer 捕获panic
	defer func() {
		err := recover()
		//判断panic
		assert.Equal(t, strings.Contains(fmt.Sprint(err),"数据库连接出错"),true)
		C_D.DropDB(DbConfigforTest())
	}()

	C_D.CreateDB(DbConfigforTest())
	testdb := connection.ConnetDB(DbConfigforTest())
	assert.NotEqual(t, testdb, nil, "数据库对象为空")
	C_D.CreateDB(DbConfigforTestFail())
	//下面的应该不会运行了
	fmt.Println("要是运行到这里就奇怪了")
	C_D.DropDB(DbConfigforTest())

}
