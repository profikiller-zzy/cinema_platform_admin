package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

func (r RGroup) CinemaRouter() {
	cinemaApiApp := api.ApiGroupApp.CinemaApi
	r.GET("/unreviewed_approval_list/", middleware.JwtAuth(1), cinemaApiApp.UnreviewedApprovalListView)
}
