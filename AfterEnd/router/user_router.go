package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

// UserRouter 传入的是路由组
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/email_login/", userApiApp.EmailLoginView)
	r.GET("/user_list/", middleware.JwtAuth(1), userApiApp.UserListView)
}
