package main

import (
	"context"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}

// SayHi implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHi(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}
