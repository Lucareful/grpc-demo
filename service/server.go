package service

import (
	"fmt"
	"log"
	"net"

	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	log.Println("Listening on", listenOn)
	pb.RegisterSimpleServer(server, &pb.UnimplementedSimpleServer{})
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
