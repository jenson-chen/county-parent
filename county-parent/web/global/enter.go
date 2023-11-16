package global

import (
	countySrv "github.com/jenson/county/county/service"
	userSrv "github.com/jenson/county/user/service"
)

var (
	// UserServiceClient 用户微服务客户端
	UserServiceClient userSrv.UserRpcClient

	// CountyServiceClient 区县微服务客户端
	CountyServiceClient countySrv.CountyRpcClient
)
