package controller

import (
	"fmt"
	"gin-cli/app/dal/model"
	"gin-cli/app/dal/query"
	"gin-cli/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.BindJSON(&jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.User{}

	if username, ok := jsondata["username"]; ok {
		params.Username = username.(string)
	}
	if nickname, ok := jsondata["nickname"]; ok {
		params.Nickname = nickname.(string)
	}
	if password, ok := jsondata["password"]; ok {
		params.Password = password.(string)
	}
	if email, ok := jsondata["email"]; ok {
		params.Email = email.(string)
	}
	if sex, ok := jsondata["sex"]; ok {
		params.Sex = sex.(string)
	}
	if status, ok := jsondata["status"]; ok {
		params.Status = int32(status.(float64))
	}
	params.Created = time.Now()
	err := query.User.WithContext(c).Create(&params)

	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "用户新增成功",
	})
}

func UpdateUser(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.User{}
	if id, ok := jsondata["id"]; ok {
		params.ID = id.(int64)
	}
	if username, ok := jsondata["username"]; ok {
		params.Username = username.(string)
	}
	if nickname, ok := jsondata["nickname"]; ok {
		params.Nickname = nickname.(string)
	}
	if password, ok := jsondata["password"]; ok {
		params.Password = password.(string)
	}
	if email, ok := jsondata["email"]; ok {
		params.Email = email.(string)
	}
	if sex, ok := jsondata["sex"]; ok {
		params.Sex = sex.(string)
	}
	if status, ok := jsondata["status"]; ok {
		params.Status = int32(status.(float64))
	}
	info, err := query.User.WithContext(c).Where(query.User.ID.Eq(params.ID)).Updates(&params)
	if err != nil {
		panic(err)
	}
	fmt.Println("更新条数", info.RowsAffected)
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "域信息更新成功",
		"data": info.RowsAffected,
	})
}

func DeleteUser(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)

	var userid int64
	if id, ok := jsondata["id"]; ok {
		userid = int64(id.(float64))
	}
	info, err := query.User.WithContext(c).Where(query.Domain.ID.Eq(userid)).Delete()
	if err != nil {
		return
	}
	fmt.Println("删除条数", info.RowsAffected)
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "用户删除成功",
		"data": info,
	})
}

func GetUserDetail(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)
	q := query.User.WithContext(c)
	for key, value := range jsondata {
		switch key {
			case "id":
				q.Where(query.User.ID.Eq(int64(value.(float64))))
			case "username":
				q.Where(query.User.Username.Eq(value.(string)))
			case "nickname":
				q.Where(query.User.Nickname.Eq(value.(string)))
			case "sex":
				q.Where(query.User.Sex.Eq(value.(string)))
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
	find, err := q.Select(query.User.ID, query.User.Nickname, query.User.Username, query.User.Sex, query.User.Email, query.User.Status).Find()
	if err != nil {
		return
	}
	count, err := q.Count()
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  utils.Success,
		"msg":   "用户列表查询成功",
		"data":  find,
		"count": count,
	})
}
