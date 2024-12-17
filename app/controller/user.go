package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)
import "gin-cli/app/service/user"

func CreateUser(c *gin.Context) {
	db := user.Queries{}
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
	db := user.Queries{}
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
	db := user.Queries{}
	err := db.DeleteUser(c, 1)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

func GetUserDetail(c *gin.Context) {
	db := user.Queries{}
	params := user.GetUserDetailParams{
		Column1: "",
		Column2: "",
		Limit:   0,
		Offset:  0,
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
		return
	}
	c.JSON(200, gin.H{
		"data":  detail,
		"count": count,
	})
}
