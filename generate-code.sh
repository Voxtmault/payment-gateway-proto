#!/bin/bash

echo "Generating gRPC Codes"

# Ensure Go binaries are in PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go Codes
protoc --go_out=./ --go-grpc_out=./ snap.proto
protoc --go_out=./ --go-grpc_out=./ payment-gateway.proto
echo "Go Codes Generated"

# Generate JavaScript code
protoc --js_out=import_style=commonjs,binary:./js/snap-service snap.proto
protoc --js_out=import_style=commonjs,binary:./js/payment-service payment-gateway.proto
echo "JS Code Generated"

# Generate gRPC-Web code
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/snap-service snap.proto
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/payment-service payment-gateway.proto
echo "JS Web Generated"