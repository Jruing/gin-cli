package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var mySecretKey = []byte("mysecretkey") // 这里的 secret key 用于签名 JWT，实际应用中要保密

// 用户登录后生成 JWT
func GenerateJWT(userId int, username string) (string, error) {
	// 创建一个 JWT token
	claims := jwt.MapClaims{
		"userId":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 secret key 签名并生成 token
	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析和验证 JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 使用的是 HS256 算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ok
		}
		return mySecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 获取 token 中的 claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
