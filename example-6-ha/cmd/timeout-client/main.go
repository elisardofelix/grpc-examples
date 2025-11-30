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

	//	const grpcServiceConfig = `{
	//	"methodConfig": [{
	//		"name": [{"service": "config.ConfigService", "method": "LongRunning"}],
	//		"timeout": "1s"
	//	}]
	//}`
	cfg := config.Config{
		MethodConfig: []config.MethodConfig{{
			Name: []config.NameConfig{{
				Service: "config.ConfigService",
				Method:  "LongRunning",
			}},
			Timeout: "1s", // You can adjust the timeout duration here as needed
		}},
	}

	serviceConfig, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(string(serviceConfig)),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	_, err = client.LongRunning(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
}
