# Example 2 — gRPC example (Go)

A slightly more structured gRPC example in Go showing a server, simple clients, and an internal service package.

This example contains:

- `cmd/server/main.go` — the gRPC server entrypoint.
- `cmd/client/main.go` — a client that calls the server (happy path).
- `cmd/client-error/main.go` — a client which demonstrates an error path.
- `internal/hello/service.go` — example service implementation.
- `proto/` — protobuf definitions and generated Go bindings (includes `generate.go`, `hello.pb.go`, and `hello_grpc.pb.go`).
- `Makefile` — convenience targets to run server, clients and an example `grpcurl` call.

Prerequisites
-------------

- Go (recommended >= 1.20).
- Optionally: `protoc` and the Go protobuf plugins if you want to regenerate proto files.

You can quickly install common dependencies (Ubuntu) using the repository convenience script from the repo root:

```bash
./install_dependencies.sh
```

This installs `protoc`, the Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`) and `grpcurl`. It may require `sudo` and will add `$HOME/go/bin` to your `~/.bashrc` if needed — open a new shell or run `source ~/.bashrc` after it finishes.

Quick start
-----------

From the `example-2` directory you can use the provided Makefile targets:

```bash
# run the gRPC server
make run-server

# run the happy-path client (calls the server)
make run-client

# run the client that demonstrates an error
make run-client-error

# run a grpcurl example (server must be listening, this uses localhost:50051)
make run-grpcurl
```

The Makefile's `run-grpcurl` target demonstrates how to call the `HelloService/SayHello` method with `grpcurl` and expects the server to be listening on `localhost:50051`.

Run with `go run` directly
-------------------------

You can also run the binaries directly with `go run` from the repository root or from inside `example-2`:

```bash
go run ./example-2/cmd/server/main.go
go run ./example-2/cmd/client/main.go
```

Regenerating protobuf bindings
-----------------------------

If you change `proto/hello.proto` you can regenerate Go bindings in two ways:

1) Using the helper in `proto` (if present):

```bash
# from example-2
cd example-2
go run ./proto/generate.go
```

2) Using `protoc` directly (example):

```bash
cd example-2
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

Make sure `protoc-gen-go` and `protoc-gen-go-grpc` are installed and on your PATH.

Troubleshooting
---------------

- If you see "cannot find package" errors: run `go mod tidy` in `example-2`.
- If `grpcurl` fails to connect, ensure the server is running and listening on the expected address (Makefile uses `localhost:50051`).

Notes
-----

- The generated `.pb.go` files are included, so you don't strictly need `protoc` to run the example.
- If you want, I can add a README-level example `grpcurl` command with expected output, or add short usage docs for each client command. Let me know which you'd prefer.