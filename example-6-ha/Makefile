.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/config.proto

.PHONY: run-servers
run-servers:
	PORT=50051 go run cmd/server/main.go &
	PORT=50052 go run cmd/server/main.go &
	PORT=50053 go run cmd/server/main.go

.PHONY: run-client
run-custom-client:
	go run cmd/custom-load-balancer-client/main.go

.PHONY: run-retry-client
run-retry-client:
	go run cmd/retry-client/main.go

.PHONY: run-timeout-client
run-timeout-client:
	go run cmd/timeout-client/main.go

.PHONY: run-roundrobin-client
run-roundrobin-client:
	go run cmd/roundrobin-client/main.go