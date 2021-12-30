package service

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/luenci/grpc-demo/config"
	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
)

type StreamServer struct {
	pb.UnimplementedStreamServerServer
}

func (s *StreamServer) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// StreamServiceRun 启动流式服务.
func StreamServiceRun() error {
	// 创建grpc服务器
	server := grpc.NewServer()
	// 注册服务
	pb.RegisterStreamServerServer(server, &StreamServer{})
	// 监听地址
	listener, err := net.Listen(config.Network, config.StreamAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", config.StreamAddress, err)
	}
	log.Println("Listening on", config.StreamAddress)
	// 启动服务
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to stream serve gRPC server: %w", err)
	}
	return nil
}
