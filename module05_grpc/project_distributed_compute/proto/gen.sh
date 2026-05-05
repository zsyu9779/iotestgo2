#!/bin/bash
cd "$(dirname "$0")"
mkdir -p computepb
protoc --proto_path=. \
  --go_out=computepb --go_opt=module=iotestgo/module05_grpc/project_distributed_compute/proto/computepb \
  --go-grpc_out=computepb --go-grpc_opt=module=iotestgo/module05_grpc/project_distributed_compute/proto/computepb \
  compute.proto
echo "Generated: computepb/compute.pb.go, computepb/compute_grpc.pb.go"
