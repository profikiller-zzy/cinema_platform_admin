package data_statistics_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dataStatisticsResponse struct {
	BoxOfficeDay   float64 `json:"box_office_day"`
	BoxOfficeWeek  float64 `json:"box_office_week"`
	BoxOfficeMonth float64 `json:"box_office_month"`
	TotalUsers     int     `json:"total_users"`
	TotalOrders    int     `json:"total_orders"`
}

func (DataApi) DataStatistics(c *gin.Context) {
	var dsRes dataStatisticsResponse
	DB := global.Db.Session(&gorm.Session{Logger: global.MysqlLog})

	boxOfficeDay, err := model.GetTodayBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeDay = boxOfficeDay

	boxOfficeWeek, err := model.GetWeeklyBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeWeek = boxOfficeWeek

	boxOfficeMonth, err := model.GetMonthlyBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeMonth = boxOfficeMonth

	fmt.Printf("当日票房总数为: %.2f,  %.2f,  %.2f\n", boxOfficeDay, boxOfficeWeek, boxOfficeMonth)
}
