package movie_api

import (
	"AfterEnd/model/response"
	"AfterEnd/service/image_service"
	"github.com/gin-gonic/gin"
)

func (MovieApi) MovieCoverUpload(c *gin.Context) {
	// 获取上传的封面图片
	fileHeader, err := c.FormFile("coverPicture")
	if err != nil {
		response.FailWithMessage("文件不存在", c)
		return
	}

	// 调用图片上传服务处理上传的文件
	upRes := image_service.ImageService{}.ImageUploadService(fileHeader, c)
	// 假如上传图片未成功且失败原因不是图片已存在
	if !upRes.IsSuccess && upRes.Msg != "图片已存在" {
		response.FailWithMessage(upRes.Msg, c)
		return
	}
	response.OKWithData(upRes, c)
}
