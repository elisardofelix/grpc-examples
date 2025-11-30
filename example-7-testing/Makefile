.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/chat.proto

.PHONY: mock-gen
mock-gen:
	go generate ./...

.PHONY: run-server
run-server:
	go run cmd/server-with-deps/main.go

.PHONY: run-client
run-client:
	go run cmd/client/main.go

.PHONY: test-unit
test-unit:
	go test -v ./...

.PHONY: test-integration
test-integration:
	go test --tags integration -v ./...