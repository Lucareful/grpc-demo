package main

import (
	"log"

	client "github.com/luenci/grpc-demo/client/example"
	"golang.org/x/sync/errgroup"
)

func main() {

	g := errgroup.Group{}
	g.Go(func() error {
		log.Println("start simple server...")
		return client.SimpleClientRun()
	})

	g.Go(func() error {
		log.Println("start Stream server...")
		return client.StreamClientRun()
	})

	if err := g.Wait(); err != nil {
		log.Println("Get errors: ", err)
	}

}
