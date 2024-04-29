package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	_SUCCESS = 0
	_FAIl    = 7
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpResult(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: _FAIl,
		Data: nil,
		Msg:  message,
	})
}

// 对httpresult方法进行封装，将code隐藏

func HTTPOk(data interface{}, msg string, c *gin.Context) {
	HttpResult(_SUCCESS, data, msg, c)
}

func HTTPFail(data interface{}, msg string, c *gin.Context) {
	HttpResult(_FAIl, data, msg, c)
}
