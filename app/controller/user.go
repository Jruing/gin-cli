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
func CreateUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	params := user.CreateUserParams{
		Nickname: pgtype.Text{},
		Username: "",
		Password: "",
		Sex:      pgtype.Text{},
		Email:    "",
		Status:   0,
		Created:  pgtype.Timestamp{},
	}
	_, err := db.CreateUser(c, params)
	if err != nil {
		return
	}
}

func UpdateUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	params := user.UpdateUserParams{
		ID:       0,
		Nickname: pgtype.Text{},
		Username: "",
		Password: "",
		Sex:      pgtype.Text{},
		Email:    "",
		Status:   0,
	}
	err := db.UpdateUser(c, params)
	if err != nil {
		return
	}
}

func DeleteUser(c *gin.Context) {
	db := user.New(utils.Pgconn)
	err := db.DeleteUser(c, 1)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

func GetUserDetail(c *gin.Context) {
	nickname := c.Param("nickname")
	username := c.Param("username")
	page := c.Param("page")
	limit := c.Param("limit")
	limit1, err := strconv.ParseInt(limit, 0, 32)
	page1, err := strconv.ParseInt(page, 0, 32)

	db := user.New(utils.Pgconn)
	params := user.GetUserDetailParams{
		Column1: nickname,
		Column2: username,
		Limit:   limit1,
		Offset:  (page1 - 1) * limit1,
	}
	detail, err := db.GetUserDetail(c, params)
	if err != nil {
		return
	}
	count_params := user.GetUserCountParams{
		Column1: "",
		Column2: "",
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
