# Example 5 — Interceptor examples (Go)

This example demonstrates various gRPC interceptor patterns and related client behaviors in Go:

- deadline handling
- authentication (JWT) middleware
- metadata handling
- request validation
- a dedicated interceptor-server / interceptor-client demo

Repository layout
-----------------

- `cmd/server/` — main server that exercises handlers used by several demo clients (deadline, auth, metadata, validate).
- `cmd/deadline-client/` — client demonstrating deadline/cancellation handling.
- `cmd/auth-client/` — client demonstrating JWT-based authentication (works/fails depending on JWT_SECRET env).
- `cmd/meta-data-client/` — client demonstrating sending metadata to the server, MaxCallRecvMsgSize and MaxCallSendMsgSize implementation, and concept of header (pre request data) and trailer (post request data)
- `cmd/validate-client/` — client demonstrating JWT-based input validation behavior and claims management.
- `cmd/interceptor-server/` — separate server showing interceptor wiring and behavior.
- `cmd/interceptor-client/` — client that talks to the interceptor server using the interceptor on the client itself.
- `proto/` — protobuf definitions and generated Go bindings.
- `Makefile` — convenience targets to generate protos, tidy modules and run the server/clients (including interceptor variants).

Prerequisites
-------------

- Go (recommended >= 1.20).
- Optionally: `protoc` and the Go protobuf plugins if you plan to regenerate proto bindings:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

You can use the repository-level installer to add `protoc` and the Go plugins on Ubuntu:

```bash
# from repository root
./install_dependencies.sh
```

Makefile targets / Quick start
-----------------------------

From the `example-5-interceptor` directory you can use the provided Makefile targets:

```bash
# generate protobuf Go bindings
make proto-gen

# tidy go modules
make tidy

# run the main server used by many demo clients (requires JWT_SECRET env var for auth demo)
JWT_SECRET=super-secret-key-123 make run-server

# run the deadline-demo client
make run-d-client

# run the auth client with a valid secret (succeeds)
JWT_SECRET=super-secret-key-123 make run-auth-client-works

# run the auth client with an invalid secret (fails)
JWT_SECRET=wrong-secret-key-456 make run-auth-client-fails

# run the metadata demo client
make run-meta-client

# run the validation demo client
make run-validate-client

# run the interceptor server and client (separate example)
make run-interceptor-server
make run-interceptor-client
```

Notes about the auth demo
------------------------

- The Makefile sets an example `JWT_SECRET` environment variable when running the server or client in the auth demo cases. The server and client use this value to generate/verify a simple JWT used for authentication in the demo. Use the same secret for both server and client to simulate a valid auth flow.

Troubleshooting
---------------

- If you see "cannot find package" errors: run `make tidy` (or `go mod tidy`) in `example-5-interceptor`.
- If the auth client fails with an authentication error, verify the `JWT_SECRET` environment variable used when running the server and client match.
- For deadline/cancellation demos, increase logging on the server or adjust deadlines in the client to observe different behaviors.

Next steps / extras
-------------------

- I can add short example requests and the expected server responses for each client (helpful for quick verification).
- I can add a small script that runs the server and a client in the background and captures their logs for easier demo playback.

Enjoy exploring interceptor patterns and middleware with these small demos. Let me know if you'd like sample outputs added to the README or to wire a demo script for automated runs.