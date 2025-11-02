# Example 3 — Streaming gRPC example (Go)

A collection of streaming gRPC examples written in Go. This example demonstrates server-streaming, client-streaming and bidirectional streaming patterns using a small `streaming` service.

Repository layout
-----------------

- `cmd/server/main.go` — gRPC server (serves streaming RPCs).
- `cmd/server-stream-client/main.go` — client that demonstrates a server-streaming call.
- `cmd/client-stream-client/main.go` — client that demonstrates a client-streaming call.
- `cmd/bi-directional-stream-client/main.go` — client that demonstrates bidirectional streaming.
- `internal/streaming/service.go` — server-side service implementation for streaming RPCs.
- `proto/` — protobuf definition and generated Go bindings (`streaming.proto`, `streaming.pb.go`, `streaming_grpc.pb.go`).
- `Makefile` — convenience targets to generate protos, tidy modules and run server/clients.

Prerequisites
-------------

- Go (recommended >= 1.20).
- Optional: `protoc` and the `protoc-gen-go` / `protoc-gen-go-grpc` plugins if you want to regenerate proto files.

This repository includes a convenience installer script at the repo root to install `protoc`, the Go plugins and `grpcurl` on Ubuntu:

```bash
# from the repository root
./install_dependencies.sh
```

The script may require `sudo` and will add `$HOME/go/bin` to `~/.bashrc` if needed — open a new shell or run `source ~/.bashrc` after it finishes to update your PATH.

Makefile targets / Quick start
-----------------------------

From the `example-3` directory you can use the provided Makefile targets:

```bash
# generate protobuf Go bindings (uses protoc on ./proto/*.proto)
make proto-gen

# tidy go modules
make tidy

# run the gRPC server
make run-server

# run the server-streaming client (server must be running)
make run-server-stream-client

# run the client-streaming client (server must be running)
make run-client-stream-client

# run the bidirectional streaming client (server must be running)
make run-bi-directional-stream-client
```

You can also run the commands directly with `go run` from the repository root:

```bash
go run ./example-3/cmd/server/main.go
go run ./example-3/cmd/server-stream-client/main.go
```

Regenerating protobuf bindings
-----------------------------

If you edit `proto/streaming.proto` you can regenerate the bindings using the Makefile target which runs `protoc` with source-relative paths:

```bash
cd example-3
make proto-gen
```

Or run `protoc` directly (example):

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
```

Ensure `protoc-gen-go` and `protoc-gen-go-grpc` are installed and on your PATH.

Troubleshooting
---------------

- "cannot find package": run `make tidy` (or `go mod tidy`) in `example-3`.
- If a client cannot connect: confirm the server is running and that the server address/port match what's configured in the client code.
- If `protoc` fails: ensure `protoc` and its Go plugins are installed (`./install_dependencies.sh` can help on Ubuntu).

Notes
-----

- Generated `.pb.go` files are included under `proto/` so `protoc` is not strictly required to run the examples.
- If you'd like, I can add short example input/output snippets for each streaming client, or add a tiny `scripts/` folder with helper scripts to run the server + a client in separate terminals.

Enjoy exploring streaming RPC patterns with this example. Let me know if you want sample `grpcurl` commands or automated test scripts to exercise each streaming scenario.