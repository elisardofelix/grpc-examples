# Example 4 — StreamFile with TLS and mTLS (Go)

This example demonstrates streaming file transfer over gRPC with TLS (Server only validation) and mutual TLS (mTLS) (Client and Server validation) in Go. It includes helpers to generate a local CA and server/client certificates, and separate server/client programs for TLS and mTLS modes.

Repository layout
-----------------

- `cmd/server/` — server using TLS (server presents a certificate).
- `cmd/client/` — client that connects to the TLS server (verifies server cert).
- `cmd/mtls-server/` — server configured to require client certificates (mTLS).
- `cmd/mtls-client/` — client configured to present a client certificate for mTLS.
- `internal/stream/` — service implementation for streaming file transfer logic.
- `proto/` — protobuf definition and generated Go bindings (`streaming.proto`, `streaming.pb.go`, `streaming_grpc.pb.go`).
- `certs/` — sample certificate artifacts (CA, server and client cert/key). Check this folder for pre-created files or generate new ones with the Makefile.
- `Makefile` — convenience targets to generate protos, generate certificates, tidy modules and run the server/clients.

Prerequisites
-------------

- Go (recommended >= 1.20).
- `protoc` and the Go protobuf plugins if you want to regenerate proto files (optional):
  - protoc (https://grpc.io/docs/protoc-installation/)
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- `openssl` is required if you want to use the Makefile to generate CA/server/client certificates locally.

You can also use the repository convenience installer from the repo root to install the Go plugins and `grpcurl` (Ubuntu):

```bash
# from the repository root
./install_dependencies.sh
```

Makefile targets / Quick start
-----------------------------

From the `example-4-streamfile-tls-mtls` directory you can run the provided make targets:

```bash
# generate protobuf Go bindings
make proto-gen

# tidy go modules
make tidy

# generate a local CA and server/client certs (requires openssl)
make gen-certs

# run the TLS server (uses certs/server.crt and certs/server.key)
make run-server

# run the TLS client (verifies server cert against certs/ca.crt)
make run-client

# run the mTLS server (requires client certs; server verifies client cert)
make run-mtls-server

# run the mTLS client (presents client cert to server)
make run-mtls-client
```

Notes about certificates and TLS/mTLS
------------------------------------

- The `gen-certs` target (and its sub-targets) generate a CA (`certs/ca.crt`, `certs/ca.key`), a server key/cert pair (`certs/server.key` / `certs/server.crt`), and a client key/cert pair (`certs/client.key` / `certs/client.crt`). Generated server and client certs are signed by the local CA.

- For TLS mode (server + client):
  - The server uses `certs/server.crt`/`certs/server.key`.
  - The client trusts the server by using `certs/ca.crt` as the CA bundle to verify the server certificate.

- For mTLS mode (mutual TLS):
  - The server is configured to require and verify client certificates against `certs/ca.crt`.
  - The mTLS client is configured to present `certs/client.crt` and `certs/client.key` when connecting.

Security note: the provided certificates are for local testing only. Do not use these generated keys/certs in production.

How to run a quick demo
-----------------------

1. Generate certs (if they aren't already present):

```bash
cd example-4-streamfile-tls-mtls
make gen-certs
```

2. Start the TLS server in one terminal:

```bash
make run-server
```

3. Run the TLS client in another terminal (it will verify the server cert):

```bash
make run-client
```

4. To test mTLS, start the mTLS server and run the mTLS client:

```bash
make run-mtls-server
make run-mtls-client
```

5. Open your browser and visit:

```
http://localhost:8080/
```

The client will serve the last-received file (or a small web page linking to it). If you change the client or server ports, update the URL accordingly.

Protobuf regeneration
---------------------

If you edit `proto/streaming.proto` you can regenerate the Go bindings with the `proto-gen` target:

```bash
cd example-4-streamfile-tls-mtls
make proto-gen
```

Or run `protoc` directly as shown in the Makefile.

Troubleshooting
---------------

- "cannot find package": run `make tidy` (or `go mod tidy`) in the module directory.
- Openssl errors when generating certs: ensure `openssl` is installed and available on PATH.
- Certificate verification failures: confirm the client/server are using the expected CA file (`certs/ca.crt`) and that the hostnames/IPs match those in the cert subjectAltName entries; the Makefile uses `localhost` and `127.0.0.1`.