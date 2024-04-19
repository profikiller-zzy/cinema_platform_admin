package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

// UserRouter 传入的是路由组
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/admin_login/", userApiApp.AdminLoginView)
	r.GET("/users/", middleware.JwtAuth(1), userApiApp.UserListView)
	//r.GET("/individual_user/", middleware.JwtAuth(1), userApiApp.IndividualUserView)
	r.POST("/users/", middleware.JwtAuth(1), userApiApp.UserCreateView)
	r.PUT("/users/", middleware.JwtAuth(1), userApiApp.UserEditView)
	r.DELETE("/users/", middleware.JwtAuth(1), userApiApp.UserRemoveView)
}
