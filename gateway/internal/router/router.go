package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化路由
	router := gin.Default()
	// 初始化路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}
