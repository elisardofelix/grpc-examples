# Example 1 — gRPC Hello (Go)

A minimal gRPC example written in Go. This example demonstrates a simple "Hello" service using protobuf and a single Go `main.go` program under `example-1/`.

## Repository layout

- `main.go` — entry point for the example server (or client depending on implementation).
- `go.mod` — Go module file for dependencies.
- `proto/` — protobuf definitions and generated files:
  - `hello.proto` — service and message definitions.
  - `hello.pb.go` — generated Go bindings for the proto file.
  - `generate.go` — helper to regenerate proto bindings (optional).

## Prerequisites

- Go (recommended >= 1.20). Install from https://golang.org/dl/.
- Optionally: `protoc` and the `protoc-gen-go`/`protoc-gen-go-grpc` plugins if you want to regenerate the proto files:
  - protoc (https://grpc.io/docs/protoc-installation/)
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

  Alternatively, this repository includes a convenience script to install the common dependencies on Ubuntu:

  ```bash
  # from the repository root
  ./install_dependencies.sh
  ```

  The script installs `protoc`, the Go protobuf plugins, and `grpcurl`. It may require `sudo` and appends your Go bin (`$HOME/go/bin`) to `~/.bashrc` if necessary — open a new shell or run `source ~/.bashrc` after it finishes so your PATH is updated.

## Build & Run

From the `example-1` directory:

Run directly (recommended for quick testing):

```bash
cd example-1
go run main.go
```

Or build a binary then run it:

```bash
cd example-1
go build -o hello-example .
./hello-example
```

If the example includes both server and client modes, check `main.go` for how to switch behavior (flags or different entrypoints).

## Regenerating protobuf bindings (optional)

If you change `proto/hello.proto` you'll need to regenerate Go bindings. Two common ways:

1) Using the helper (if present):

```bash
# from repo root or example-1
cd example-1
go generate ./...
```

2) Using protoc directly (example run from `example-1`):

```bash
cd example-1
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

Make sure `protoc-gen-go` and `protoc-gen-go-grpc` are on your PATH (installed via `go install` as shown above).

## Notes

- The generated `.pb.go` files are included in `proto/` so you don't strictly need `protoc` to run the example.
- If you need to add dependencies, use `go mod tidy` inside `example-1`.

## Troubleshooting

- "cannot find package" errors: run `go mod tidy` in `example-1`.
- Protobuf generation errors: ensure `protoc` and the Go plugins are installed and available on PATH.

Enjoy exploring this minimal gRPC example in Go! If you want, I can also add a short example client or expand the README with example requests and expected responses.