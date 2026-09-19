package middleware

import (
	"strings"

	jwt "github.com/bluenotbloo/boys-help-boys/common/jwt"
	"github.com/bluenotbloo/boys-help-boys/gateway/config"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件用于验证 JWT token
func JWTAuth() gin.HandlerFunc {
	cfg := config.GetConfig()
	secret := []byte(cfg.JWT.Secret)

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			abortUnauthorized(c, "invalid auth format")
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwt.ParseToken(tokenStr, secret) // 解析token 验签
		if err != nil {
			abortUnauthorized(c, "invalid or expired token")
			return
		}

		// 将用户信息存入上下文，供后续处理使用
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Set("tenant_id", claims.TenantID)

		c.Next()
	}
}

// 便捷获取
func GetUserID(c *gin.Context) (uint64, bool) {
	v, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}
	id, ok := v.(uint64)
	return id, ok
}

func abortUnauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(401, gin.H{
		"code":    401,
		"message": msg,
	})
}
