#!/bin/bash

echo "Generating gRPC Codes"

# Ensure Go binaries are in PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go Codes
protoc --go_out=./ --go-grpc_out=./ snap.proto
protoc --go_out=./ --go-grpc_out=./ payment-gateway.proto
protoc --go_out=./ --go-grpc_out=./ management.proto
echo "Go Codes Generated"