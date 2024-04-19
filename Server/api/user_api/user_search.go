package user_api

import (
	"AfterEnd/model"
	"AfterEnd/model/ctype"
	"AfterEnd/model/response"
	"AfterEnd/service/common_service"
	"AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

// UserSearchView 管理员搜索用户名，返回与用户名有关的用户
func (UserApi) UserSearchView(c *gin.Context) {
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
		//// 数据脱敏
		//user.Tel = desen.DesensitizationTel(user.Tel)
		//user.Email = desen.DesensitizationEmail(user.Email)
		userRepList = append(userRepList, UserResponse{
			MODEL:     user.MODEL,
			UserName:  user.UserName,
			AvatarUrl: user.AvatarUrl,
			Age:       user.Age,
			Tel:       user.Tel,
			Email:     user.Email,
			UserType:  user.UserType,
			RoleID:    int(user.UserType),
		})
	}
	response.OKWithPagingData(userRepList, count, c)
	return
}
