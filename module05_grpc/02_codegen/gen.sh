#!/bin/bash
cd "$(dirname "$0")"
protoc --proto_path=. \
  --go_out=hellopb --go_opt=module=iotestgo/module05_grpc/02_codegen/hellopb \
  --go-grpc_out=hellopb --go-grpc_opt=module=iotestgo/module05_grpc/02_codegen/hellopb \
  hello.proto
echo "Generated: hellopb/hello.pb.go, hellopb/hello_grpc.pb.go"
