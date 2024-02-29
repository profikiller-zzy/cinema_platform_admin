package model

import "time"

// Movie 电影
type Movie struct {
	MODEL
	CNName      string        `json:"cn_name"`      // 电影的中文名
	ENName      string        `json:"en_name"`      // 电影的英文名
	ReleaseDate time.Time     `json:"release_date"` // 上映时间
	PlayTime    time.Duration `json:"play_time"`    // 电影的时长
	Director    string        `json:"director"`     // 电影的导演
}
