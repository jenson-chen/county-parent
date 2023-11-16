package etc

import (
	"fmt"
	"github.com/jenson/county/core/constant"
	"github.com/jenson/county/county/api"
	"github.com/jenson/county/county/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterAndStartRemoteCountyServer() {
	addr := fmt.Sprintf(":%d", constant.CountyServicePort)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen %v\n", err)
	}

	//创建一个gRPC服务器实例
	srv := grpc.NewServer()
	cs := api.CountyService{}
	//将service结构体注册为gRPC服务
	service.RegisterCountyRpcServer(srv, &cs)
	log.Printf("user grpc running %s\n", addr)
	err = srv.Serve(listen)
}
