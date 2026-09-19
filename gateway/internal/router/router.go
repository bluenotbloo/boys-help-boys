package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 初始化路由
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")
		c.JSON(200, gin.H{
			"message": "Hello, " + path + "!",
		})
	})
}
