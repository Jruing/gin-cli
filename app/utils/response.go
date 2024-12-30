package utils

import "github.com/gin-gonic/gin"

type DetailResponseStruct struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func (d DetailResponseStruct) Response(c *gin.Context) {
	d.Code = Success
	d.Msg = "success"
	d.Data = nil
	d.Count = 0
	c.JSON(200, d)
}
