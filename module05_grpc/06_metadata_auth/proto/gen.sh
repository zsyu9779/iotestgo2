#!/bin/bash
cd "$(dirname "$0")"
mkdir -p authpb
protoc --proto_path=. \
  --go_out=authpb --go_opt=module=iotestgo/module05_grpc/06_metadata_auth/proto/authpb \
  --go-grpc_out=authpb --go-grpc_opt=module=iotestgo/module05_grpc/06_metadata_auth/proto/authpb \
  hello.proto
echo "Generated: authpb/hello.pb.go, authpb/hello_grpc.pb.go"
