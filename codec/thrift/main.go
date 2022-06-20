package main

import (
	proxyless "github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless/greetservice"
	"log"
)

func main() {
	svr := proxyless.NewServer(new(GreetServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
