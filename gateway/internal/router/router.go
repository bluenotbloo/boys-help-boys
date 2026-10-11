package router

import (
	"net/http/httputil"
	"net/url"

	"github.com/bluenotbloo/boys-help-boys/gateway/config"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 占位路由
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

	// 用户服务反向代理：/api/v1/user/* → user 服务的 grpc-gateway HTTP 端口
	userProxy := newUserServiceProxy()
	router.Any("/api/v1/user/*path", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})
}

// newUserServiceProxy 创建指向 user 服务的反向代理
func newUserServiceProxy() *httputil.ReverseProxy {
	cfg := config.GetConfig()
	target, err := url.Parse("http://" + cfg.Upstream.UserService)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(target)
}