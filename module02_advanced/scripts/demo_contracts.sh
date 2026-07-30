#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

interfaces_output="$(go run ./module02_advanced/01_interfaces)"
grep -q 'Woof!' <<<"$interfaces_output"
grep -q 'String: I am a string' <<<"$interfaces_output"

errors_output="$(go run ./module02_advanced/02_errors_defer)"
grep -q 'Main continues after recover' <<<"$errors_output"
grep -q 'Deferred 1: Cleanup resources' <<<"$errors_output"

pipeline_output="$(go run ./module02_advanced/project_log_analyzer)"
grep -q 'Processed 100 logs, found ' <<<"$pipeline_output"

echo "module02 demo contracts: PASS"
