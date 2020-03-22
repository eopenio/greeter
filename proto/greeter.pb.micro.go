// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/greeter.proto

package greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}
