package user_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/utils/desen"
	"github.com/gin-gonic/gin"
)

type ID struct {
	ID uint `form:"id"` // 单个用户的id
}

// IndividualUserView 为管理员返回单个用户的信息
func (UserApi) IndividualUserView(c *gin.Context) {
	var id ID
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	var user model.User
	err = global.Db.Take(&user, "id = ?", id.ID).Error
	if err != nil {
		response.FailWithMessage("没有与该id匹配的用户", c)
		return
	}
	var userRep UserResponse
	// 数据脱敏
	user.Tel = desen.DesensitizationTel(user.Tel)
	user.Email = desen.DesensitizationEmail(user.Email)
	userRep = UserResponse{
		MODEL:    user.MODEL,
		UserName: user.UserName,
		Age:      user.Age,
		Tel:      user.Tel,
		Email:    user.Email,
		UserType: user.UserType,
	}
	response.OKWithData(userRep, c)
	return
}
