package model

import "errors"

type Files struct{
	Id uint64				`gorm:"primary_key"`
	File_path string
}

func (files *Files)FilesIsValid() (err error) {
	switch {
	//case files.Id == 0:
	//	err = errors.New("id错误")
	case files.File_path == "":
		err = errors.New("Order id 不能为空")
	default:
		err = nil
	}
	return
}