package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yby/todo_list/service"
)

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	userID := c.Query("user_id")
	password := c.Query("password")

	token, err := service.LoginUser(userID, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"token": token,
	})
}

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("password")

	uid, token, err := service.RegisterUser(pwd, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "ok",
		"user_id": uid,
		"token":   token,
	})
}
