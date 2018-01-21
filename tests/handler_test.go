package tests

import (
	"os"
	"testing"
	. "QY_Homework/db/config"
	"github.com/json-iterator/go"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/gavv/httpexpect"
	"QY_Homework/service"
	"QY_Homework/router"
	"time"
	"QY_Homework/db/connection"
	"QY_Homework/model"
	//"gcoresys/common/logger"
	"fmt"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HandlerTestSuite struct {
	suite.Suite
}

func TearDownTest(s *HandlerTestSuite) {
	fmt.Print("测试结束")
}

func TestSuite(t *testing.T) {
	//logger.InitLogger(logger.LvlDebug,nil)
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	suite.Run(t, new(HandlerTestSuite))
}

//新建demo_order 测试(无文件)
func (s *HandlerTestSuite)Testcreatedemowithoutfile() {
	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithForm(map[string]string{
		"order_id"   : "xx1280a",
		"user_name"  : "ban11111",
		"amount"     : "12.1245",
		"status"     : "under process",
	}).Expect()
	resp.Status(400)

	var respJson *service.BaseResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败,")
	s.Equal(false, respJson.Success, respJson.Info)
	s.Equal("文件不能为空，请先上传文件", respJson.Info)
}

//新建demo_order 测试(有文件)
func(s *HandlerTestSuite)Testcreatedemowithfile() {
	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithFile("file_url", "./testfile/123.png").WithForm(map[string]string{
		"order_id"   : "xx0098dd",
		"user_name"  : "ban22222",
		"amount"     : "99.1206",
		"status"     : "开始",
	}).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	if err != nil {
		fmt.Println("json转换成对象失败!!", err)
	}
	s.Equal(true, respJson.Success, respJson.Info)
	var DBdata []*model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
	s.Equal(true, len(DBdata)>0)
	s.NotEqual(0, DBdata[len(DBdata)-1].Id,"Id为0,不合法")
	s.Equal("xx0098dd", DBdata[len(DBdata)-1].Order_id,"Order id 不一致")
	s.Equal("ban22222", DBdata[len(DBdata)-1].User_name,"User name 不一致")
	s.Equal(99.1206, DBdata[len(DBdata)-1].Amount,"Amount 不一致")
	s.Equal("开始", DBdata[len(DBdata)-1].Status,"Status 不一致")
	s.Equal(true, len(DBdata[len(DBdata)-1].File_url)>0, "file url 为空！")
}

//更新demo_order 测试(更新第一条数据)
func(s *HandlerTestSuite)Testupdatedemo() {
	os.Chdir("F:/coding/go/src/QY_Homework/tests")
	resp := httpexpect.New(s.T(), TestServer).PUT(Updatedemo).WithMultipart().WithFile("file_url", "./testfile/update.jpg").WithForm(map[string]string{
		"id"		 : "1",
		"order_id"   : "123456789",
		"user_name"  : "ban123456",
		"amount"     : "0.123456",
		"status"     : "结束",
	}).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	//logger.Info("errr","er",err)
	s.Equal(nil, err, "json转换成对象失败!!")
	s.Equal(true, respJson.Success, respJson.Info)

	var DBdata []*model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
	s.Equal(true, len(DBdata)>0)

	if respJson.Success == true {
		s.Equal(respJson.Id, DBdata[0].Id, "Id为0,不合法")
		s.Equal(respJson.Order_id, DBdata[0].Order_id, "Order id 不一致")
		s.Equal(respJson.User_name, DBdata[0].User_name, "User name 不一致")
		s.Equal(respJson.Amount, DBdata[0].Amount, "Amount 不一致")
		s.Equal(respJson.Status, DBdata[0].Status, "Status 不一致")
		s.Equal(respJson.File_url, DBdata[0].File_url, "file url 不一致")
	}
}

//测试 查询单条数据
func (s *HandlerTestSuite)TestGetdemoinfo() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemoinfo).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败!!")
	var DBdata model.Demo_order
	openedDb := connection.ConnetDB(NewDbConfig())
	openedDb.Where("id = ?", 1).First(&DBdata)
	s.Equal(respJson.Demo_order, &DBdata, "返回的数据与数据库中的数据不匹配")
}
//测试 查询单条数据失败
func (s *HandlerTestSuite)TestGetdemoinfofail() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemoinfo+"99999").Expect()
	resp.Status(400)
	var respJson service.BaseResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	fmt.Println("******************************************",respJson)
	s.Equal(nil, err, "json转换成对象失败!!")
	s.Equal(false, respJson.Success, respJson.Info)
}

func (s *HandlerTestSuite)TestGetdemolist() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemolist).Expect()
	resp.Status(200)
}