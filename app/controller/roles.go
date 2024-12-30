package controller

import (
	"fmt"
	"gin-cli/app/dal/model"
	"gin-cli/app/dal/query"
	"gin-cli/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.BindJSON(&jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.Role{}
	if rolename, ok := jsondata["rolename"]; ok {
		params.Rolename = rolename.(string)
	}
	if status, ok := jsondata["status"]; ok {
		params.Status = int32(status.(float64))
	}
	params.Created = time.Now()
	err := query.Role.WithContext(c).Create(&params)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "角色新增成功",
	})
}

func UpdateRole(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.Role{}
	if id, ok := jsondata["id"]; ok {
		params.ID = id.(int64)
	}
	if rolename, ok := jsondata["rolename"]; ok {
		params.Rolename = rolename.(string)
	}
	if status, ok := jsondata["status"]; ok {
		params.Status = status.(int32)
	}
	info, err := query.Role.WithContext(c).Where(query.Role.ID.Eq(params.ID)).Updates(&params)
	if err != nil {
		panic(err)
	}
	fmt.Println("更新条数", info.RowsAffected)
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "角色信息更新成功",
		"data": info.RowsAffected,
	})
}

func DeleteRole(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)

	var roleid int64
	if id, ok := jsondata["id"]; ok {
		roleid = int64(id.(float64))
	}
	info, err := query.Role.WithContext(c).Where(query.Role.ID.Eq(roleid)).Delete()
	if err != nil {
		return
	}
	fmt.Println("删除条数", info.RowsAffected)
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "角色删除成功",
		"data": info,
	})
}

func GetRoleDetail(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)
	q := query.Role.WithContext(c)
	for key, value := range jsondata {
		switch key {
		case "rolename":
			q.Where(query.Role.Rolename.Eq(value.(string)))
		case "status":
			q.Where(query.Role.Status.Eq(int32(value.(float64))))
		}
	}
	var Limit int
	if limit, ok := jsondata["limit"]; ok {
		Limit = int(limit.(float64))
		q.Limit(Limit)
	}
	if page, ok := jsondata["page"]; ok {
		q.Offset((int(page.(float64)) - 1) * Limit)
	}
	find, err := q.Select(query.Role.ALL).Find()
	if err != nil {
		return
	}
	count, err := q.Count()
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  utils.Success,
		"msg":   "角色列表查询成功",
		"data":  find,
		"count": count,
	})
}
