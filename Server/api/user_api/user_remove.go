package user_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserRemoveView 这个函数目前还不完善，删除用户之前需要先将指定用户的全部依赖删除
func (UserApi) UserRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，请确认请求类型为JSON，error：%s", err.Error()), c)
		return
	}

	var umRes model.UserRemoveResponse
	var umResList []model.UserRemoveResponse = make([]model.UserRemoveResponse, len(rmReq.IDList))
	for index, userID := range rmReq.IDList {
		if err, umRes = model.SoftDeleteUser(userID); err == nil {
			count++
			umResList[index] = umRes
		} else {
			umResList[index] = umRes
		}
	}

	response.OKWithData(umResList, c)
}
