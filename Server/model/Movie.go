package model

import "time"

// Movie 电影
type Movie struct {
	MODEL
	CNName      string        `json:"cn_name"`                                     // 电影的中文名
	ENName      string        `json:"en_name"`                                     // 电影的英文名
	ReleaseDate time.Time     `json:"release_date"`                                // 上映时间
	PlayTime    time.Duration `json:"play_time"`                                   // 电影的时长
	Director    string        `json:"director"`                                    // 电影的导演
	Reviews     []Review      `gorm:"foreignKey:MovieID" json:"reviews"`           // 该电影对应的评论
	Screenings  []Screening   `gorm:"foreignKey:MovieID" json:"screenings"`        // 与该电影关联的排片信息
	Actors      []Actor       `gorm:"many2many:movie_actor_perform" json:"actors"` // 该电影的演员表
}
