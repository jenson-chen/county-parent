package main

import (
	"fmt"
	"github.com/jenson/county/core/constant"
	"github.com/jenson/county/web/etc"
	"github.com/jenson/county/web/router"
)

// Main start entry
func main() {
	etc.LoadConfig() //加载配置文件

	addr := fmt.Sprintf("%s:%d", constant.WebServiceHost, constant.WebServicePort)

	routers := router.InitRouters()
	if err := routers.Run(addr); err != nil {
		panic("Failed to start the route, and the program ended")
	}

}
