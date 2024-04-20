package movie_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MovieRemoveView 下架电影
func (MovieApi) MovieRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，请确认请求类型为JSON，error：%s", err.Error()), c)
		return
	}

	var mrRes model.ListRemoveResponse
	var umResList []model.ListRemoveResponse = make([]model.ListRemoveResponse, len(rmReq.IDList))
	for index, movieID := range rmReq.IDList {
		if err, mrRes = model.SoftDeleteMovie(movieID); err == nil {
			count++
			umResList[index] = mrRes
		} else {
			umResList[index] = mrRes
		}
	}

	response.OKWithData(umResList, c)
}
