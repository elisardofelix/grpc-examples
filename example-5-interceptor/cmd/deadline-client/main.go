package main

import (
	"context"
	"log"
	"time"

	"github.com/elisardofelix/grpc-examples/example-5-interceptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	// Create a context with a timeout
	// You can adjust the timeout duration to see how it affects the RPC call.
	// Here is an example that will succeed:
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// will success because the server takes 5 seconds to respond
	// Here is an example implemented that will timeout before the server responds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	_, err = client.LongRunning(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RPC call successfully made")
}
