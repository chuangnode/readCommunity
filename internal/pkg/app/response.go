package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readCommunity/internal/pkg/errcode"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"data,omitempty"`
}

func ResponseWithMsg(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseWithError(c *gin.Context, code int, msg string, err interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code":  code,
		"Msg":   msg,
		"Error": err,
	})
}

func ResponseInfo(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  errcode.GetMsg(code),
		Data: data,
	})
}
