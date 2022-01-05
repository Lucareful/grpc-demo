package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/luenci/grpc-demo/config"

	"github.com/luenci/grpc-demo/service"
	"golang.org/x/sync/errgroup"
)

// main gRPC server 入口.
func main() {
	// 初始化配置
	config.InitConf()
	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = service.CheckToken(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	g := errgroup.Group{}
	g.Go(func() error {
		log.Println("start simple server...")
		return service.SimpleServiceRun(interceptor)
	})

	g.Go(func() error {
		log.Println("start Stream server...")
		return service.StreamServiceRun()
	})

	if err := g.Wait(); err != nil {
		log.Println("Get errors: ", err)
	}
}
