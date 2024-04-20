package movie_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/service/image_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MovieCreateRequest struct {
	MovieName string `json:"movie_name" binding:"required" msg:"请输入电影名称"` // 电影的名字
	//CoverPictureID  uint          `json:"banner_id" structs:"banner_id"`   // 文章封面id
	//CoverPictureUrl string        `json:"banner_url" structs:"banner_url"` // 电影封面的存储地址
	ReleaseDate time.Time     `json:"release_date" binding:"required" msg:"请输入上映时间"` // 上映时间
	PlayTime    time.Duration `json:"play_time" binding:"required" msg:"请输入电影时长"`    // 电影的时长
	Director    string        `json:"director" binding:"required" msg:"请输入电影导演"`     // 电影的导演
	Actors      string        `json:"actors" binding:"required" msg:"请输入电影演员"`       // 电影的演员
}

func (MovieApi) MovieCreateView(c *gin.Context) {
	// 获取上传的封面图片
	fileHeader, err := c.FormFile("coverPicture")
	if err != nil {
		response.FailWithMessage("文件不存在", c)
		return
	}

	// 调用图片上传服务处理上传的文件
	upRes := image_service.ImageService{}.ImageUploadService(fileHeader, c)
	// 假如上传图片未成功且失败原因不是图片已存在
	if !upRes.IsSuccess && upRes.Msg != "图片已存在" {
		response.FailWithMessage(upRes.Msg, c)
		return
	}

	coverPictureID := upRes.PictureID
	coverPictureUrl := upRes.FilePath
	// 获取其他表单字段的值
	movieName := c.PostForm("movie_name")
	releaseDateStr := c.PostForm("release_date")
	playTimeStr := c.PostForm("play_time")
	director := c.PostForm("director")
	actors := c.PostForm("actors")

	// 将字符串转换为相应类型
	releaseDate, err := time.Parse(time.RFC3339, releaseDateStr)
	if err != nil {
		response.FailWithMessage("上映时间格式错误", c)
		return
	}
	playTime, err := time.ParseDuration(playTimeStr)
	if err != nil {
		response.FailWithMessage("电影时长格式错误", c)
		return
	}

	// 创建电影对象并存储到数据库
	err = global.Db.Create(&model.Movie{
		MovieName:       movieName,
		CoverPictureID:  coverPictureID,
		CoverPictureUrl: coverPictureUrl,
		ReleaseDate:     releaseDate,
		PlayTime:        playTime,
		Director:        director,
		Actors:          actors,
	}).Error

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("电影%s添加成功!", movieName), c)
	return
}
