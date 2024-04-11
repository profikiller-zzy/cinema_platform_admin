package user_api

import (
	"AfterEnd/global"
	"AfterEnd/model/ctype"
	"AfterEnd/model/response"
	"AfterEnd/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 影院用户  3 普通用户
}

func (UserApi) UserCreateView(c *gin.Context) {
	var ucReq UserCreateRequest
	if err := c.ShouldBindJSON(&ucReq); err != nil {
		response.FailBecauseOfParamError(err, &ucReq, c)
		return
	}
	err := user_service.UserService{}.CreateUser(ucReq.UserName, ucReq.Password, ucReq.Role, "")
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("用户%s创建成功!", ucReq.UserName), c)
	return
}
