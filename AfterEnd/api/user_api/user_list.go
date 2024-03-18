package user_api

import (
	"AfterEnd/model"
	"AfterEnd/model/ctype"
	"AfterEnd/model/response"
	"AfterEnd/service/common_service"
	"AfterEnd/utils/desen"
	"AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	model.MODEL
	UserName string     `json:"user_name"` // 用户名，唯一
	Age      string     `json:"age"`       // 年龄
	Tel      string     `json:"tel"`       // 电话号码
	Email    string     `json:"email"`     // 邮箱，用户可以通过邮箱登录
	UserType ctype.Role `json:"user_type"` // 用户类别，用于区分普通用户、影院用户、平台管理员(属于自定义类型，后期需要添加)，3为普通用户，2为电影院用户，1为平台管理员
}

type UserListRequest struct {
	model.PageInfo
	Role int `json:"role" form:"role"`
}

// UserListView 为管理员返回全部用户信息
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}
	var userRepList = make([]UserResponse, 0)
	var userList = make([]model.User, 0)
	var count int64
	// 对用户列表进行分页
	userList, count, err = common_service.PagingList(model.User{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})

	for _, user := range userList {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员，不显示用户名
			user.UserName = ""
		}
		// 数据脱敏
		user.Tel = desen.DesensitizationTel(user.Tel)
		user.Email = desen.DesensitizationEmail(user.Email)
		userRepList = append(userRepList, UserResponse{
			MODEL:    user.MODEL,
			UserName: user.UserName,
			Age:      user.Age,
			Tel:      user.Tel,
			Email:    user.Email,
			UserType: user.UserType,
		})
	}
	response.OKWithPagingData(userRepList, count, c)
	return
}
