package model

// User 用户
type User struct {
	MODEL
	UserName string `gorm:"size:64" json:"user_name"` // 用户名，唯一
	Password string `gorm:"size:64" json:"password"`  // 密码
	Age      string `gorm:"size 4" json:"age"`        // 年龄
	Email    string `json:"email"`                    // 邮箱
	UserType string `json:"user_type"`                // 用户类别，用于区分普通用户、影院用户、平台管理员(属于自定义类型，后期需要添加)，3为普通用户，2为电影院用户，1为平台管理员
}
