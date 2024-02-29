package model

type MovieActorPerform struct {
	ActorID uint `gorm:"primaryKey" json:"actor_id"` // 演员ID
	MovieID uint `gorm:"primaryKey" json:"movie_id"` // 电影ID
}
