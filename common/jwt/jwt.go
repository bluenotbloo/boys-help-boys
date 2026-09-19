// token 签发/验证
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 业务自定义载荷
type CustomClaims struct {
	UserID               uint64   `json:"user_id"`             // 用户 ID
	Username             string   `json:"username"`            // 用户名
	Roles                []string `json:"roles,omitempty"`     // 扩展：RBAC
	TenantID             string   `json:"tenant_id,omitempty"` // 扩展：多租户
	jwt.RegisteredClaims          // 内嵌标准字段：exp, iat, nbf, iss, sub, aud, jti
}

// 颁发 token
//
// userID: 用户 ID
// username: 用户名
// roles: 用户角色列表
// tenantID: 租户 ID
// issuer: 签发者
// subject: 主题
// secret: 签名密钥
// expiresIn: 过期时间
func GenerateToken(
	userID uint64,
	username string,
	roles []string,
	tenantID string,
	issuer string,
	subject string,
	secret []byte,
	expiresIn time.Duration,
) (string, error) {
	claims := CustomClaims{
		UserID:   userID,   // 用户 ID
		Username: username, // 用户名
		Roles:    roles,    // 角色
		TenantID: tenantID, // 租户 ID
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                // 生效时间
			Issuer:    issuer,                                        // 签发者
			Subject:   subject,                                       // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用 HS256 签名算法
	return token.SignedString(secret)                          // 返回签名后的 token 字符串
}

// 解析并校验 token
//
// tokenStr: JWT token 字符串
// secret: 签名密钥
func ParseToken(tokenStr string, secret []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,        // 解析 token 字符串
		&CustomClaims{}, // 解析为自定义claims
		func(t *jwt.Token) (interface{}, error) { // 回调函数用于提供签名密钥
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { //  检查签名算法是否为 HMAC
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
