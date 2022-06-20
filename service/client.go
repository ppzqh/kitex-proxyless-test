package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless/greetservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/xds"
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
			fmt.Println(err)

		} else {
			fmt.Println(resp.String())
		}
		time.Sleep(time.Second)
	}
	return nil
}
func NewProxylessClient(target string) TestService {
	xds.Init()
	cli, err := greetservice.NewClient(target, client.WithXDSSuite())
	if err != nil {
		panic("construct client failed")
	}
	return &ProxylessClient{cli: cli}
}
