package router

import (
	"gin-cli/app/service/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateDomain(c *gin.Context) {
	db := domain.Queries{}
	params := domain.CreateDomainParams{
		Domain:  "",
		Status:  0,
		Created: pgtype.Timestamp{},
	}
	_, err := db.CreateDomain(c, params)
	if err != nil {
		return
	}
}

func UpdateDomain(c *gin.Context) {
	db := domain.Queries{}
	params := domain.UpdateDomainParams{
		ID:      0,
		Domain:  "",
		Status:  0,
		Created: pgtype.Timestamp{},
	}
	err := db.UpdateDomain(c, params)
	if err != nil {
		return
	}
}

func DeleteDomain(c *gin.Context) {
	db := domain.Queries{}
	err := db.DeleteDomain(c, 1)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

func GetDomainDetail(c *gin.Context) {
	db := domain.Queries{}
	params := domain.GetDomainDetailParams{
		Limit:  0,
		Offset: 0,
	}
	detail, err := db.GetDomainDetail(c, params)
	if err != nil {
		return
	}
	count_params := domain.GetDomainDetailParams{
		Column1: "",
		Limit:   0,
		Offset:  0,
	}
	count, err := db.GetDomainDetail(c, count_params)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data":  detail,
		"count": count,
	})
}
