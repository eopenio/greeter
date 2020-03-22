package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	proto "github.com/eopenio/greeter/proto"
)


func main() {
	// 创建新服务
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建新的客户端
	greeter := proto.NewGreeterService("greeter", service.Client())

	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Eter"})
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Greeting)
}

