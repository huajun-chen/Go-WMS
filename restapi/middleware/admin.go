package middleware

import (
	"Go-WMS/global"
	"Go-WMS/param"
	"Go-WMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IsAdminAuth 用户权限中间件
// 参数：
//		无
// 返回值：
//		gin.HandlerFunc：Gin的处理程序
func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Token信息
		claims, _ := c.Get("claims")
		// 获取当前用户信息
		currentUser := claims.(*utils.CustomClaims)
		// 判断是否具有权限
		if currentUser.AuthorityID != 1 {
			failStruct := param.Resp{
				Code: 10015,
				Msg:  global.I18nMap["10015"],
			}
			c.JSON(http.StatusOK, failStruct)
			// 中断之后的中间件
			c.Abort()
			return
		}
		// 继续向下执行
		c.Next()
	}
}
