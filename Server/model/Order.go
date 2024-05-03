package model

import (
	"AfterEnd/global"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// Order 订单
type Order struct {
	MODEL
	TicketPrice float64 `json:"ticket_price"` // 订单票价
	UserID      uint    `json:"user_id"`      // 外键，下订单的用户
	ScreeningID uint    `json:"screening_id"` // 电影的排片信息ID
}

// GetTodayBoxOffice 统计今日票房
func GetTodayBoxOffice(db *gorm.DB) (float64, error) {
	var totalBoxOffice float64
	// 获取当日的日期
	today := time.Now().Format("2006-01-02")
	// 查询当日的所有排片信息
	var screeningsDay []Screening
	err := db.Where("DATE(start_time) = ?", today).Preload("Orders").Find(&screeningsDay).Error
	if err != nil {
		return 0, err
	}

	// 计算本周票房
	for _, screening := range screeningsDay {
		for _, order := range screening.Orders {
			totalBoxOffice += order.TicketPrice
		}
	}

	return totalBoxOffice, nil
}

// GetWeeklyBoxOffice 统计本周票房
func GetWeeklyBoxOffice(db *gorm.DB) (float64, error) {
	var totalBoxOffice float64
	// 获取本周的起始日期和结束日期
	startOfWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday())+1)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	// 查询本周内的排片信息
	var screenings []Screening
	if err := db.Where("DATE(start_time) BETWEEN ? AND ?", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02")).Preload("Orders").Find(&screenings).Error; err != nil {
		return 0, err
	}

	// 计算本周票房
	for _, screening := range screenings {
		for _, order := range screening.Orders {
			totalBoxOffice += order.TicketPrice
		}
	}

	return totalBoxOffice, nil
}

// GetMonthlyBoxOffice 统计本月票房
func GetMonthlyBoxOffice(db *gorm.DB) (float64, error) {
	var totalBoxOffice float64
	// 获取本月的起始日期和结束日期
	startOfMonth := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	// 查询本月内的排片信息
	var screenings []Screening
	if err := db.Where("DATE(start_time) BETWEEN ? AND ?", startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02")).Preload("Orders").Find(&screenings).Error; err != nil {
		return 0, err
	}

	// 计算本月票房
	for _, screening := range screenings {
		for _, order := range screening.Orders {
			totalBoxOffice += order.TicketPrice
		}
	}

	return totalBoxOffice, nil
}

func CountOrders(db *gorm.DB) (int64, error) {
	var count int64
	result := db.Model(Order{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// CalculateSevenDaysStatistics 函数用于返回近七天内每天的票房数和订单数
func CalculateSevenDaysStatistics() ([]float64, []int64, error) {
	// 获取当前时间
	now := time.Now()

	// 初始化结果切片
	boxOfficePerDay := make([]float64, 7)
	ordersPerDay := make([]int64, 7)

	// 循环计算近七天的数据
	for i := 0; i < 7; i++ {
		// 计算当前日期之前的第 i 天的日期
		targetDate := now.AddDate(0, 0, -i)

		// 查询当天的排片信息
		var screenings []Screening
		if err := global.Db.Where("DATE(start_time) = ?", targetDate.Format("2006-01-02")).Find(&screenings).Error; err != nil {
			return nil, nil, err
		}

		// 初始化当天的票房和订单数量
		totalBoxOffice := 0.0
		totalOrders := int64(0)

		// 遍历当天的排片信息，计算票房和订单数量
		for _, screening := range screenings {
			// 假如当天没有排片
			if screening.ID == 0 {
				continue
			}

			// 查询当天该排片对应的订单数量和票房
			var ordersCount int64
			var totalTicketPrice float64
			if err := global.Db.Model(&Order{}).
				Where("screening_id = ?", screening.ID).
				Count(&ordersCount).
				Select("COALESCE(SUM(ticket_price), 0)"). // 使用 COALESCE 函数将 NULL 值转换为 0
				Scan(&totalTicketPrice).Error; err != nil {
				return nil, nil, err
			}

			// 累加票房和订单数量
			totalBoxOffice += totalTicketPrice
			totalOrders += ordersCount
		}

		// 将当天的票房和订单数量存入结果切片
		boxOfficePerDay[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", totalBoxOffice), 64)
		ordersPerDay[i] = totalOrders
	}

	return boxOfficePerDay, ordersPerDay, nil
}
