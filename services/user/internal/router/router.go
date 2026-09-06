package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化路由
	router := gin.Default()
	// 初始化路由
	router.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Register success",
		})
	})
}
