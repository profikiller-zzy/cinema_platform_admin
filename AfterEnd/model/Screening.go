package model

import "time"

// Screening 排片信息
type Screening struct {
	ID          uint      `gorm:"primaryKey" json:"id"`                 // 主键ID
	MovieID     uint      `json:"movie_id"`                             // 与之关联的电影
	CinemaID    uint      `json:"cinema_id"`                            // 与之关联的影院
	HallID      uint      `json:"hall_id"`                              // 与之关联的影厅
	StartTime   time.Time `json:"start_time"`                           // 开始时间
	EndTime     time.Time `json:"end_time"`                             // 预计散场时间
	TicketPrice float32   `json:"ticket_price"`                         // 票价
	Language    string    `json:"language"`                             // 语言类型
	Orders      []Order   `gorm:"foreignKey:ScreeningID" json:"orders"` // 该排片对应的订单记录
}
