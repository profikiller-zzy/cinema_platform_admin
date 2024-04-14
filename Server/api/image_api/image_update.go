package image_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" msg:"请选择图片ID"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

func (ImageApi) ImageUpdateView(c *gin.Context) {
	var iuReq ImageUpdateRequest
	err := c.ShouldBindJSON(&iuReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &iuReq, c)
		return
	}
	var image model.ProfilePhoto
	err = global.Db.Take(&image, iuReq.ID).Error
	if err != nil { // 没有找到
		response.FailWithMessage("文件不存在", c)
		return
	}
	// 找到了就将传入的name替换掉旧的name
	err = global.Db.Model(&image).Update("name", iuReq.Name).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("修改广告成功", c)
		return
	}
	response.OKWithMessage("图片名称修改成功", c)
}
