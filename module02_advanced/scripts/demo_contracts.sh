#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

interfaces_output="$(go run ./module02_advanced/blocks/01_interfaces_errors/demo/01_interfaces)"
grep -q 'Woof!' <<<"$interfaces_output"
grep -q 'String: I am a string' <<<"$interfaces_output"
grep -q 'Pointer method set: after' <<<"$interfaces_output"

errors_output="$(go run ./module02_advanced/blocks/01_interfaces_errors/demo/02_errors_defer)"
grep -q 'user "alice": user not found' <<<"$errors_output"
grep -q 'is not found: true' <<<"$errors_output"
grep -q 'invalid field: username' <<<"$errors_output"
grep -q 'recovered at boundary: unexpected state' <<<"$errors_output"

file_io_output="$(go run ./module02_advanced/blocks/01_interfaces_errors/demo/03_file_io)"
grep -q 'Read after seek(7): World' <<<"$file_io_output"
grep -q 'Expanded Scanner accepted token: true' <<<"$file_io_output"

channels_output="$(go run ./module02_advanced/blocks/02_goroutines_channels/demo/04_channels)"
grep -q '^1: 2$' <<<"$channels_output"
grep -q '^10: 29$' <<<"$channels_output"
grep -q '^100: 541$' <<<"$channels_output"

pipeline_output="$(go run ./module02_advanced/integrated_lab/log_analyzer/solution)"
grep -q 'Processed 100 logs, found ' <<<"$pipeline_output"

echo "module02 demo contracts: PASS"
