package service

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
)

type simple struct {
	// 向后兼容
	pb.UnimplementedSimpleServer
}

func (s *simple) GetSimpleInfo(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	return &pb.SimpleResponse{
		Code:  200,
		Value: "Hello: " + request.Data + "world",
	}, nil
}

// Run Server 启动服务.
func Run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	log.Println("Listening on", listenOn)
	pb.RegisterSimpleServer(server, &simple{})
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
