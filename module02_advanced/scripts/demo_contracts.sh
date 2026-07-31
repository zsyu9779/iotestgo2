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

channels_output="$(go run ./module02_advanced/blocks/02_goroutines_channels/demo/04_channels)"
grep -q 'Unbuffered handoff: 42' <<<"$channels_output"
grep -q 'Buffered before receive: len=2 cap=2' <<<"$channels_output"
grep -q 'Closed channel: first=7/true second=0/false' <<<"$channels_output"
grep -q 'Nil channel disabled case: ready' <<<"$channels_output"

pipeline_output="$(go run ./module02_advanced/integrated_lab/log_analyzer/solution)"
grep -q 'Processed 100 logs, found ' <<<"$pipeline_output"

echo "module02 demo contracts: PASS"
