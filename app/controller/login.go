package controller

import (
	"gin-cli/app/dal/query"
	"gin-cli/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"user1": "password1", // 模拟一个用户数据库，用户名:密码
	"user2": "password2",
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	query.SetDefault(utils.MysqlClient)
	q := query.User.WithContext(c)
	user, err := q.Where(query.User.Username.Eq(req.Username),query.User.Password.Eq(req.Password)).First()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}else {
		token, err := utils.GenerateJWT(int(user.ID), user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	// // 验证用户名和密码
	// if password, exists := users[req.Username]; exists && password == req.Password {
	// 	// 用户验证通过，生成 JWT
	// 	token, err := utils.GenerateJWT(1, req.Username) // 假设用户 ID 为 1
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
	// 		return
	// 	}

	// 	// 返回 JWT 给客户端
	// 	c.JSON(http.StatusOK, gin.H{"token": token})
	// } else {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	// }
}
