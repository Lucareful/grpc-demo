package example

import (
	"context"
	"io"
	"log"

	"github.com/luenci/grpc-demo/config"
	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
)

// StreamClientRun 调用服务端的ListValue方法
func StreamClientRun() error {
	// 连接服务器
	conn, err := grpc.Dial(config.StreamAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Stream net.Connect err: %v", err)
		return err
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewStreamServerClient(conn)

	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc ",
	}
	// 调用我们的服务(ListValue方法)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
		return err
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
			return err
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
	return nil
}
