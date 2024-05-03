package model

import (
	"AfterEnd/global"
	"AfterEnd/model/ctype"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	MODEL
	UserName  string         `gorm:"size:64" json:"user_name"`              // 用户名，唯一
	Password  string         `gorm:"size:64" json:"password"`               // 密码
	AvatarUrl string         `gorm:"size:256" json:"avatar"`                // 头像地址
	Age       string         `gorm:"size 4" json:"age"`                     // 年龄
	Tel       string         `gorm:"size:18" json:"tel"`                    // 电话号码
	Email     string         `gorm:"size:128" json:"email"`                 // 邮箱，用户可以通过邮箱登录
	UserType  ctype.Role     `gorm:"size:16" json:"user_type"`              // 用户类别，用于区分普通用户、影院用户、平台管理员(属于自定义类型，后期需要添加)，3为普通用户，2为电影院用户，1为平台管理员
	Orders    []Order        `gorm:"foreignKey:UserID" json:"orders"`       // 该用户的订单记录
	Reviews   []Review       `gorm:"foreignKey:UserID" json:"reviews"`      // 该用户发表的所有评论
	Cinema    []Cinema       `gorm:"foreignKey:CinemaUserID" json:"cinema"` // 如果该用户是影院用户(经营着一家影院)，则这是该用户所拥有的影院
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`               // 软删除字段
}

func SoftDeleteUser(userID uint) (error, ListRemoveResponse) {
	// 根据用户ID查询用户
	var user User
	var umRes ListRemoveResponse
	result := global.Db.Where("id = ?", userID).First(&user)
	err := result.Error
	// 删除失败
	if err != nil {
		umRes = ListRemoveResponse{
			UserID:    user.ID,
			IsSuccess: false,
			Msg:       fmt.Sprintf("%s", err),
		}
		return result.Error, umRes
	}

	// 检查是否是管理员
	if user.UserType == 1 {
		umRes = ListRemoveResponse{
			UserID:    user.ID,
			IsSuccess: false,
			Msg:       "管理员不允许删除",
		}
		return errors.New("管理员不允许删除"), umRes
	}

	// 软删除用户
	result = global.Db.Delete(&user)
	err = result.Error
	if err != nil {
		umRes = ListRemoveResponse{
			UserID:    user.ID,
			IsSuccess: false,
			Msg:       fmt.Sprintf("%s", err),
		}
		return result.Error, umRes
	}

	umRes = ListRemoveResponse{
		UserID:    user.ID,
		IsSuccess: true,
		Msg:       fmt.Sprintf("删除用户 %d 成功!", user.ID),
	}
	return nil, umRes
}

func CountUsers(db *gorm.DB) (int64, error) {
	var count int64
	result := db.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
