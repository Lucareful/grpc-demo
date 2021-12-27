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
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	simStore := pb.NewSimpleClient(conn)
	res, err := simStore.GetSimpleInfo(
		context.Background(), &pb.SimpleRequest{
			Data: "luenci",
		})
	if err != nil {
		return fmt.Errorf("failed to PutSimpleInfo: %w", err)
	}

	log.Printf("Successfully PutSimpleInfo: %s", res)
	return nil
}
