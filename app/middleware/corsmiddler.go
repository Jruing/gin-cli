package middleware

import (
	"gin-cli/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 中间件处理跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 验证用户状态
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/v1/login" || c.Request.URL.Path == "/api/v1/register" {
			c.Next()
			return
		} else {
			token := c.GetHeader("Authorization")
			if token == "" {
				c.JSON(http.StatusOK, gin.H{
					"code": utils.TokenError,
					"msg":  "token不能为空",
				})
				c.Abort()
				return
			} else {
				// 验证 JWT Token
				claims, err := utils.ValidateJWT(token)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": utils.TokenError, "msg": "token验证失败"})
					c.Abort()
					return
				}

				// 把解析出的 userId 和 username 放到上下文中
				c.Set("userId", claims["userId"])
				c.Set("username", claims["username"])

				c.Next()
			}
		}
	}
}
