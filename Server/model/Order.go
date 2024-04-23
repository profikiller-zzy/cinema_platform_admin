package model

import (
	"gorm.io/gorm"
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
