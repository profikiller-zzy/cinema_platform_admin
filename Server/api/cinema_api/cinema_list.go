package cinema_api

import (
	"AfterEnd/model"
	"AfterEnd/model/response"
	"AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
)

type CinemaListResponse struct {
	model.MODEL
	Name               string `json:"name"`                // 电影院名称
	AddressInformation string `json:"address_information"` // 地址信息
	TelNumber          string `json:"tel_number"`          // 电影院的联系电话
	CinemaUserID       uint   `json:"cinema_user_id"`      // 管理该电影院的影院用户
}

// CinemaListView 为管理员返回影院列表
func (CinemaApi) CinemaListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}
	var cinemaRepList = make([]CinemaListResponse, 0)
	var cinemaList = make([]model.Cinema, 0)
	var count int64
	// 对用户列表进行分页
	cinemaList, count, err = common_service.PagingList(model.Cinema{}, common_service.Option{
		PageInfo: pageModel,
		Likes:    []string{"name"},
	})

	for _, cinema := range cinemaList {
		cinemaRepList = append(cinemaRepList, CinemaListResponse{
			MODEL:              cinema.MODEL,
			Name:               cinema.Name,
			AddressInformation: cinema.AddressInformation,
			TelNumber:          cinema.TelNumber,
			CinemaUserID:       cinema.CinemaUserID,
		})
	}
	response.OKWithPagingData(cinemaRepList, count, c)
	return
}
