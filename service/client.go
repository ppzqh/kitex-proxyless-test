package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless/greetservice"
	"time"
)

type ProxylessClient struct {
	cli greetservice.Client
}

func (c *ProxylessClient) Run() error {
	for {
		ctx := context.Background()
		req := &proxyless.HelloRequest{Message: "Hello!"}
		resp, err := c.cli.SayHello(ctx, req)
		if err != nil {
			return err
		}
		fmt.Println(resp.String())
		time.Sleep(time.Second)
	}
	return nil
}
func NewProxylessClient(target string) TestService {
	client, err := greetservice.NewClient(target)
	if err != nil {
		panic("construct client failed")
	}

	return &ProxylessClient{cli: client}
}
