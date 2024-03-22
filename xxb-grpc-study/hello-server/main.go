package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	pb "xxb-grpc-study/hello-server/proto"
)

//hello server 服务器端
type server struct {
	pb.UnimplementedSayHelloServer
}

//方法重写
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	//获取元数据的信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appId"]; ok {
		appId = v[0]
	}
	if v, ok := md["appKey"]; ok {
		appKey = v[0]
	}

	if appId != "newbie" || appKey != "123123" {
		return nil, errors.New("Token不正确")
	}
	fmt.Printf("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//从提供的根证书颁发机构证书文件构建TLS凭据，两个参数分别是自签名证书和私钥文件
	//creds, _ := credentials.NewServerTLSFromFile("D:\\Go_Dev\\src\\xxb-grpc-study\\key\\test.pem", "D:\\Go_Dev\\src\\xxb-grpc-study\\key\\test.key")
	////开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//在grpc服务端中去注册我们自己编写的服务,注册一定是引用注册
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		log.Println("启动失败~")
	}
}
