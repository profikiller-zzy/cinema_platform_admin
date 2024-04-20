package movie_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
	"time"
)

type MovieResponse struct {
	model.MODEL
	MovieName       string        `json:"movie_name"`                      // 电影的名字
	CoverPictureUrl string        `json:"banner_url" structs:"banner_url"` // 电影封面的存储地址
	ReleaseDate     time.Time     `json:"release_date"`                    // 上映时间
	PlayTime        time.Duration `json:"play_time"`                       // 电影的时长
	Director        string        `json:"director"`                        // 电影的导演
	Actors          string        `json:"actors"`                          // 该电影的演员
}

type MovieListRequest struct {
	model.PageInfo
}

// MovieListView 为管理员返回全部用户信息
func (MovieApi) MovieListView(c *gin.Context) {
	var pageModel MovieListRequest
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}
	var movieRepList = make([]MovieResponse, 0)
	var movieList = make([]model.Movie, 0)
	var count int64
	// 对电影列表进行分页
	movieList, count, err = common_service.PagingList(model.Movie{}, common_service.Option{
		PageInfo: pageModel.PageInfo,
		Likes:    []string{"movie_name"},
	})

	for _, movie := range movieList {
		movieRepList = append(movieRepList, MovieResponse{
			MODEL:           movie.MODEL,
			MovieName:       movie.MovieName,
			CoverPictureUrl: movie.CoverPictureUrl,
			ReleaseDate:     movie.ReleaseDate,
			PlayTime:        movie.PlayTime,
			Director:        movie.Director,
			Actors:          movie.Actors,
		})
	}
	response.OKWithPagingData(movieRepList, count, c)
	return
}
