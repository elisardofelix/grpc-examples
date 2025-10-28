proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto

tidy:
	go mod tidy

run-server:
	go run cmd/server/main.go

run-server-stream-client:
	go run cmd/server-stream-client/main.go