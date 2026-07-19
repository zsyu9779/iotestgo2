#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "$script_dir/../../.." && pwd)"
cd "$repo_root"

assert_contains() {
    local output="$1"
    local expected="$2"
    if ! grep -Fq "$expected" <<<"$output"; then
        printf 'missing expected output: %s\n' "$expected" >&2
        exit 1
    fi
}

vars_output="$(go run ./module01_basics/blocks/01_go_basics/demo/02_vars_types)"
assert_contains "$vars_output" 'blank identifier: 1 3'
assert_contains "$vars_output" 'const expression length: 3'
assert_contains "$vars_output" 'iota edge cases: 0 1 2 250 250 5 6'
assert_contains "$vars_output" 'iota reset: 0 1'
assert_contains "$vars_output" 'defined type conversion: 12 alias assignment: 12'
assert_contains "$vars_output" 'truncated float: 1'

printf 'module01 demo contracts: PASS\n'
