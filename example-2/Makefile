# Makefile for running the Go gRPC server and grpcurl command
.PHONY: run-serve run-grpcurl

# Command to run the Go gRPC server
run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go

run-client-error:
	go run cmd/client-error/main.go

# Example grpcurl command to call the gRPC server
# Explanation of the flags used:
#grpcurl -d '{"name":"World"}'-> body of the request in JSON format
#		-import-path ./proto -> path to proto files use to marshal/unmarshal
#		-proto hello.proto -> specific proto file use to marshal/unmarshal
#		-plaintext localhost:50051 -> server address
#		hello.HelloService/SayHello -> service and method to call
run-grpcurl:
	grpcurl -d '{"name":"World"}' \
		-import-path ./proto \
		-proto hello.proto \
		-plaintext localhost:50051 \
		hello.HelloService/SayHello