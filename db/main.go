package main

import (
	"QY_Homework/db/config"
	"fmt"
	"QY_Homework/db/C_D"
)

func main() {
	//根据配置，创建数据库
	message := C_D.CreateDB(config.NewDbConfig())
	fmt.Println(message)
	//message = C_D.DropDB(config.NewDbConfig())
	//fmt.Println(message)
}





