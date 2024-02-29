package model

import (
	"AfterEnd/global"
	"time"
)

// MODEL 通用模板
type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`            // 主键ID
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"` // 更新时间
}

func AutoMigration() {
	var err error
	// 自定义多对多关系表
	//err = global.Db.SetupJoinTable(&model.MenuModel{}, "Banners", &model.MenuBanner{})
	//if err != nil {
	//	global.Log.Warn(err.Error())
	//}
	// 对scheme自动迁移
	err = global.Db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&User{},
			&Announcement{},
			&Review{},
			&Cinema{},
			&Movie{},
			&Actor{},
			&Order{},
			&Hall{},
			&Screening{})
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("数据表迁移成功！")
}
