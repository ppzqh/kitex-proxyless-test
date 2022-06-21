package main

import (
	"flag"
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/service"
)

var (
	serviceName       string
	targetServiceName string
)

const (
	testClient = "kitex-client"
	testServer = "kitex-server"
	suffix = ".default.svc.cluster.local"
)

func initFlags() {
	flag.StringVar(&serviceName, "serviceName", "", "specify the name of your service")
	flag.StringVar(&targetServiceName, "targetServiceName", "", "specify the name of your target service")

	flag.Parse()
}

func main() {
	initFlags()
	var svc service.TestService
	switch serviceName {
	case testClient:
		if targetServiceName == "" {
			panic("please set the target name")
		}
		svc = service.NewProxylessClient(targetServiceName+suffix)
	case testServer:
		svc = service.NewProxylessServer()
	default:
		panic("Tag wrong")
	}
	fmt.Println("TEST SERVICE START")
	err := svc.Run()
	if err != nil {
		errMsg := fmt.Errorf("test failed with error: %s", err.Error())
		fmt.Println(errMsg)
	}
}
