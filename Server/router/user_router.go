package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

// UserRouter 传入的是路由组
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/admin_login/", userApiApp.AdminLoginView)
	r.GET("/user_list/", middleware.JwtAuth(1), userApiApp.UserListView)
	r.GET("/individual_user/", middleware.JwtAuth(1), userApiApp.IndividualUserView)
	r.POST("/user_create/", middleware.JwtAuth(1), userApiApp.UserCreateView)
	r.PUT("/user_edit/", middleware.JwtAuth(1), userApiApp.UserEditView)
	r.DELETE("/user_remove/", middleware.JwtAuth(1), userApiApp.UserRemoveView)
}
