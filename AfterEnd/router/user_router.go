package router

import "AfterEnd/api"

// UserRouter 传入的是路由组
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/email_login/", userApiApp.EmailLoginView)
}
