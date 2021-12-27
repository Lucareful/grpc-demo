package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	simStore := pb.NewSimpleClient(conn)
	if _, err := simStore.GetSimpleInfo(
		context.Background(), &pb.SimpleRequest{
			Data: "luenci",
		}); err != nil {
		return fmt.Errorf("failed to PutSimpleInfo: %w", err)
	}

	log.Println("Successfully PutSimpleInfo")
	return nil
}
