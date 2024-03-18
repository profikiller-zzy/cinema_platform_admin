package model

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
