package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type StatusCode int

var (
	Success StatusCode = 200
	Error   StatusCode = 400
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func OK(msg string, data any, c *gin.Context) {
	Result(int(Success), data, msg, c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(int(Success), nil, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(int(Success), data, "操作成功", c)
}

func Fail(c *gin.Context) {
	Result(int(Error), nil, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(int(Error), nil, message, c)
}

func FailWithData(message string, data any, c *gin.Context) {
	Result(int(Error), data, message, c)
}
