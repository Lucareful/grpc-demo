package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"time"

	"github.com/luenci/grpc-demo/config"

	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type simple struct {
	// 向后兼容
	pb.UnimplementedSimpleServer
}

// GetSimpleInfo 获取信息.
func (s *simple) GetSimpleInfo(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, request, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

// handle 处理请求.
func handle(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := &pb.SimpleResponse{
			Code:  200,
			Value: "Hello: " + req.Data + "world",
		}
		data <- res
	}
}

// SimpleServiceRun Server 启动服务.
func SimpleServiceRun() error {
	listener, err := net.Listen(config.Network, config.SimpleAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", config.SimpleAddress, err)
	}

	server := grpc.NewServer()
	log.Println("Listening on", config.SimpleAddress)
	pb.RegisterSimpleServer(server, &simple{})
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
