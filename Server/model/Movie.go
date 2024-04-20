package model

import (
	"AfterEnd/global"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Movie 电影
type Movie struct {
	MODEL
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`              // 软删除字段
	MovieName       string         `json:"movie_name"`                           // 电影的名字
	CoverPictureID  uint           `json:"banner_id" structs:"banner_id"`        // 文章封面id
	CoverPictureUrl string         `json:"banner_url" structs:"banner_url"`      // 电影封面的存储地址
	ReleaseDate     time.Time      `json:"release_date"`                         // 上映时间
	PlayTime        time.Duration  `json:"play_time"`                            // 电影的时长
	Director        string         `json:"director"`                             // 电影的导演
	Actors          string         `json:"actors"`                               // 该电影的演员
	Reviews         []Review       `gorm:"foreignKey:MovieID" json:"reviews"`    // 该电影对应的评论
	Screenings      []Screening    `gorm:"foreignKey:MovieID" json:"screenings"` // 与该电影关联的排片信息
	//Actors          []Actor       `gorm:"many2many:movie_actor_perform" json:"actors"` // 该电影的演员表
}

// MovieCreate 完成电影入库操作
func (m *Movie) MovieCreate() (err error) {
	err = global.Db.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// SoftDeleteMovie 对电影进行软删除
func SoftDeleteMovie(movieID uint) (error, ListRemoveResponse) {
	// 根据用户ID查询用户
	var movie Movie
	var umRes ListRemoveResponse
	result := global.Db.Where("id = ?", movieID).First(&movie)
	err := result.Error
	// 删除失败
	if err != nil {
		umRes = ListRemoveResponse{
			UserID:    movie.ID,
			IsSuccess: false,
			Msg:       fmt.Sprintf("%s", err),
		}
		return result.Error, umRes
	}

	// 软删除电影
	result = global.Db.Delete(&movie)
	err = result.Error
	if err != nil {
		umRes = ListRemoveResponse{
			UserID:    movie.ID,
			IsSuccess: false,
			Msg:       fmt.Sprintf("%s", err),
		}
		return result.Error, umRes
	}

	umRes = ListRemoveResponse{
		UserID:    movie.ID,
		IsSuccess: true,
		Msg:       fmt.Sprintf("下架电影 %s 成功!", movie.MovieName),
	}
	return nil, umRes
}
