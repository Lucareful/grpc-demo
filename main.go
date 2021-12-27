package main

import (
	"log"

	"github.com/luenci/grpc-demo/service"
)

// main gRPC server 入口.
func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
