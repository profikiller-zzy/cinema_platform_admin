package movie_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
	"time"
)

// MovieRank 返回的电影排行榜
type MovieRank struct {
	Rank        uint    `json:"rank"`
	MovieID     uint    `json:"movie_id"`
	MovieName   string  `json:"movie_name"`
	ReleaseDate string  `json:"release_date"`
	BoxOffice   float64 `json:"box_office"`
}

// MovieBoxOfficeRanking 需要将电影以最近一周票房总数进行排行
func (MovieApi) MovieBoxOfficeRanking(c *gin.Context) {
	// TODO 查询出给定时间区间中的电影票房排行前100名
	DB := global.Db.Session(&gorm.Session{Logger: global.MysqlLog})
	moviesInfo, err := model.GetAllMovies(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询出错，err：%s", err.Error()), c)
	}

	rankings := make([]MovieRank, 0, 100) // 开辟长度为 100 的切片，但初始长度为 0
	endTime := time.Now()
	startTime := endTime.AddDate(0, 0, -7)
	for _, value := range moviesInfo {
		releaseDate := value.ReleaseDate.Format("2006-01-02")
		boxOffice, err := model.BoxOfficeOfMovieInRecentWeek(value.ID, DB, startTime, endTime)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("在计算电影 '%s' 的票房时出错！, err : %s", value.MovieName, err.Error()), c)
			return
		}
		// 如果票房为 0，则跳过该电影
		if boxOffice == 0 {
			continue
		}
		rank := MovieRank{
			MovieID:     value.ID,
			MovieName:   value.MovieName,
			BoxOffice:   boxOffice,
			ReleaseDate: releaseDate,
		}
		rankings = append(rankings, rank)
	}

	// 对排行榜按票房总额降序排序
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].BoxOffice > rankings[j].BoxOffice
	})

	// 如果排行榜的长度大于 100，则截断切片
	if len(rankings) > 100 {
		rankings = rankings[:100]
	}

	// 填入排名
	for i := 0; i < len(rankings); i++ {
		rankings[i].Rank = uint(i + 1)
	}

	// 返回结果
	response.OKWithData(rankings, c)
}
