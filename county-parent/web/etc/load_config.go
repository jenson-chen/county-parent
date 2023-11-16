package etc

import (
	"fmt"
	"github.com/jenson/county/core/constant"
	countySrv "github.com/jenson/county/county/service"
	userSrv "github.com/jenson/county/user/service"
	"github.com/jenson/county/web/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func LoadConfig() {
	loadUserServiceClient()
	loadCountyServiceClient()
}

func loadUserServiceClient() {
	// 以下为gRPC代码 已经完成客户端到服务端的调用
	addr := fmt.Sprintf(":%d", constant.UserServicePort)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start the User-gRPC client")
	}

	//初始化客户端
	client := userSrv.NewUserRpcClient(conn)
	global.UserServiceClient = client
}

func loadCountyServiceClient() {
	// 以下为gRPC代码 已经完成客户端到服务端的调用
	addr := fmt.Sprintf(":%d", constant.CountyServicePort)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start the User-gRPC client")
	}

	//初始化客户端
	client := countySrv.NewCountyRpcClient(conn)
	global.CountyServiceClient = client
}
