#!/bin/bash
cd "$(dirname "$0")"
mkdir -p hellopb
protoc --proto_path=. \
  --go_out=hellopb --go_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  --go-grpc_out=hellopb --go-grpc_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  hello.proto
echo "Generated: hellopb/hello.pb.go, hellopb/hello_grpc.pb.go"
