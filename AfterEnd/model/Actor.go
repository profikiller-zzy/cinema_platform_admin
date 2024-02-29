package model

import "time"

type Actor struct {
	MODEL
	ActorName string    `json:"actor_name"` // 演员名称
	BirthDay  time.Time `json:"birth_day"`  // 演员生日
}
