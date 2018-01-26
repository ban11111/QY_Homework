package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"QY_Homework/db/C_D"
	"QY_Homework/db/connection"
	"strings"
	"fmt"
	"QY_Homework/test_configs"
)

func TestDB(t *testing.T) {
	//defer 捕获panic
	defer func() {
		err := recover()
		//判断panic
		assert.Equal(t, strings.Contains(fmt.Sprint(err),"数据库连接出错"),true)
		C_D.DropDB(test_configs.DbConfigforTest())
	}()

	C_D.CreateDB(test_configs.DbConfigforTest())
	testdb := connection.ConnetDB()
	assert.NotEqual(t, testdb, nil, "数据库对象为空")
	C_D.CreateDB(test_configs.DbConfigforTestFail())
	//下面的应该不会运行了
	fmt.Println("要是运行到这里就奇怪了")
	C_D.DropDB(test_configs.DbConfigforTest())

}
