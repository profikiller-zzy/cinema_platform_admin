package user_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/ctype"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserRoleRequest struct {
	UserType ctype.Role `json:"user_type" binding:"required,oneof=1 2 3" msg:"输入正确的权限等级"`
	Age      string     `gorm:"size 4" json:"age"`     // 年龄
	Tel      string     `gorm:"size:18" json:"tel"`    // 电话号码
	Email    string     `gorm:"size:128" json:"email"` // 邮箱
	UserID   uint       `json:"user_id" binding:"required" msg:"请确认需要修改的用户ID"`
}

// UserEditView 编辑用户信息
func (UserApi) UserEditView(c *gin.Context) {
	var urReq UserRoleRequest
	err := c.ShouldBindJSON(&urReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &urReq, c)
		return
	}

	var userModel model.User
	err = global.Db.First(&userModel, urReq.UserID).Error
	if err != nil {
		response.FailWithMessage("输入的用户ID错误", c)
		return
	}
	err = global.Db.Model(&userModel).Updates(map[string]interface{}{
		"user_type": urReq.UserType,
		"age":       urReq.Age,
		"tel":       urReq.Tel,
		"email":     urReq.Email,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("用户 %d 权限成功变更为%s,用户年龄、电话号码和邮箱分别变更为'%s','%s','%s'", urReq.UserID, urReq.UserType.String(), urReq.Age, urReq.Tel, urReq.Email), c)
}
