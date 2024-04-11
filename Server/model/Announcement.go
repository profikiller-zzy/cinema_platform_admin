package model

// Announcement 电影院发出的公告
type Announcement struct {
	MODEL
	CinemaID    uint   `json:"cinema_id"`   // 发出该公告的电影院
	Description string `json:"description"` // 公告的内容
}
