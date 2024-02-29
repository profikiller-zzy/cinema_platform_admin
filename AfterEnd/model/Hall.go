package model

// Hall 电影院影厅
type Hall struct {
	MODEL
	HallName string `json:"hall_name"` // 影厅名称
	CinemaID uint   `json:"cinema_id"` // 与该影厅关联的影院
}
