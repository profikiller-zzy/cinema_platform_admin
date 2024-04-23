package api

import (
	"AfterEnd/api/cinema_api"
	"AfterEnd/api/data_statistics_api"
	"AfterEnd/api/image_api"
	"AfterEnd/api/movie_api"
	"AfterEnd/api/user_api"
)

// ApiGroup 是对整个Api定义的结构体的统合，方便链式调用
type ApiGroup struct {
	UserApi   user_api.UserApi
	CinemaApi cinema_api.CinemaApi
	ImageApi  image_api.ImageApi
	MovieApi  movie_api.MovieApi
	DataApi   data_statistics_api.DataApi
}

// ApiGroupApp 实例化ApiGroup对象
var ApiGroupApp = new(ApiGroup)
