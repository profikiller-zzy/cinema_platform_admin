package user_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/utils/jwts"
	"AfterEnd/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string `json:"password" binding:"required" msg:"请输入密码"`   // 密码
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var ELReq EmailLoginRequest
	err := c.ShouldBindJSON(&ELReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	// 验证用户是否存在
	var userModel model.User
	err = global.Db.Take(&userModel, "user_name = ? or email = ?", ELReq.UserName, ELReq.UserName).Error
	if err != nil { // 该用户不存在
		global.Log.Warnln("用户名不存在")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码是否正确
	pwdIsCorrect := pwd.VerifyPwd(ELReq.Password, userModel.Password)
	if !pwdIsCorrect {
		global.Log.Warnln("用户名密码错误")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 验证成功，生成Token
	tokenString, err := jwts.GenerateToken(jwts.JwtPayLoad{
		UserID:   userModel.ID,
		UserName: userModel.UserName,
		Role:     int(userModel.UserType),
	})
	if err != nil {
		global.Log.Warnln(err.Error())
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	response.OKWithData(tokenString, c)
}
