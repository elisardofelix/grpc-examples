package hello

import (
	"context"
	"fmt"

	"github.com/elisardofelix/grpc-examples/example-2/proto"
)

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func (s Service) SayHello(ctx context.Context, request *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", request.GetName()),
	}, nil
}
