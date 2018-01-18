package model

type Demo_order struct {
	Id         uint 			`gorm:"primary_key"`
	Order_id   string
	User_name  string
	Amount     float64
	Status     string
	File_url   string
}