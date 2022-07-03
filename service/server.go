package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless/greetservice"
	"github.com/cloudwego/kitex/server"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	resp = proxyless.NewHelloResponse()
	resp.SetMessage("Hello!")
	fmt.Println("SayHello is called")
	return
}

type ProxylessServer struct {
	svr server.Server
}

func NewProxylessServer() TestService {
	return &ProxylessServer{svr: greetservice.NewServer(&GreetServiceImpl{})}
}

func (s *ProxylessServer) Run() error {
	err := s.svr.Run()
	return err
}
