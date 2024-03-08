package api

import "AfterEnd/api/user_api"

// ApiGroup 是对整个Api定义的结构体的统合，方便链式调用
type ApiGroup struct {
	UserApi user_api.UserApi
}

// ApiGroupApp 实例化ApiGroup对象
var ApiGroupApp = new(ApiGroup)
