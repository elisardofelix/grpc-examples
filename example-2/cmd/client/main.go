package main

import (
	"context"
	"log"

	"github.com/elisardofelix/grpc-examples/example-2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	resp, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %v", resp)
}
