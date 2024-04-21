package movie_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MovieCreateRequest struct {
	MovieName       string        `json:"movie_name" binding:"required" msg:"请输入电影名称"`   // 电影的名字
	CoverPictureID  uint          `json:"cover_picture_id" structs:"cover_picture_id"`   // 图片id
	CoverPictureUrl string        `json:"cover_picture_url" structs:"cover_picture_url"` // 电影封面的存储地址
	ReleaseDate     time.Time     `json:"release_date" binding:"required" msg:"请输入上映时间"` // 上映时间
	PlayTime        time.Duration `json:"play_time" binding:"required" msg:"请输入电影时长"`    // 电影的时长
	Director        string        `json:"director" binding:"required" msg:"请输入电影导演"`     // 电影的导演
	Actors          string        `json:"actors" binding:"required" msg:"请输入电影演员"`       // 电影的演员
}

func (MovieApi) MovieCreateView(c *gin.Context) {
	var mcReq MovieCreateRequest
	if err := c.ShouldBindJSON(&mcReq); err != nil {
		response.FailBecauseOfParamError(err, &mcReq, c)
		return
	}

	// 创建电影对象并存储到数据库
	err := global.Db.Create(&model.Movie{
		MovieName:       mcReq.MovieName,
		CoverPictureID:  mcReq.CoverPictureID,
		CoverPictureUrl: mcReq.CoverPictureUrl,
		ReleaseDate:     mcReq.ReleaseDate,
		PlayTime:        mcReq.PlayTime,
		Director:        mcReq.Director,
		Actors:          mcReq.Actors,
	}).Error

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("电影%s添加成功!", mcReq.MovieName), c)
	return
}
