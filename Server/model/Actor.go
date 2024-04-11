package model

import "time"

type Actor struct {
	MODEL
	ActorName string    `json:"actor_name"`                                  // 演员名称
	BirthDay  time.Time `json:"birth_day"`                                   // 演员生日
	Movies    []Movie   `gorm:"many2many:movie_actor_perform" json:"movies"` // 该演员演过的所有电影
}
