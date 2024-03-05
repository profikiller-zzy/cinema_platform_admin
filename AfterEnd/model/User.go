package model

import "AfterEnd/model/ctype"

// User 用户
type User struct {
	MODEL
	UserName string     `gorm:"size:64" json:"user_name"`              // 用户名，唯一
	Password string     `gorm:"size:64" json:"password"`               // 密码
	Age      string     `gorm:"size 4" json:"age"`                     // 年龄
	Email    string     `gorm:"size:128" json:"email"`                 // 邮箱
	UserType ctype.Role `gorm:"size:16" json:"user_type"`              // 用户类别，用于区分普通用户、影院用户、平台管理员(属于自定义类型，后期需要添加)，3为普通用户，2为电影院用户，1为平台管理员
	Orders   []Order    `gorm:"foreignKey:UserID" json:"orders"`       // 该用户的订单记录
	Reviews  []Review   `gorm:"foreignKey:UserID" json:"reviews"`      // 该用户发表的所有评论
	Cinema   []Cinema   `gorm:"foreignKey:CinemaUserID" json:"cinema"` // 如果该用户是影院用户(经营着一家影院)，则这是该用户所拥有的影院
}
