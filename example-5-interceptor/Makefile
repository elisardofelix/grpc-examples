proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto

tidy:
	go mod tidy

run-server:
	JWT_SECRET=super-secret-key-123 go run cmd/server/main.go

run-d-client:
	go run cmd/deadline-client/main.go

run-auth-client-works:
	JWT_SECRET=super-secret-key-123 go run cmd/auth-client/main.go

run-auth-client-fails:
	JWT_SECRET=wrong-secret-key-456 go run cmd/auth-client/main.go

run-meta-client:
	go run cmd/meta-data-client/main.go

run-validate-client:
	go run cmd/validate-client/main.go

# Commands to run the interceptor example
run-interceptor-server:
	go run cmd/interceptor-server/main.go

run-interceptor-client:
	go run cmd/interceptor-client/main.go