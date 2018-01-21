package service

import (
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"errors"
)

//上传文件，并返回file—url
func Uploadfiles(file []*multipart.FileHeader,id uint64, c *gin.Context) (file_url string, err error){
	if file == nil {
		return "", errors.New("文件不能为空，请先上传文件")
	}
	//cwd, _ := os.Getwd()
	//fmt.Println("Work dir:", cwd)
	if err = os.Chdir(UploadPath); err != nil {
		return "", err
	}
	fpath := "f" + strconv.FormatUint(id,10)
	file_path := UploadPath + fpath + "/"
	if err = os.Mkdir(fpath, 0777); err != nil {
		return "", err
	}
	for _, f := range file{
		err = c.SaveUploadedFile(f, file_path + string([]rune(f.Filename)[10:]))
		if err != nil {
			return "", err
		}
	}
	return "localhost:8080" + PublicURL + fpath + "/", err
}

//更新文件
func Updatefiles(file []*multipart.FileHeader,id uint64, c *gin.Context) (file_url string, err error){
	if file == nil {
		return "", errors.New("文件不能为空，请先上传文件")
	}
	fpath := "f" + strconv.FormatUint(id,10)
	file_path := UploadPath + fpath + "/"
	for _, f := range file{
		err = c.SaveUploadedFile(f, file_path + string([]rune(f.Filename)[10:]))
		if err != nil {
			return "", err
		}
	}
	return "localhost:8080" + PublicURL + fpath + "/", err
}