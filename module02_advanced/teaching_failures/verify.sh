#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

run_case() {
	local name="$1"
	local expected="$2"
	shift 2
	local output_file
	output_file="$(mktemp)"

	set +e
	"$@" >"$output_file" 2>&1
	local status=$?
	set -e

	if [[ $status -eq 0 ]]; then
		echo "$name: expected non-zero exit" >&2
		cat "$output_file" >&2
		rm -f "$output_file"
		exit 1
	fi
	if ! grep -Eq "$expected" "$output_file"; then
		echo "$name: diagnostic did not match $expected" >&2
		cat "$output_file" >&2
		rm -f "$output_file"
		exit 1
	fi
	rm -f "$output_file"
	echo "$name: PASS"
}

base="./module02_advanced/teaching_failures/testdata"
run_case send_closed 'send on closed channel' go run "$base/send_closed"
run_case double_close 'close of closed channel' go run "$base/double_close"
run_case deadlock 'all goroutines are asleep|deadlock' go run "$base/deadlock"
run_case goroutine_panic 'panic: child goroutine failed' go run "$base/goroutine_panic"
run_case data_race 'DATA RACE' go run -race "$base/data_race"
run_case concurrent_map_write 'concurrent map writes' go run "$base/concurrent_map_write"

echo "module02 teaching failures: PASS"
