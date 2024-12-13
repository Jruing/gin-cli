package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DetailResponseStruct struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

// 定义一个处理函数，返回自定义格式的数据
func DetailResponse(c *gin.Context) {
	// 数据
	data := map[string]string{
		"key": "value",
	}

	// 返回自定义格式的响应
	c.JSON(http.StatusOK, DetailResponseStruct{
		Code:  http.StatusOK,
		Msg:   "成功",
		Data:  data,
		Count: 20,
	})
}
