package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

func (r RGroup) CinemaRouter() {
	cinemaApiApp := api.ApiGroupApp.CinemaApi
	r.GET("/cinemas/", middleware.JwtAuth(1), cinemaApiApp.CinemaListView)
}
