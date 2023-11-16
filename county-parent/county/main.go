package main

import "github.com/jenson/county/county/etc"

func main() {
	etc.LoadConfig() //加载本地配置文件
	etc.RegisterAndStartRemoteCountyServer()
}
