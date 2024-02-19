package model

import "time"

// MODEL 通用模板
type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`            // 主键ID
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"` // 更新时间
}
