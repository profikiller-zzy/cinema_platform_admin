package model

// Order 订单
type Order struct {
	MODEL
	TicketPrice float32 `json:"ticket_price"`      // 订单票价
	UserID      uint    `json:"user_id"`           // 外键，下订单的用户
	User        User    `gorm:"foreignKey:UserID"` // 与该订单关联的用户
	ScreeningID uint    `json:"screening_id"`      // 电影的排片信息ID
}
