package main

import (
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/service"
	"os"
)

const (
	serviceNameKey = "MY_SERVICE_NAME"
	testClient = "kitex-client"
	testServer = "kitex-server"
	benchmarkClient = "proxyless-benchmark-client"
	suffix = ".default.svc.cluster.local:8888"
)

func main() {
	serviceName, ok := os.LookupEnv(serviceNameKey)
	if !ok {
		panic("Please specify the service name")
	}
	var svc service.TestService
	switch serviceName {
	case testClient:
		svc = service.NewProxylessClient(testClient+suffix, testServer+suffix)
	case testServer:
		svc = service.NewProxylessServer()
	case benchmarkClient:
		svc = service.NewBenchmarkRunner(testServer+suffix)
	}
	fmt.Println("TEST SERVICE START")
	err := svc.Run()
	if err != nil {
		errMsg := fmt.Errorf("test failed with error: %s", err.Error())
		fmt.Println(errMsg)
	}
}
