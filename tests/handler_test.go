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
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HandlerTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	go router.Start_Server()
	time.Sleep(100 * time.Millisecond)
	suite.Run(t, new(HandlerTestSuite))
}

//新建demo_order 测试
func (s *HandlerTestSuite)Testcreatedemo() {
	resp := httpexpect.New(s.T(), TestServer).POST(Createdemo).WithMultipart().WithForm(map[string]string{
		"order_id"   : "xx1280a",
		"user_name"  : "ban11111",
		"amount"     : "12.1245",
		"status"     : "under process",
	}).Expect()
	resp.Status(200)
	var respJson *service.SuccessResp
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	if err != nil {
		fmt.Println("json转换成对象失败,", err)
	}
	s.Equal(true, respJson.Success, respJson.Info)
	var DBdata []*model.Demo_order
	connection.ConnetDB(NewDbConfig()).Model(&model.Demo_order{}).Find(&DBdata)
	s.Equal(true, len(DBdata)>0)
	s.Equal(respJson.Id, DBdata[len(DBdata)-1].Id,"Id是否一致")
	s.Equal(respJson.Order_id, DBdata[len(DBdata)-1].Order_id,"是否一致")
	s.Equal(respJson.User_name, DBdata[len(DBdata)-1].User_name,"是否一致")
	s.Equal(respJson.Amount, DBdata[len(DBdata)-1].Amount,"是否一致")
	s.Equal(respJson.Status, DBdata[len(DBdata)-1].Status,"是否一致")
}

//更新demo_order 测试
func(s *HandlerTestSuite)Testupdatedemo() {

}