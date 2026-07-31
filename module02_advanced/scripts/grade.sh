#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

echo "[1/6] 检查 Module 02 Go 格式"
formatted="$(find module02_advanced -type f -name '*.go' -print0 | xargs -0 gofmt -l)"
if [[ -n "$formatted" ]]; then
	printf '%s\n' "$formatted"
	exit 1
fi

echo "[2/6] 运行 go vet"
go vet ./module02_advanced/...

echo "[3/6] 运行普通测试"
go test ./module02_advanced/...

echo "[4/6] 运行 race detector"
go test -race ./module02_advanced/...

echo "[5/6] 构建 Module 02"
go build ./module02_advanced/...

echo "[6/6] 验证教师 Homework Solution"
make module02-homework-solution

echo "module02 grade: PASS"
