package model

import "AfterEnd/model/ctype"

// CinemaApproval 电影院临时审批资料
type CinemaApproval struct {
	MODEL
	Name               string      `gorm:"size:64" json:"name"`                 // 电影院名称
	AddressInformation string      `gorm:"size:128" json:"address_information"` // 地址信息
	TelNumber          string      `gorm:"size:32" json:"tel_number"`           // 电影院的联系电话
	CinemaUserID       uint        `json:"cinema_user_id"`                      // 提交该资料的用户
	State              ctype.State `gorm:"size:16" json:"state"`                // 审批资料的状态
}
