package tests

import (
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

	"fmt"
	"gcoresys/common/logger"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HandlerTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	logger.InitLogger(logger.LvlDebug,nil)
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	suite.Run(t, new(HandlerTestSuite))
}

//新建demo_order 测试(无文件)
//func (s *HandlerTestSuite)Testcreatedemowithoutfile() {
//	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithForm(map[string]string{
//		"order_id"   : "xx1280a",
//		"user_name"  : "ban11111",
//		"amount"     : "12.1245",
//		"status"     : "under process",
//	}).Expect()
//	resp.Status(200)
//	var respJson *service.SuccessResp
//	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
//	if err != nil {
//		fmt.Println("json转换成对象失败,", err)
//	}
//	s.Equal(true, respJson.Success, respJson.Info)
//	var DBdata []*model.Demo_order
//	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
//	s.Equal(true, len(DBdata)>0)
//	s.Equal(respJson.Id, DBdata[len(DBdata)-1].Id,"Id是否一致")
//	s.Equal(respJson.Order_id, DBdata[len(DBdata)-1].Order_id,"是否一致")
//	s.Equal(respJson.User_name, DBdata[len(DBdata)-1].User_name,"是否一致")
//	s.Equal(respJson.Amount, DBdata[len(DBdata)-1].Amount,"是否一致")
//	s.Equal(respJson.Status, DBdata[len(DBdata)-1].Status,"是否一致")
//	s.Equal("", respJson.File_url)
//}

//func(s *HandlerTestSuite)Testcreatedemowithfile() {
//	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithFile("file_url", "./testfile/123.png").WithForm(map[string]string{
//		"order_id"   : "xx0098dd",
//		"user_name"  : "ban22222",
//		"amount"     : "99.1206",
//		"status"     : "Finished",
//	}).Expect()
//	resp.Status(200)
//	var respJson service.SuccessResp
//	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
//	if err != nil {
//		fmt.Println("json转换成对象失败,", err)
//	}
//	s.Equal(true, respJson.Success, respJson.Info)
//	var DBdata []*model.Demo_order
//	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
//	s.Equal(true, len(DBdata)>0)
//	s.NotEqual(0, DBdata[len(DBdata)-1].Id,"Id是否不为0")
//	s.Equal("xx0098dd", DBdata[len(DBdata)-1].Order_id,"是否一致")
//	s.Equal("ban22222", DBdata[len(DBdata)-1].User_name,"是否一致")
//	s.Equal(99.1206, DBdata[len(DBdata)-1].Amount,"是否一致")
//	s.Equal("Finished", DBdata[len(DBdata)-1].Status,"是否一致")
//	s.Equal(true, len(DBdata[len(DBdata)-1].File_url)>0, "是否有值")
//}

//更新demo_order 测试
func(s *HandlerTestSuite)Testupdatedemo() {
	resp := httpexpect.New(s.T(), TestServer).PUT(Updatedemo).WithMultipart().WithFile("file_url", "./testfile/update.jpg").WithForm(map[string]string{
		"id"		 : "1",
		"order_id"   : "123456789",
		"user_name"  : "ban123456",
		"amount"     : "0.123456",
		"status"     : "Started",
	}).Expect()
	resp.Status(200)
	var respJson service.SuccessResp
	fmt.Println("----------",resp.Body().Raw())
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	fmt.Println("what the fuck !")
	//logger.Info("errr","er",err)
	s.Equal(nil, err)
	s.Equal(true, respJson.Success, respJson.Info)

	var DBdata []*model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
	s.Equal(true, len(DBdata)>0)

	s.Equal(respJson.Id, DBdata[0].Id,"Id是否一致")
	s.Equal(respJson.Order_id, DBdata[0].Order_id,"是否一致")
	s.Equal(respJson.User_name, DBdata[0].User_name,"是否一致")
	s.Equal(respJson.Amount, DBdata[0].Amount,"是否一致")
	s.Equal(respJson.Status, DBdata[0].Status,"是否一致")
	s.Equal(respJson.File_url, DBdata[0].File_url, "是否一致")
}