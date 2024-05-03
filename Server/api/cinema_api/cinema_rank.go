package cinema_api

import (
	"AfterEnd/global"
	"AfterEnd/model"
	"AfterEnd/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

// CinemaRank 返回的影院排行榜
type CinemaRank struct {
	Rank       uint    `json:"rank"`
	CinemaID   uint    `json:"cinema_id"`
	CinemaName string  `json:"cinema_name"`
	Revenue    float64 `json:"revenue"`
}

// CinemaBoxOfficeRanking 需要将影院以最近一周营收总数进行排行
func (CinemaApi) CinemaBoxOfficeRanking(c *gin.Context) {
	// TODO 查询出给定时间区间中的电影票房排行前100名
	DB := global.Db.Session(&gorm.Session{Logger: global.MysqlLog})
	cinemasInfo, err := model.GetAllCinema(DB)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询出错，err：%s", err.Error()), c)
		return
	}

	rankings := make([]CinemaRank, 0, 100) // 开辟长度为 100 的切片，但初始长度为 0
	for _, value := range cinemasInfo {
		revenue, err := model.RevenueOfCinemaInRecentWeek(value.ID, DB)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("在计算影院 '%s' 的票房时出错！, err : %s", value.Name, err.Error()), c)
			return
		}
		// 如果票房为 0，则跳过该电影
		if revenue == 0 {
			continue
		}
		rank := CinemaRank{
			CinemaID:   value.ID,
			CinemaName: value.Name,
			Revenue:    revenue,
		}
		rankings = append(rankings, rank)
	}

	// 对排行榜按票房总额降序排序
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Revenue > rankings[j].Revenue
	})

	// 如果排行榜的长度大于 100，则截断切片
	if len(rankings) > 100 {
		rankings = rankings[:100]
	}

	// 填入排名
	for i := 0; i < len(rankings); i++ {
		rankings[i].Rank = uint(i + 1)
	}

	// 返回结果
	response.OKWithData(rankings, c)
}
