package controller

import (
	"fmt"
	"gin-cli/app/dal/model"
	"gin-cli/app/dal/query"
	"gin-cli/app/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateDomain(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.BindJSON(&jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.Domain{
		Domain:  "",
		Status:  0,
		Created: time.Time{},
	}

	if domainname, ok := jsondata["domain"]; ok {
		params.Domain = domainname.(string)
	}

	if status, ok := jsondata["status"]; ok {
		params.Status = int32(status.(float64))
	}

	params.Created = time.Now()
	err := query.Domain.WithContext(c).Create(&params)

	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "域新增成功",
	})
}

func UpdateDomain(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(jsondata)
	query.SetDefault(utils.MysqlClient)
	params := model.Domain{
		ID:      0,
		Domain:  "",
		Status:  0,
		Created: time.Time{},
	}
	if id, ok := jsondata["id"]; ok {
		params.ID = id.(int64)
	}

	if domainname, ok := jsondata["domain"]; ok {
		params.Domain = domainname.(string)
	}

	if status, ok := jsondata["domain"]; ok {
		params.Status = status.(int32)
	}

	//err := query.Domain.WithContext(c).Where(query.Domain.ID.Eq(params.ID)).Update()
	//if err != nil {
	//	return
	//}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "域信息更新成功",
	})
}

func DeleteDomain(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)

	var domainid int64
	if id, ok := jsondata["id"]; ok {
		domainid = int64(id.(float64))
	}
	info, err := query.Domain.WithContext(c).Where(query.Domain.ID.Eq(domainid)).Delete()
	if err != nil {
		return
	}
	fmt.Println("删除条数", info.RowsAffected)
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "域删除成功",
		"data": info,
	})
}

func GetDomainDetail(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	query.SetDefault(utils.MysqlClient)
	q := query.Domain.WithContext(c).Where()
	var Limit int
	if limit, ok := jsondata["limit"]; ok {
		Limit = int(limit.(float64))
		q.Limit(Limit)
	}
	if page, ok := jsondata["page"]; ok {
		q.Offset((int(page.(float64)) - 1) * Limit)
	}
	find, err := q.Columns(query.Domain.ID.Eq("1"))
	if err != nil {
		return
	}
	count, err := q.Count()
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  utils.Success,
		"msg":   "域名列表查询成功",
		"data":  find,
		"count": count,
	})
}
