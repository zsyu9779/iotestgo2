#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

echo '[1/4] 检查 Module 04 Go 格式'
formatted="$(find module04_gorm -type f -name '*.go' -print0 | xargs -0 gofmt -l)"
if [[ -n "$formatted" ]]; then
	printf '%s\n' "$formatted"
	exit 1
fi

echo '[2/4] 运行 go vet'
go vet ./module04_gorm/...

echo '[3/4] 运行离线测试'
go test ./module04_gorm/...

echo '[4/4] 构建所有正常包'
go build ./module04_gorm/...

echo 'module04 verify: PASS'
