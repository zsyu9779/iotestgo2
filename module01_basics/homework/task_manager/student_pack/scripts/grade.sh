#!/usr/bin/env sh
set -eu

echo "[1/4] 检查 gofmt"
unformatted=$(find . -type f -name '*.go' -exec gofmt -l {} \;)
if [ -n "$unformatted" ]; then
  echo "以下文件未通过 gofmt："
  echo "$unformatted"
  exit 1
fi

echo "[2/4] 执行 go vet"
go vet ./...

echo "[3/4] 执行测试"
go test ./...

echo "[4/4] 构建 CLI"
go build ./...

echo "Task Manager 作业验收通过"
