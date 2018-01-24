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
var CurrentPath string

type HandlerTestSuite struct {
	suite.Suite
}

func TestStart(t *testing.T) {
	resp := httpexpect.New(t, TestServer).POST(Createdemo).WithMultipart().WithFile("file_url", "./testfile/123.png").WithForm(map[string]string{
		"order_id":  "test",
		"user_name": "ban22222",
		"amount":    "12.12",
		"status":    "testnew",
	}).Expect()
	resp.Status(200)
}

func (suite *HandlerTestSuite)TearDownTest() {
}

func TestSuite(t *testing.T) {
	//测试前获取当前路径
	CurrentPath, _ = os.Getwd()
	//logger.InitLogger(logger.LvlDebug,nil)
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	os.Chdir(CurrentPath)
	TestStart(t)
	os.Chdir(CurrentPath)
	TestStart(t)
	suite.Run(t, new(HandlerTestSuite))
}

//新建demo_order 测试(无文件)
func (s *HandlerTestSuite) Testcreatedemowithoutfile() {
	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithForm(map[string]string{
		"order_id":  "xx1280a",
		"user_name": "ban11111",
		"amount":    "12.1245",
		"status":    "this one will fail",
	}).Expect()
	resp.Status(400)
	var respJson service.BaseResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败,")
	s.Equal(false, respJson.Success, respJson.Info)
	s.Equal("文件不能为空，请先上传文件", respJson.Info)
}

//新建demo_order 测试(有文件)
func (s *HandlerTestSuite) Testcreatedemowithfile() {
	os.Chdir(CurrentPath)
	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithFile("file_url", "./testfile/123.png").WithForm(map[string]string{
		"order_id":  "xx0098dd",
		"user_name": "ban22222",
		"amount":    "99.1206",
		"status":    "new",
	}).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	if err != nil {
		fmt.Println("json转换成对象失败!!", err)
	}
	s.Equal(true, respJson.Success, respJson.Info)
	var DBdata []model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
	s.Equal(true, len(DBdata) > 0)
	s.NotEqual(0, DBdata[len(DBdata)-1].Id, "Id为0,不合法")
	s.Equal("xx0098dd", DBdata[len(DBdata)-1].Order_id, "Order id 不一致")
	s.Equal("ban22222", DBdata[len(DBdata)-1].User_name, "User name 不一致")
	s.Equal(99.1206, DBdata[len(DBdata)-1].Amount, "Amount 不一致")
	s.Equal("new", DBdata[len(DBdata)-1].Status, "Status 不一致")
	s.Equal(true, len(DBdata[len(DBdata)-1].File_url) > 0, "file url 为空！")
	s.Equal(respJson.Demo_order, DBdata[len(DBdata)-1])
}

//更新demo_order 测试(更新第一条数据)
func (s *HandlerTestSuite) Testupdatedemo() {
	os.Chdir(CurrentPath)
	resp := httpexpect.New(s.T(), TestServer).PUT(Updatedemo).WithMultipart().WithFile("file_url", "./testfile/update.jpg").WithForm(map[string]string{
		//"id":        "1",
		"order_id":  "123456789",
		"user_name": "ban123456",
		"amount":    "0.123456",
		"status":    "updated",
	}).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	//logger.Info("errr","er",err)
	s.Equal(nil, err, "json转换成对象失败!!")
	s.Equal(true, respJson.Success, respJson.Info)

	if respJson.Success == true {
		s.Equal(uint64(1), respJson.Id, "Id不对")
		s.Equal("123456789", respJson.Order_id, "Order id 不一致")
		s.Equal("ban123456", respJson.User_name, "User name 不一致")
		s.Equal(0.123456, respJson.Amount, "Amount 不一致")
		s.Equal("updated", respJson.Status, "Status 不一致")
		s.Equal("localhost:8080/public/f1/", respJson.File_url, "file url 不一致")
	}
}

//测试 事务SQL
func (s *HandlerTestSuite) TestTransaction() {
	os.Chdir(CurrentPath)
	connection.ConnetDB(NewDbConfig()).Where("id = ?", "2").Delete(model.Files{})
	resp := httpexpect.New(s.T(), TestServer).PUT(Transction).WithMultipart().WithFile("file_url", "./testfile/update.jpg").WithForm(map[string]string{
		"order_id":  "should fail",
		"user_name": "should fail",
		"amount":    "123.456789",
		"status":    "should fail",
	}).Expect()
	resp.Status(400)
	var DBdata model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).First(&DBdata)
	s.NotEqual("should fail", DBdata.Order_id)
}

//测试 查询单条数据
func (s *HandlerTestSuite) TestGetdemoinfo() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemoinfo).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败!!")
	var DBdata model.Demo_order
	openedDb := connection.ConnetDB(NewDbConfig())
	openedDb.Where("id = ?", 1).First(&DBdata)
	s.Equal(respJson.Demo_order, DBdata, "返回的数据与数据库中的数据不匹配")
}

//测试 查询单条数据失败
func (s *HandlerTestSuite) TestGetdemoinfofail() {
	resp := httpexpect.New(s.T(), TestServer).GET(Getdemoinfo + "99999").Expect()
	resp.Status(400)
	var respJson service.BaseResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败!!")
	s.Equal(false, respJson.Success, respJson.Info)
}

//测试 查询默认列表数据
func (s *HandlerTestSuite) TestGetdemolist() {
	resp := httpexpect.New(s.T(), TestServer).GET(Postdemolist).Expect()
	resp.Status(200)
	var respJson service.ListResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err, "json转换成对象失败!!")
	var DBdata []model.Demo_order
	openedDb := connection.ConnetDB(NewDbConfig())
	openedDb.Find(&DBdata)
	for i := range respJson.List {
		s.Equal(respJson.List[i].Order_id, DBdata[i].Order_id)
	}
}

//测试4种排序查询列表数据
func (s *HandlerTestSuite) TestPostdemolist() {
	sort := []string{"time", "timedesc", "amount", "amountdesc"}
	for _, sortby := range sort {
		resp := httpexpect.New(s.T(), TestServer).POST(Postdemolist).WithJSON(map[string]string{
			"search": "ban",
			"sortby": sortby,
		}).Expect()
		resp.Status(200)
		var respJson service.ListResp
		err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
		s.Equal(nil, err, "json转换成对象失败!!")
		switch sortby {
		case "time":
			//s.Equal(true, respJson.List[0].CreatedAt.Before(respJson.List[1].CreatedAt))
		case "timedesc":
			//s.Equal(true, respJson.List[0].CreatedAt.After(respJson.List[1].CreatedAt))
		case "amount":
			s.Equal(true, respJson.List[0].Amount <= respJson.List[1].Amount)
		case "amountdesc":
			s.Equal(true, respJson.List[0].Amount >= respJson.List[1].Amount)
		default:
			fmt.Println("应该不会运行到这里吧")
		}
	}
}
