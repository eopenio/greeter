package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	proto "github.com/eopenio/greeter/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
    micro.Name("greeter"),
    micro.Version("v0.0.1"),
	)

	// 初始化服务，并且解析命令行传参
	service.Init()

	// RegisterGreeterHandler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

