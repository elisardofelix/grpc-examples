package main

import (
	"context"
	"log"

	"github.com/elisardofelix/grpc-examples/example-2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	resp, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: ""})
	if err != nil {
		if status, ok := status.FromError(err); ok {
			log.Fatalf("gRPC Error - Code: %v, Message: %v", status.Code(), status.Message())
		}

		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Success Response: %v", resp)
}
