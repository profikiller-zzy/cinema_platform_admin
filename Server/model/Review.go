package model

// Review 用户评论
type Review struct {
	MODEL
	Rating  int    `json:"rating"`   // 用户点赞数
	Comment string `json:"comment"`  // 用户评论内容
	UserID  uint   `json:"user_id"`  // 创建该评论的用户的用户ID
	MovieID uint   `json:"movie_id"` // 与该评论联系的电影ID
}
