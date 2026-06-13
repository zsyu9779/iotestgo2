#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
mkdir -p hellopb

GOOGLEAPIS_DIR="$(find "$(go env GOPATH)/pkg/mod" -path '*/third_party/googleapis/google/api/annotations.proto' -print -quit)"
if [ -z "$GOOGLEAPIS_DIR" ]; then
  echo "google/api/annotations.proto not found in module cache" >&2
  exit 1
fi
GOOGLEAPIS_DIR="$(dirname "$(dirname "$(dirname "$GOOGLEAPIS_DIR")")")"

protoc --proto_path=. \
  --proto_path="$GOOGLEAPIS_DIR" \
  --go_out=hellopb --go_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  --go-grpc_out=hellopb --go-grpc_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  --grpc-gateway_out=hellopb --grpc-gateway_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  hello.proto
echo "Generated: hellopb/hello.pb.go, hellopb/hello_grpc.pb.go, hellopb/hello.gw.pb.go"
