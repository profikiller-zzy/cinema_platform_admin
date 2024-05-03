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
	TotalUsers     int64   `json:"total_users"`
	TotalOrders    int64   `json:"total_orders"`
}

func (DataApi) DataStatistics(c *gin.Context) {
	var dsRes dataStatisticsResponse
	DB := global.Db.Session(&gorm.Session{Logger: global.MysqlLog})

	// 统计当日票房
	boxOfficeDay, err := model.GetTodayBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("统计当日票房出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeDay = boxOfficeDay

	// 统计本周票房
	boxOfficeWeek, err := model.GetWeeklyBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("统计本周票房出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeWeek = boxOfficeWeek

	// 统计本月票房
	boxOfficeMonth, err := model.GetMonthlyBoxOffice(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("统计本月票房出错，%s", err.Error()), c)
		return
	}
	dsRes.BoxOfficeMonth = boxOfficeMonth

	// 统计用户总数
	totalUser, err := model.CountUsers(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("统计用户总数出错，%s", err.Error()), c)
		return
	}
	dsRes.TotalUsers = totalUser

	// 统计订单总数
	totalOrder, err := model.CountOrders(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("统计订单总数出错，%s", err.Error()), c)
		return
	}
	dsRes.TotalOrders = totalOrder

	response.OKWithData(dsRes, c)
}
