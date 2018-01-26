package tests

import (
	"github.com/stretchr/testify/suite"
	"fmt"
	"testing"
	"time"
	"QY_Homework/router"
	"QY_Homework/service"
	"github.com/gavv/httpexpect"
	"os"
	"QY_Homework/test_configs"
)

type XLSXTestSuite struct {
	suite.Suite
}

func TestSuiteXLSX(t *testing.T) {
	//logger.InitLogger(logger.LvlDebug,nil)
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	suite.Run(t, new(XLSXTestSuite))
	os.Remove("/home/qydev/var/www/xlsx/Demo_Order.xlsx")
}

func (s *XLSXTestSuite) TestXLSX() {
	err := service.GetXLSX()
	s.Equal(nil, err, "文档生成失败!!")
	fmt.Println("文档生成成功！")
}

func (s *XLSXTestSuite)TestGetXLSX() {
	resp := httpexpect.New(s.T(), test_configs.TestServer).GET(test_configs.Getdemoxlsx).Expect()
	resp.Status(200)
	fmt.Println(resp.Headers())
}