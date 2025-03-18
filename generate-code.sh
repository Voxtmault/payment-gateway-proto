#!/bin/bash

echo "Generating gRPC Codes"

# Ensure Go binaries are in PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go Codes
protoc --go_out=./go --go-grpc_out=./go snap.proto
protoc --go_out=./go --go-grpc_out=./go payment-gateway.proto
protoc --go_out=./go --go-grpc_out=./go management.proto
echo "Go Codes Generated"

# Generate JavaScript code
protoc --js_out=import_style=commonjs,binary:./js/snap-service snap.proto
protoc --js_out=import_style=commonjs,binary:./js/payment-service payment-gateway.proto
protoc --js_out=import_style=commonjs,binary:./js/management-service management.proto
echo "JS Code Generated"

# Generate gRPC-Web code
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/snap-service snap.proto
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/payment-service payment-gateway.proto
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/management-service management.proto
echo "JS Web Generated"