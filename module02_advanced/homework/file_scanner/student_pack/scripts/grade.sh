#!/usr/bin/env bash
set -euo pipefail

work_dir="$(pwd)"
if [[ ! -f "$work_dir/go.mod" ]]; then
	echo "请从作业根目录运行评分脚本" >&2
	exit 1
fi

echo "[1/4] gofmt"
formatted="$(find . -type f -name '*.go' -print0 | xargs -0 gofmt -l)"
if [[ -n "$formatted" ]]; then
	printf '%s\n' "$formatted"
	exit 1
fi

echo "[2/4] go vet"
go vet ./...

echo "[3/4] go test -race"
go test -race ./...

echo "[4/4] go build"
go build ./...

echo "module02 homework grade: PASS"
