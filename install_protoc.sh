#!/bin/bash
set -e

# =========================================
# Install protoc (Protocol Buffers Compiler)
# and Go plugins on Ubuntu
# =========================================

# -----------------------------------------
# Install protoc
# -----------------------------------------
echo "Installing protoc"
sudo apt install -y protobuf-compiler

# -----------------------------------------
# Install Go protobuf plugins
# -----------------------------------------
echo "Installing Go protobuf plugins..."
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest