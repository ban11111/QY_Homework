package model

import "errors"

type Demo_order struct {
	Id         uint 			`gorm:"primary_key" json:"id"`
	Order_id   string			`json:"order_id" binding:"required"`
	User_name  string			`json:"user_name" binding:"required"`
	Amount     float64			`json:"amount" binding:"required"`
	Status     string			`json:"status"`
	File_url   string			`json:"file_url"`
}


func (demo *Demo_order)IsValid() (err error) {
	switch {
	//case demo.Id == 0:
	//	err = errors.New("id错误")
	case demo.Order_id == "":
		err = errors.New("Order id 不能为空")
	case demo.User_name == "":
		err = errors.New("User name不能为空")
	case demo.Amount == 0:
		err = errors.New("amount 不能为零")
	case demo.Status == "":
		err = errors.New("Status 不能为空")
	default:
		err = nil
	}
	return
}