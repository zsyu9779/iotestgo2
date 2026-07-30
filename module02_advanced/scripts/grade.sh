#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

echo "[1/5] 检查 Module 02 Go 格式"
formatted="$(find module02_advanced -type f -name '*.go' -print0 | xargs -0 gofmt -l)"
if [[ -n "$formatted" ]]; then
	printf '%s\n' "$formatted"
	exit 1
fi

echo "[2/5] 运行 go vet"
go vet ./module02_advanced/...

echo "[3/5] 运行普通测试"
go test ./module02_advanced/...

echo "[4/5] 运行 race detector"
go test -race ./module02_advanced/...

echo "[5/5] 构建 Module 02"
go build ./module02_advanced/...

echo "module02 grade: PASS"
