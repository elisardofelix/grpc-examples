package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elisardofelix/grpc-examples/example-6-ha/internal/resolve"
	"github.com/elisardofelix/grpc-examples/example-6-ha/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	ctx := context.Background()

	builder := &resolve.Builder{}
	resolver.Register(builder)

	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

	// chris://
	conn, err := grpc.NewClient(fmt.Sprintf("%s://", builder.Scheme()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	for i := range 12 {
		log.Printf("making request %d", i)

		res, err := client.GetServerAddress(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("response received: %s", res.GetAddress())
		time.Sleep(time.Second)
	}
}
