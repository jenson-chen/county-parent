package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jenson/county/core/e"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/core/request"
	"github.com/jenson/county/web/global"
	"time"
)

// Greet method is to welcome the user and return a friendly greeting
func (ua UserApi) Greet(c *gin.Context) {
	client := global.UserServiceClient
	ctx, cancelFunc := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancelFunc()

	var req request.GreetReq
	_ = c.ShouldBindJSON(&req)

	resp, err := client.SayHello(ctx, &models.GreetReq{Username: req.Username})
	if err != nil {
		e.FailWithMessage(err.Error(), c)
		return
	}
	e.OKWithData(resp.GetResult(), c)
}

func (ua UserApi) Login(c *gin.Context) {
	client := global.UserServiceClient
	ctx, cancelFunc := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancelFunc()

	var req models.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		e.FailWithMessage(err.Error(), c)
		return
	}
	resp, err := client.Login(ctx, &req)
	if err != nil {
		e.FailWithMessage(err.Error(), c)
		return
	}
	e.OK("登录成功", resp, c)
}
