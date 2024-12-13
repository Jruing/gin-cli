package login

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"net/http"
	"time"
)

const (
	jwtSecret      = "your_secret_key" // 请替换为你的密钥
	accessTokenExp = 5 * time.Minute   // 访问 Token 过期时间
)

type Claims struct {
	UserID     int    `json:"userID"`
	Username   string `json:"username"`
	GrantScope string `json:"grantScope"`
	jwt.RegisteredClaims
}

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func generateTokenUsingHs256(userid int, username string) (string, error) {
	claim := Claims{
		UserID:     userid,
		Username:   username,
		GrantScope: "read_user_info",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jruing",                                           // 签发者
			Subject:   "zhangsan",                                         // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},         //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExp)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),    //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			ID:        randStr(10),                                        // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(jwtSecret))
	return token, err
}

func parseTokenHs256(token_string string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(token_string, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

// 用户结构体
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	// 绑定 JSON 请求到结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "root" && user.Password == "123456" {
		fmt.Println("密码校验成功")
		hs256, err := generateTokenUsingHs256(1, user.Username)
		if err != nil {
			return
		}
		fmt.Println("token===", hs256)
	} else {
		fmt.Println("密码校验失败")
	}
}
