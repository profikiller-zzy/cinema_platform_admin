package movie_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MovieEditRequest struct {
	MovieName   string        `json:"movie_name"`                                     // 电影的名字
	ReleaseDate time.Time     `json:"release_date"`                                   // 上映时间
	PlayTime    time.Duration `json:"play_time"`                                      // 电影的时长
	Director    string        `json:"director"`                                       // 电影的导演
	Actors      string        `json:"actors"`                                         // 该电影的演员
	MovieID     uint          `json:"movie_id" binding:"required" msg:"请确认需要修改的电影ID"` // 需要修改的电影ID
}

// MovieEditView 编辑电影信息
func (MovieApi) MovieEditView(c *gin.Context) {
	var meReq MovieEditRequest
	err := c.ShouldBindJSON(&meReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &meReq, c)
		return
	}

	var movieModel model.Movie
	err = global.Db.First(&movieModel, meReq.MovieID).Error
	if err != nil {
		response.FailWithMessage("该电影不存在，请检查电影ID是否正确", c)
		return
	}

	err = global.Db.Model(&movieModel).Updates(map[string]interface{}{
		"MovieName":   meReq.MovieName,
		"ReleaseDate": meReq.ReleaseDate,
		"PlayTime":    meReq.PlayTime,
		"Director":    meReq.Director,
		"Actors":      meReq.Actors,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("电影 %q 信息更新成功", meReq.MovieName), c)
}
