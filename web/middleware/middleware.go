package middleware

import (
	"bookstore-go/global"
	"bookstore-go/model"
	"net/http"
	"strings"

	"bookstore-go/jwt"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "认证令牌格式错误",
			})
			c.Abort()
			return
		}

		// 提取token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析JWT token
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "认证令牌无效",
			})
			c.Abort()
			return
		}

		// 检查token类型
		if claims.TokenType != "access" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "token类型错误，请使用access token",
			})
			c.Abort()
			return
		}

		// 查询用户信息
		var user model.User
		if err := global.DbClient.First(&user, int(claims.UserID)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "用户不存在",
			})
			c.Abort()
			return
		}

		// 检查是否为管理员
		if !user.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    -1,
				"message": "权限不足，需要管理员权限",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("admin_user", user)
		c.Set("admin_user_id", user.ID)
		c.Next()
	}
}
