package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/luenci/grpc-demo/config"

	pb "github.com/luenci/grpc-demo/protos/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SimpleClientRun is a client for the Simple service.
func SimpleClientRun() error {
	//从输入的证书文件中为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("./cert/server.pem", "grpc-demo")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	conf := config.GetConf()
	token := Token{
		AppID:     conf.Token.AppID,
		AppSecret: conf.Token.AppSecret,
	}

	conn, err := grpc.Dial(config.SimpleAddress, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", config.SimpleAddress, err)
	}
	log.Printf("Connected to Simple service: %s", config.SimpleAddress)

	simStore := pb.NewSimpleClient(conn)

	clientDeadline := time.Now().Add(time.Duration(5 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	res, err := simStore.GetSimpleInfo(
		ctx, &pb.SimpleRequest{
			Data: "luenci",
		})
	if err != nil {
		// 获取错误状态
		sta, ok := status.FromError(err)
		if ok {
			// 判断是否为调用超时
			if sta.Code() == codes.DeadlineExceeded {
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
