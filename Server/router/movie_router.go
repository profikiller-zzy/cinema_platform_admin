package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

func (r RGroup) MoiveRouter() {
	movieApiApp := api.ApiGroupApp.MovieApi
	r.POST("/movies/", middleware.JwtAuth(1), movieApiApp.MovieCreateView)
	r.GET("/movies/", middleware.JwtAuth(1), movieApiApp.MovieListView)
}
