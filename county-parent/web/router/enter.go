package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jenson/county/core/e"
	"github.com/jenson/county/web/api"
)

// InitRouters example Initializing the route mapping
func InitRouters() *gin.Engine {
	engine := gin.Default()

	instance := api.ApiGroupInstance

	engine.GET("/ping", func(c *gin.Context) {
		e.OKWithMessage("pong", c)
	})

	userApi := instance.UserApi
	user := engine.Group("/user")
	{
		user.GET("/greet", userApi.Greet)
		user.POST("/login", userApi.Login)
	}

	countyApi := instance.CountyApi
	county := engine.Group("/county")
	{
		county.GET("/query", countyApi.QueryDetailByEditor)
	}
	return engine
}
