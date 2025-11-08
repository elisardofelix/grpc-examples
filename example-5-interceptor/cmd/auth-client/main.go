package main

import (
	"context"
	"log"
	"os"

	"github.com/elisardofelix/grpc-examples/example-5-interceptor/internal/auth"
	"github.com/elisardofelix/grpc-examples/example-5-interceptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	ctx := context.Background()

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal("JWT_SECRET is required")
	}

	authService, err := auth.NewService(jwtSecret)
	if err != nil {
		log.Fatalf("failed to initialise auth service: %v", err)
	}

	token, err := authService.IssueToken(ctx, "user-id-1234")
	if err != nil {
		log.Fatalf("failed to issue token: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	_, err = client.Protected(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successful response")
}
