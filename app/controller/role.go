package controller

import (
	"gin-cli/app/service/role"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateRole(c *gin.Context) {
	db := role.Queries{}
	params := role.CreateRoleParams{
		Rolename: "",
		Status:   0,
		Created:  pgtype.Timestamp{},
	}
	_, err := db.CreateRole(c, params)
	if err != nil {
		return
	}
}

func UpdateRole(c *gin.Context) {
	db := role.Queries{}
	params := role.UpdateRoleParams{
		ID:       0,
		Rolename: "",
		Status:   0,
		Created:  pgtype.Timestamp{},
	}
	err := db.UpdateRole(c, params)
	if err != nil {
		return
	}
}

func DeleteRole(c *gin.Context) {
	db := role.Queries{}
	err := db.DeleteRole(c, 1)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

func GetRoleDetail(c *gin.Context) {
	db := role.Queries{}
	params := role.GetRoleDetailParams{
		Limit:  0,
		Offset: 0,
	}
	detail, err := db.GetRoleDetail(c, params)
	if err != nil {
		return
	}
	count_params := role.GetRoleDetailParams{
		Limit:  0,
		Offset: 0,
	}
	count, err := db.GetRoleDetail(c, count_params)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data":  detail,
		"count": count,
	})
}
