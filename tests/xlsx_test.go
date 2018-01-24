package tests

import (
	"github.com/stretchr/testify/suite"
	"fmt"
	"testing"
	"time"
	"QY_Homework/router"
	"QY_Homework/service"
	"github.com/gavv/httpexpect"
)

type XLSXTestSuite struct {
	suite.Suite
}

//func TearDownTest(s *XLSXTestSuite) {
//	fmt.Print("测试结束")
//}

func TestSuiteXLSX(t *testing.T) {
	//logger.InitLogger(logger.LvlDebug,nil)
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	suite.Run(t, new(XLSXTestSuite))
}

func (s *XLSXTestSuite) TestXLSX() {
	err := service.GetXLSX()
	s.Equal(nil, err, "文档生成失败!!")
	fmt.Println("文档生成成功！")
}

func (s *XLSXTestSuite)TestGetXLSX() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemoxlsx).Expect()
	resp.Status(200)
	fmt.Println(resp.Headers())
}