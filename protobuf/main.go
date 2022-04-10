package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protos "google.golang.org/helloworld/helloworld/currency"
	"google.golang.org/helloworld/helloworld/server"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	// l, err := net.Listen("tcp", fmt.Sprintf(":%d`", 9092))
	l, err := net.Listen("tcp", "localhost:9092")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	fmt.Println("11")
	log.Info("reate success")
	// listen for requests
	gs.Serve(l)
}
