package movie_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type TimePeriodRequest struct {
	MovieID   string    `json:"movie_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type BoxOfficeResponse struct {
	TotalBoxOffice  float64   `json:"total_box_office"`
	BoxOfficePerDay []float64 `json:"box_office_per_day"`
	Date            []string  `json:"date"`
	MovieName       string    `json:"movie_name"`
}

// BoxOfficeOfMovieInSpecificTimePeriod 查询指定电影指定时间内的票房
func (MovieApi) BoxOfficeOfMovieInSpecificTimePeriod(c *gin.Context) {
	var tpReq TimePeriodRequest
	err := c.ShouldBindJSON(&tpReq)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	// 将字符串类型的 movie_id 转换为整数类型
	movieID, err := strconv.Atoi(tpReq.MovieID)
	if err != nil {
		// 处理转换错误
		return
	}

	DB := global.Db.Session(&gorm.Session{Logger: global.MysqlLog})

	// 计算指定时间段内的每天日期
	dateRange := make([]string, 0)
	for d := tpReq.StartTime; d.Before(tpReq.EndTime); d = d.AddDate(0, 0, 1) {
		dateRange = append(dateRange, d.Format("2006-01-02"))
	}

	// 查询每天的票房并累加得到总票房
	boxOfficePerDay := make([]float64, len(dateRange))
	totalBoxOffice := 0.0
	for i, date := range dateRange {
		boxOffice, err := model.BoxOfficeOfMovieInSpecificDay(uint(movieID), DB, date)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("服务器出错！, err : %s", err.Error()), c)
			return
		}
		boxOfficePerDay[i] = boxOffice
		totalBoxOffice += boxOffice
	}

	movieName, err := model.GetMovieNameByID(uint(movieID), DB)
	if err != nil {
		response.FailWithMessage("获取电影名称失败！", c)
		return
	}

	// 返回总票房和每天的票房数据
	resp := BoxOfficeResponse{
		TotalBoxOffice:  totalBoxOffice,
		BoxOfficePerDay: boxOfficePerDay,
		Date:            dateRange,
		MovieName:       movieName,
	}
	response.OKWithData(resp, c)
}
