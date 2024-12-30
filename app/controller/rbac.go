package controller

import (
	"fmt"
	"gin-cli/app/utils"

	"github.com/gin-gonic/gin"
)

func AddRoleForUser(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	var role string
	if value, ok := jsondata["role"]; ok {
		role = value.(string)
	}
	var user []string
	if value, ok := jsondata["user"]; ok {
		user = value.([]string)
	}
	for k, v := range user {
		fmt.Println(k, v)
		_, err := utils.Enforcer.AddRoleForUser(v, role)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"code": utils.Success,
			"msg":  "添加成功",
		})
	}
}

func DeleteRoleForUser(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	var role string
	if value, ok := jsondata["role"]; ok {
		role = value.(string)
	}
	var user []string
	if value, ok := jsondata["user"]; ok {
		user = value.([]string)
	}
	for k, v := range user {
		fmt.Println(k, v)
		_, err := utils.Enforcer.DeleteRoleForUser(v, role)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"code": utils.Success,
			"msg":  "删除成功",
		})
	}
}

func AddPolicies(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	var role string
	if value, ok := jsondata["role"]; ok {
		role = value.(string)
	}
	var menu []string
	if value, ok := jsondata["menu"]; ok {
		menu = value.([]string)
	}
	var operation []string
	if value, ok := jsondata["operation"]; ok {
		operation = value.([]string)
	}
	var policies [][]string
	for _, menuItem := range menu {
		for _, opItem := range operation {
			policies = append(policies, []string{role, menuItem, opItem})
		}
	}
	_, err := utils.Enforcer.AddPolicies(policies)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "添加成功",
	})
}

func RemovePolicies(c *gin.Context) {
	jsondata := make(map[string]interface{})
	_ = c.Bind(&jsondata)
	var role string
	if value, ok := jsondata["role"]; ok {
		role = value.(string)
	}
	var menu []string
	if value, ok := jsondata["menu"]; ok {
		menu = value.([]string)
	}
	var operation []string
	if value, ok := jsondata["operation"]; ok {
		operation = value.([]string)
	}
	var policies [][]string
	for _, menuItem := range menu {
		for _, opItem := range operation {
			policies = append(policies, []string{role, menuItem, opItem})
		}
	}
	_, err := utils.Enforcer.RemovePolicies(policies)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": utils.Success,
		"msg":  "删除成功",
	})

}
