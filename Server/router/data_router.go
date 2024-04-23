package router

import "AfterEnd/api"

func (r RGroup) DataRouter() {
	dataApiApp := api.ApiGroupApp.DataApi
	r.GET("/data_statistics/", dataApiApp.DataStatistics)
}
