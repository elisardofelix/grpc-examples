package main

import (
	"log"
	"net"

	"github.com/elisardofelix/grpc-examples/example-2/internal/hello"
	"github.com/elisardofelix/grpc-examples/example-2/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, helloService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
