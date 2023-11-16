package main

import "github.com/jenson/county/user/etc"

func main() {
	etc.LoadConfig()
	etc.RegisterAndStartRemoteServer() //启动服务
}
