package flag

import (
	"AfterEnd/global"
	"AfterEnd/model"
)

// AutoMigration 对mysql数据库的schema实现自动迁移
func AutoMigration() {
	var err error
	// 对scheme自动迁移
	// 自定义多对多关系表
	//err = global.Db.SetupJoinTable(&model.Movie{}, "Actors", &model.MovieActorPerform{})
	//if err != nil {
	//	global.Log.Error(err.Error())
	//	return
	//}
	//err = global.Db.SetupJoinTable(&model.Actor{}, "Movies", &model.MovieActorPerform{})
	//if err != nil {
	//	global.Log.Error(err.Error())
	//	return
	//}
	err = global.Db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
			&model.Announcement{},
			&model.Review{},
			&model.Cinema{},
			&model.Movie{},
			//&model.Actor{},
			&model.Order{},
			&model.Hall{},
			&model.Screening{},
			//&model.MovieActorPerform{},
			&model.CinemaApproval{},
			&model.ProfilePhoto{})
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("数据表迁移成功！")
}
