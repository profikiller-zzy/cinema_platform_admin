package middleware

import (
	"AfterEnd/model/response"
	"AfterEnd/service"
	"AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

// JwtAuth 很多操作只有用户(包括普通用户、影院用户和管理员)登录才能进行，该中间件用于验证用户的登录状态和权限
func JwtAuth(role int) gin.HandlerFunc {
	// 如何判断发送请求的是admin还是user
	// 从浏览器请求头中获取token，使用token判断是不是管理员
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.VerifyToken(tokenString)
		if err != nil {
			response.FailWithMessage("非法token", c)
			c.Abort()
			return
		}
		// 判断该token是否在redis黑名单中
		isInvalid, err := service.ServiceApp.UserServiceApp.CheckTokenInBlackList(tokenString)
		if isInvalid && err == nil {
			response.FailWithMessage("该token已失效，请重新登录", c)
			c.Abort()
			return
		}
		switch role {
		case 1:
			{ // 验证平台管理员
				if claims.JwtPayLoad.Role != 1 { // 该用户不是平台管理员
					response.FailWithMessage("您没有进行该操作的权限，请检查您的账户", c)
					c.Abort()
					return
					break
				}
			}
		case 2:
			{
				if claims.JwtPayLoad.Role != 2 { // 该用户不是影院用户
					response.FailWithMessage("您没有进行该操作的权限，请检查您的账户", c)
					c.Abort()
					return
					break
				}
			}
		default:
			break

		}
		c.Set("claims", claims)
	}
}
