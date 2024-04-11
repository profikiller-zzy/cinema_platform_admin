package cinema_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UnreviewedApprovalResponse struct {
	model.MODEL
	Name               string `json:"name"`                // 电影院名称
	AddressInformation string `json:"address_information"` // 地址信息
	TelNumber          string `json:"tel_number"`          // 电影院的联系电话
	CinemaUserID       uint   `json:"cinema_user_id"`      // 提交该资料的用户
}

type UnreviewedApprovalListRequest struct {
	model.PageInfo
	Role int `json:"role" form:"role"`
}

// UnreviewedApprovalListView 为管理员返回未审批的影院资料列表
func (CinemaApi) UnreviewedApprovalListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	var unreviewedAppList = make([]UnreviewedApprovalResponse, 0)
	var ApprovalList = make([]model.CinemaApproval, 0)
	offset := pageModel.PageSize * (pageModel.PageNum - 1)

	var count int64
	global.Db.Model(model.CinemaApproval{}).Where("state = ?", "审核中").Count(&count)

	// 对未审批的影院资料进行查询和分页
	err = global.Db.Where("state = ?", "审核中").Offset(offset).Limit(pageModel.PageSize).Order("create_at asc").Find(&ApprovalList).Error
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("error: %s", err.Error()), c)
		return
	}

	for _, approval := range ApprovalList {
		unreviewedAppList = append(unreviewedAppList, UnreviewedApprovalResponse{
			MODEL:              approval.MODEL,
			Name:               approval.Name,
			AddressInformation: approval.AddressInformation,
			TelNumber:          approval.TelNumber,
			CinemaUserID:       approval.CinemaUserID,
		})
	}

	response.OKWithPagingData(unreviewedAppList, count, c)
}
