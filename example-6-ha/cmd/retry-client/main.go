package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/elisardofelix/grpc-examples/example-6-ha/internal/config"
	"github.com/elisardofelix/grpc-examples/example-6-ha/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	ctx := context.Background()

	cfg := config.Config{
		MethodConfig: []config.MethodConfig{{
			Name: []config.NameConfig{{
				Service: "config.ConfigService",
			}},
			RetryPolicy: &config.RetryPolicy{
				MaxAttempts:          4,
				InitialBackoff:       "1s",
				MaxBackoff:           "10s",
				BackoffMultiplier:    2,
				RetryableStatusCodes: []string{"INTERNAL", "UNAVAILABLE"},
			},
		}},
	}

	serviceConfig, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}

	conn, _ := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(string(serviceConfig)),
	)

	client := proto.NewConfigServiceClient(conn)

	_, err = client.Flaky(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successful response received")
}
