package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	res, err := simStore.GetSimpleInfo(
		ctx, &pb.SimpleRequest{
			Data: "luenci",
		})
	if err != nil {
		// 获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			// 判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}

	// 打印返回值
	log.Println(res.Value)

	log.Printf("Successfully PutSimpleInfo: %s", res)
	return nil
}
