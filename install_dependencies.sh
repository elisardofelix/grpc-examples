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
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  
# -----------------------------------------
# Install mockgen (GoMock)
# -----------------------------------------
echo "Installing mockgen..."    
go install github.com/golang/mock/mockgen@latest

# Add Go bin to PATH if not already
if [[ ":$PATH:" != *":$HOME/go/bin:"* ]]; then
  echo "export PATH=\$PATH:\$HOME/go/bin" >> ~/.bashrc
  export PATH=$PATH:$HOME/go/bin
fi

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest