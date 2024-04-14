package user_service

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/ctype"
	"AfterEnd/utils/pwd"
	"errors"
)

const Avatar = "/uploads/avatar/default.png" // 默认的头像地址

func (UserService) CreateUser(userName, password string, role ctype.Role, age string, email, tel string) error {
	// 判断用户名是否存在
	var userModel model.User
	err := global.Db.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashPwd := pwd.BcryptPw(password)

	// 入库
	err = global.Db.Create(&model.User{
		UserName: userName,
		Password: hashPwd,
		UserType: role,
		Email:    email,
		Age:      age,
		Tel:      tel,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
