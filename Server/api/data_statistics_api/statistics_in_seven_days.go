package data_statistics_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type StatisticsInSevenDaysResponse struct {
	BoxOfficePerDay []float64 `json:"box_office_per_day"`
	OrdersPerDay    []int64   `json:"orders_per_day"`
	Date            []string  `json:"date"`
}

func (DataApi) StatisticsInSevenDays(c *gin.Context) {
	var res StatisticsInSevenDaysResponse
	boxOfficePerDay, ordersPerDay, err := model.CalculateSevenDaysStatistics()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("服务器出错, err: %s", err.Error()), c)
		return
	}
	dateStringList := GetSevenDaysDateList()

	res = StatisticsInSevenDaysResponse{
		BoxOfficePerDay: boxOfficePerDay,
		OrdersPerDay:    ordersPerDay,
		Date:            dateStringList,
	}
	response.OKWithData(res, c)
}

// GetSevenDaysDateList 函数用于返回最近七天的日期字符串列表，日期靠前的排在前面
func GetSevenDaysDateList() []string {
	// 获取当前时间
	now := time.Now()
	// 初始化结果切片
	dateList := make([]string, 7)
	// 循环计算最近七天的日期字符串
	for i := 6; i >= 0; i-- { // 从6开始，递减到0
		// 计算当前日期之前的第 i 天的日期
		targetDate := now.AddDate(0, 0, -i)
		// 格式化日期字符串为"月日"形式，例如"4月23日"
		dateStr := targetDate.Format("1月2日")
		// 将日期字符串存入结果切片
		dateList[6-i] = dateStr // 将日期靠前的放在前面
	}
	return dateList
}
