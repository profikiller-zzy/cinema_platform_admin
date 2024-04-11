package cinema_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type ID struct {
	ID uint `form:"id"` // 单个用户的id
}

// IndividualApprovalView
func (CinemaApi) IndividualApprovalView(c *gin.Context) {
	var id ID
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	var cinemaApp model.CinemaApproval
	err = global.Db.Take(&cinemaApp, "id = ?", id.ID).Error
	if err != nil {
		response.FailWithMessage("未查找到响应的审核信息", c)
		return
	}
	var unAppRep UnreviewedApprovalResponse

	unAppRep = UnreviewedApprovalResponse{
		MODEL:              cinemaApp.MODEL,
		Name:               cinemaApp.Name,
		AddressInformation: cinemaApp.AddressInformation,
		TelNumber:          cinemaApp.TelNumber,
		CinemaUserID:       cinemaApp.CinemaUserID,
	}
	response.OKWithData(unAppRep, c)
	return
}
