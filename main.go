package main

import (
	"log"

	"github.com/luenci/grpc-demo/service"
	"golang.org/x/sync/errgroup"
)

// main gRPC server 入口.
func main() {

	g := errgroup.Group{}
	g.Go(func() error {
		log.Println("start simple server...")
		return service.SimpleServiceRun()
	})

	g.Go(func() error {
		log.Println("start Stream server...")
		return service.StreamServiceRun()
	})

	if err := g.Wait(); err != nil {
		log.Println("Get errors: ", err)
	}
}
