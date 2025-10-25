# grpc-examples with golang

A small collection of minimal gRPC examples written in Go. Each example is self-contained in its own folder and includes generated protobuf files so you can run them quickly.

## Example projects

- [example-1](example-1/README.md) — Minimal gRPC "Hello" example in Go (quick start, proto included).
- example-2 — A slightly more structured example with a server, internal package and generated grpc files.

Repository tree (top-level view):

```
./
├── example-1/      # minimal gRPC Hello (Go)
│   ├── main.go
│   └── proto/
├── example-2/      # structured example (server, internal, proto)
└── install_dependencies.sh
```

See each example's README for build/run instructions. For `example-1` specifically, open `example-1/README.md` or click the link above.

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
