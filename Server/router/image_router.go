package router

import (
	"AfterEnd/api"
)

func (r RGroup) ImageRouter() {
	imageApiApp := api.ApiGroupApp.ImageApi
	r.POST("/image_upload/", imageApiApp.ImageUploadingView)
}
