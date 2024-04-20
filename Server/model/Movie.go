package model

import (
	"AfterEnd/global"
	"time"
)

// Movie 电影
type Movie struct {
	MODEL
	MovieName       string        `json:"movie_name"`                           // 电影的名字
	CoverPictureID  uint          `json:"banner_id" structs:"banner_id"`        // 文章封面id
	CoverPictureUrl string        `json:"banner_url" structs:"banner_url"`      // 电影封面的存储地址
	ReleaseDate     time.Time     `json:"release_date"`                         // 上映时间
	PlayTime        time.Duration `json:"play_time"`                            // 电影的时长
	Director        string        `json:"director"`                             // 电影的导演
	Reviews         []Review      `gorm:"foreignKey:MovieID" json:"reviews"`    // 该电影对应的评论
	Screenings      []Screening   `gorm:"foreignKey:MovieID" json:"screenings"` // 与该电影关联的排片信息
	//Actors          []Actor       `gorm:"many2many:movie_actor_perform" json:"actors"` // 该电影的演员表
	Actors string `json:"actors"` // 该电影的演员
}

// MovieCreate 完成电影入库操作
func (m *Movie) MovieCreate() (err error) {
	err = global.Db.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}
