package router

import (
	"AfterEnd/api"
	"AfterEnd/middleware"
)

func (r RGroup) MoiveRouter() {
	movieApiApp := api.ApiGroupApp.MovieApi
	r.POST("/movies/", middleware.JwtAuth(1), movieApiApp.MovieCreateView)
	r.GET("/movies/", middleware.JwtAuth(1), movieApiApp.MovieListView)
	r.GET("/movies_rank/", middleware.JwtAuth(1), movieApiApp.MovieBoxOfficeRanking)
	r.DELETE("/movies/", middleware.JwtAuth(1), movieApiApp.MovieRemoveView)
	r.PUT("/movies/", middleware.JwtAuth(1), movieApiApp.MovieEditView)
	r.POST("/movie_cover/", middleware.JwtAuth(1), movieApiApp.MovieCoverUpload)
}
