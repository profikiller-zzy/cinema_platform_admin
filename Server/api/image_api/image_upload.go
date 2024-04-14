package image_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/service/image_service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// ImageUploadingView 上传图片至七牛云存储空间
func (ImageApi) ImageUploadingView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.LogFail(err, c)
		return
	}

	var FileHeaderList []*multipart.FileHeader = form.File["image"]
	if len(FileHeaderList) == 0 {
		response.FailWithMessage("没有指定任何文件或者文件不存在", c)
		return
	}

	var upResList []model.FileUploadResponse = make([]model.FileUploadResponse, len(FileHeaderList))
	for index, FileHeader := range FileHeaderList {
		upResList[index] = image_service.ImageService{}.ImageUploadService(FileHeader, c)
	}
	response.OKWithData(upResList, c)
}
