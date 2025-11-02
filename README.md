# grpc-examples with golang

A small collection of minimal gRPC examples written in Go. Each example is self-contained in its own folder and includes generated protobuf files so you can run them quickly.

## Example projects

- [example-1](example-1/README.md) — Minimal gRPC "Hello" example in Go (quick start, proto included).
- [example-2](example-2/README.md) — Structured example with server, clients, internal package and Makefile targets.
- [example-2-todo](example-2-todo/README.md) — Small Todo gRPC example demonstrating proto generation via `go generate`, server and client.
- [example-3](example-3/README.md) — Streaming RPC patterns (server, server-stream, client-stream, bidi-stream clients).
- [example-3-streamfile](example-3-streamfile/README.md) — File streaming example: client/server that stream files (client serves an HTTP check on port 8080).
- [example-4-streamfile-tls-mtls](example-4-streamfile-tls-mtls/README.md) — File streaming with TLS and mTLS, certificate generation helpers and examples.

Repository tree (top-level view):

```
./
├── example-1/               # minimal gRPC Hello (Go)
│   ├── main.go
│   └── proto/
├── example-2/               # structured example (server, client, internal, proto)
│   ├── cmd/
│   │   ├── server/
│   │   ├── client/
│   │   └── client-error/
│   ├── internal/
│   │   └── hello/
│   └── proto/
├── example-2-todo/          # todo example (proto, server, client, internal)
│   ├── cmd/
│   │   ├── server/
│   │   └── client/
│   ├── internal/
│   │   └── todo/
│   └── proto/
├── example-3/               # streaming patterns (server + various streaming clients)
│   ├── cmd/
│   │   ├── server/
│   │   ├── server-stream-client/
│   │   ├── client-stream-client/
│   │   └── bi-directional-stream-client/
│   ├── internal/
│   │   └── streaming/
│   └── proto/
├── example-3-streamfile/    # file streaming example (client serves HTTP check on :8080)
│   ├── cmd/
│   │   ├── server/
│   │   └── client/
│   ├── internal/
│   │   └── stream/
│   └── proto/
├── example-4-streamfile-tls-mtls/ # file streaming with TLS and mTLS (certs + examples)
│   ├── cmd/
│   │   ├── server/
│   │   ├── client/
│   │   ├── mtls-server/
│   │   └── mtls-client/
│   ├── internal/
│   │   └── stream/
│   ├── certs/
│   └── proto/
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
