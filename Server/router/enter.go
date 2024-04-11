package router

import (
	"AfterEnd/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		global.Log.Warnln(err.Error())
	}

	router.StaticFS("uploads", http.Dir("uploads"))

	// 该路由组是管理员的路由组
	apiRouter := router.Group("/api/")

	apiGroupApp := RGroup{
		RouterGroup: apiRouter,
	}
	apiGroupApp.UserRouter()
	apiGroupApp.CinemaRouter()

	return router
}
