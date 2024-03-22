package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "xxb-grpc-study/hello-client/proto"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "newbie",
		"appKey": "123123",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("D:\\Go_Dev\\src\\xxb-grpc-study\\key\\test.pem", "*.newbie.com")

	//WithTransportCredentials返回一个DialOption，用于配置连接级别的安全凭证(例如，TLS/SSL)。这不能与WithCredentialsBundle一起使用。
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("Did not connect:%v", err)
	}
	defer conn.Close()
	//建立连接
	client := pb.NewSayHelloClient(conn)
	//rpc方法调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: " newbie"})
	fmt.Println(resp.GetResponseMsg())
}
