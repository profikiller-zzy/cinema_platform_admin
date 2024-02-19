package model

type Cinema struct {
	MODEL
	Name               string `gorm:"size:64" json:"name"`                 // 电影院名称
	AddressInformation string `gorm:"size:128" json:"address_information"` // 地址信息
	Director           string `gorm:"size:16" json:"director"`             // 导演
	CinemaUserID       uint   `gorm:""`
}
