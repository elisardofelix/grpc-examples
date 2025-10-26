# Example 2 - TODO — gRPC example (Go)

A small gRPC Todo example in Go demonstrating a server, client, and an internal service package.

Contents
--------

- `cmd/server/main.go` — gRPC server entrypoint.
- `cmd/client/main.go` — a simple client that calls the server.
- `internal/todo/service.go` — example service implementation.
- `proto/` — protobuf definition and generated Go bindings (`todo.proto`, `todo.pb.go`, `todo_grpc.pb.go`) and a `generate.go` helper.
- `Makefile` — convenience targets for generating protos, running the server and client, and tidying modules.

Prerequisites
-------------

- Go (recommended >= 1.20).
- Optional: `protoc` and the Go protobuf plugins if you want to regenerate proto files yourself.

This repository includes a convenience installer script at the repo root that installs `protoc`, the Go plugins and `grpcurl` on Ubuntu:

```bash
# from repository root
./install_dependencies.sh
```

The script may require `sudo` and will add `$HOME/go/bin` to `~/.bashrc` if needed — open a new shell or run `source ~/.bashrc` after it finishes to update your PATH.

Quick start
-----------

From the `example-2-todo` directory you can use the provided Makefile targets:

```bash
# generate proto bindings
make proto-gen

# tidy modules
make tidy

# run the gRPC server
make run-server

# in another terminal, run the client
make run-client
```

You can also run the server and client directly with `go run`:

```bash
# run server
go run ./example-2-todo/cmd/server/main.go

# run client
go run ./example-2-todo/cmd/client/main.go
```

Regenerating protobuf bindings
-----------------------------

This project includes generated protobuf files in `proto/` and a `generate.go` helper. If you edit `proto/todo.proto` you can regenerate Go bindings by running the `proto-gen` make target which calls `go generate`:

```bash
cd example-2-todo
make proto-gen
```

Alternatively, use `protoc` directly (example run from `example-2-todo`):

```bash
protoc --go_out=. --go-grpc_out=. proto/todo.proto
```

Ensure `protoc-gen-go` and `protoc-gen-go-grpc` are installed and available on your PATH.

Troubleshooting
---------------

- "cannot find package": run `make tidy` (or `go mod tidy`) in `example-2-todo`.
- If the client cannot connect, make sure the server is running and listening on the configured address in `cmd/server/main.go`.

Notes
-----

- Generated `.pb.go` files are included so you don't strictly need `protoc` to run the example.
- If you want, I can add example requests/responses and a sample `grpcurl` command demonstrating expected output.

Enjoy exploring this small gRPC Todo example in Go — tell me if you want a README addition with example `grpcurl` output or runnable snippets for testing.