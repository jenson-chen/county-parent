package etc

import (
	"fmt"
	"github.com/jenson/county/core/constant"
	"github.com/jenson/county/user/api"
	"github.com/jenson/county/user/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

// RegisterAndStartRemoteServer Register and start a remote service
func RegisterAndStartRemoteServer() {
	addr := fmt.Sprintf(":%d", constant.UserServicePort)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen %v\n", err)
	}

	//创建一个gRPC服务器实例
	srv := grpc.NewServer()
	s := api.UserService{}
	//将service结构体注册为gRPC服务
	service.RegisterUserRpcServer(srv, &s)
	log.Printf("user grpc running %s\n", addr)
	err = srv.Serve(listen)
}
