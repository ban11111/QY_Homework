package service

import (
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

//上传文件，并返回file—url
func Uploadfiles(file []*multipart.FileHeader,id uint64, c *gin.Context) (file_url string){
	if file == nil {
		return
	}
	os.Chdir(UploadPath)
	fpath := "f" + strconv.FormatUint(id,10)
	file_path := UploadPath + fpath + "/"
	os.Mkdir(fpath, 0777)
	for _, f := range file{
		c.SaveUploadedFile(f, file_path + string([]rune(f.Filename)[10:]))

	}
	return "localhost:8080" + PublicURL + fpath + "/"
}


func Updatefiles(){

}