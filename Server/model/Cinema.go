package model

import (
	"gorm.io/gorm"
)

// Cinema 电影院
type Cinema struct {
	MODEL
	Name               string         `gorm:"size:64" json:"name"`                 // 电影院名称
	AddressInformation string         `gorm:"size:128" json:"address_information"` // 地址信息
	TelNumber          string         `gorm:"size:32" json:"tel_number"`           // 电影院的联系电话
	CinemaUserID       uint           `json:"cinema_user_id"`                      // 管理该电影院的影院用户
	Announcements      []Announcement `json:"announcements"`                       // 该电影院发出的公告
	Halls              []Hall         `gorm:"foreignKey:CinemaID" json:"halls"`    // 该电影院的影厅
}

type CinemaInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetAllCinema 获取所有的电影院
func GetAllCinema(db *gorm.DB) ([]CinemaInfo, error) {
	var cinemas []CinemaInfo

	// 查询电影的ID和名称
	if err := db.Model(&Cinema{}).Select("id, name").Find(&cinemas).Error; err != nil {
		return nil, err
	}

	return cinemas, nil
}

// RevenueOfCinemaInRecentWeek 根据电影ID，查询出该电影特定时间范围内的总票房
func RevenueOfCinemaInRecentWeek(cinemaID uint, db *gorm.DB) (float64, error) {
	// 先将目标电影院院所拥有的所有影厅ID查出来
	var hallIDs []uint
	if err := db.Model(&Hall{}).Where("cinema_id = ?", cinemaID).Pluck("id", &hallIDs).Error; err != nil {
		return 0, err
	}

	// 将排片在这些影厅中的排片信息查出来
	var screenings []Screening
	if err := db.Where("hall_id IN ?", hallIDs).Preload("Orders").Find(&screenings).Error; err != nil {
		return 0, err
	}

	var revenue float64
	// 计算票房
	for _, screening := range screenings {
		for _, order := range screening.Orders {
			revenue += order.TicketPrice
		}
	}
	return revenue, nil
}
