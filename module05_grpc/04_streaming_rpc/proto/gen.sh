#!/bin/bash
cd "$(dirname "$0")"
mkdir -p chatpb
protoc --proto_path=. \
  --go_out=chatpb --go_opt=module=iotestgo/module05_grpc/04_streaming_rpc/proto/chatpb \
  --go-grpc_out=chatpb --go-grpc_opt=module=iotestgo/module05_grpc/04_streaming_rpc/proto/chatpb \
  chat.proto
echo "Generated: chatpb/chat.pb.go, chatpb/chat_grpc.pb.go"
