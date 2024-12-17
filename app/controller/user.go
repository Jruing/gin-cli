package controller

import (
	"gin-cli/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"strconv"
)
import "gin-cli/app/service/user"

// @Summary 创建用户
// @Description 创建用户
// @ID create-user
// @Accept  json
// @Produce  json
// @Param nickname form string true "昵称"
// @Param username form string true "用户名"
// @Param password form string true "密码"
// @Success 200 {object} controllers.user
// @Router /users/createuser [post]

// 新增用户
func CreateUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	params := user.CreateUserParams{
		Nickname: pgtype.Text{},
		Username: json["username"].(string),
		Password: json["password"].(string),
		Sex:      pgtype.Text{},
		Email:    json["email"].(string),
		Status:   json["staus"].(int32),
		Created:  pgtype.Timestamp{},
	}
	err := db.CreateUser(c, params)
	if err != nil {
		return
	}

}

// 更新用户信息
func UpdateUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	id := c.Param("id")
	userid, _ := strconv.ParseInt(id, 10, 32)
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	params := user.UpdateUserParams{
		ID:       int32(userid),
		Nickname: pgtype.Text{},
		Username: json["username"].(string),
		Password: json["password"].(string),
		Sex:      pgtype.Text{},
		Email:    json["email"].(string),
		Status:   json["status"].(int32),
	}
	err = db.UpdateUser(c, params)
	if err != nil {
		return
	}
}

// 删除用户
func DeleteUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	id := c.Param("id")
	userid, _ := strconv.ParseInt(id, 10, 32)
	err := db.DeleteUser(c, int32(userid))
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}

// 获取用户详情
func GetUserDetail(c *gin.Context) {
	nickname := c.Param("nickname")
	username := c.Param("username")
	page := c.Param("page")
	limit := c.Param("limit")
	limit1, err := strconv.ParseInt(limit, 10, 32)
	page1, err := strconv.ParseInt(page, 10, 32)

	db := user.New(utils.Pgconn)
	params := user.GetUserDetailParams{
		Column1: nickname,
		Column2: username,
		Limit:   int32(limit1),
		Offset:  (int32(page1) - 1) * int32(limit1),
	}
	detail, err := db.GetUserDetail(c, params)
	if err != nil {
		return
	}
	count_params := user.GetUserCountParams{
		Column1: nickname,
		Column2: username,
	}
	count, err := db.GetUserCount(c, count_params)
	if err != nil {
		c.JSON(0, gin.H{
			"data": []string{},
			"msg":  "未查询到数据",
		})
	}
	c.JSON(200, gin.H{
		"data":  detail,
		"count": count,
	})
}
