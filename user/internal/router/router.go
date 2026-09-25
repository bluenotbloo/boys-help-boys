package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 初始化路由
	router.POST("/user/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login success",
		})
	})

	router.POST("/user/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Register success",
		})
	})
}
