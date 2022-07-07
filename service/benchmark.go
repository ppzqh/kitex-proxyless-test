package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/codec/thrift/kitex_gen/proxyless/greetservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/xds"
	"sync"
	"sync/atomic"
	"time"
)

type BenchmarkRunner struct {
	targetService string
}

func NewBenchmarkRunner(target string) *BenchmarkRunner {
	return &BenchmarkRunner{targetService: target}
}

func (r *BenchmarkRunner) normalBenchmark() {
	normalCli, err := greetservice.NewClient(r.targetService)
	if err != nil {
		if err != nil {
			klog.Error("[proxy] construct client error: %v\n", err)
			return
		}
	}
	klog.Info("Start normal benchmark")
	benchmark(normalCli)
}

func (r *BenchmarkRunner) proxylessBenchmark() {
	err := xds.Init()
	if err != nil {
		klog.Error("[proxyless] xds init error: %v\n", err)
		return
	}
	cli, err := greetservice.NewClient(r.targetService,
		client.WithXDSSuite(),
	)
	if err != nil {
		klog.Error("[proxyless] construct client error: %v\n", err)
		return
	}
	klog.Info("Start proxyless benchmark")
	benchmark(cli)
}

func (r *BenchmarkRunner) Run() error {
	//// normal client
	//r.normalBenchmark()

	// proxyless client
	r.proxylessBenchmark()
	return nil
}

func benchmark(cli greetservice.Client) {
	var (
		totalReq    uint64 = 10000000
		concurrency        = 10
	)

	var currCnt uint64
	var errCnt uint64
	var wg sync.WaitGroup
	startTime := time.Now()
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for {
				req := &proxyless.HelloRequest{Message: "Hello!"}
				_, err := cli.SayHello(context.Background(), req)
				if err != nil {
					atomic.AddUint64(&errCnt, 1)
				}
				atomic.AddUint64(&currCnt, 1)
				if currCnt >= totalReq {
					wg.Done()
					return
				}
			}
		}()
	}
	wg.Wait()
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)
	qps := float64(totalReq) / totalTime.Seconds()
	fmt.Println("Benchmark Finished")
	fmt.Printf("Total request: %d, error num: %d, cost: %dms, QPS: %f\n", totalReq, errCnt, endTime.Sub(startTime).Milliseconds(), qps)
}
