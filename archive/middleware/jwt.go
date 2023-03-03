package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	util "github.com/yby/todo_list/pkg/util/jwt"
)

// JWT JWT中间件
func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取token
		token := c.Query("token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "验证过期",
			})
			c.Abort()
			return
		}

		// 解析Token
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "验证过期",
			})
			c.Abort()
			return
		}

		// 验证
		now := time.Now()
		if now.After(claims.ExpiresAt.Time) || claims.Issuer != "Todo-List" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "验证过期",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
