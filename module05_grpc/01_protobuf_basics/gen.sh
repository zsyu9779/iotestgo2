#!/bin/bash
cd "$(dirname "$0")"
protoc --proto_path=. --go_out=examplepb \
  --go_opt=module=iotestgo/module05_grpc/01_protobuf_basics/examplepb \
  example.proto
echo "Generated: examplepb/example.pb.go"
