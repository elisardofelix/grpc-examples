# Example 3 — StreamFile gRPC example (Go)

This example demonstrates streaming file transfer with gRPC in Go. It includes a server and a client that stream file data using protobuf-defined streaming RPCs.

Repository layout
-----------------

- `cmd/server/main.go` — server entrypoint that exposes streaming RPC(s) to send/receive file data.
- `cmd/client/main.go` — client that demonstrates sending or receiving file streams to/from the server.
- `internal/stream/service.go` — server-side service implementation for streaming file transfer logic.
- `proto/` — protobuf definitions and generated Go bindings (`streaming.proto`, `streaming.pb.go`, `streaming_grpc.pb.go`).
- `Makefile` — convenience targets to generate protobuf bindings, tidy modules, and run the server/client.

Prerequisites
-------------

- Go (recommended >= 1.20).
- `protoc` and the Go protobuf plugins if you want to regenerate proto files:
  - protoc (https://grpc.io/docs/protoc-installation/)
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

You can also use the repository convenience installer from the repo root to install `protoc` and the Go plugins on Ubuntu:

```bash
# from the repository root
./install_dependencies.sh
```

Makefile targets / Quick start
-----------------------------

From the `example-3-streamfile` directory you can use the provided Makefile targets:

```bash
# generate protobuf Go bindings (runs protoc on ./proto/*.proto)
make proto-gen

# tidy go modules
make tidy

# run the gRPC server
make run-server

# run the client (talks to the server)
make run-client
```

You can also run the server/client directly with `go run` from the repository root:

```bash
go run ./example-3-streamfile/cmd/server/main.go
go run ./example-3-streamfile/cmd/client/main.go
```

Regenerating protobuf bindings
-----------------------------

If you edit `proto/streaming.proto` you can regenerate the Go bindings using the Makefile target which runs `protoc` with source-relative paths:

```bash
cd example-3-streamfile
make proto-gen
```

Or run `protoc` directly (example):

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
```

Ensure `protoc-gen-go` and `protoc-gen-go-grpc` are installed and on your PATH.

Quick usage notes
-----------------

- Check the configured server address and port in `cmd/server/main.go` and `cmd/client/main.go` if they need to match your environment.
- The example includes generated `.pb.go` files inside `proto/`, so `protoc` isn't strictly required to run the example unless you change the proto.

HTTP check (client)
-------------------

The client exposes a small HTTP server on port `8080` so you can easily verify the streamed file in your browser after a transfer completes. To use it:

1. Start the gRPC server (from `example-3-streamfile`):

```bash
make run-server
```

2. In another terminal, run the client which performs the streaming transfer and serves the received file over HTTP:

```bash
make run-client
```

3. Open your browser and visit:

```
http://localhost:8080/
```

The client will serve the last-received file (or a small web page linking to it). If you change the client or server ports, update the URL accordingly.

Troubleshooting
---------------

- "cannot find package": run `make tidy` (or `go mod tidy`) in `example-3-streamfile`.
- Client connection issues: ensure the server is running and both server and client use the same host:port.
- Protobuf generation errors: verify `protoc` and the Go plugins are installed and available on PATH.