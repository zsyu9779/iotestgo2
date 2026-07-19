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

control_output="$(go run ./module01_basics/blocks/01_go_basics/demo/03_control_funcs)"
assert_contains "$control_output" 'outer score: 50'
assert_contains "$control_output" 'switch init: owner has full access'
assert_contains "$control_output" 'loop body: 0'
assert_contains "$control_output" 'loop body: 2'
assert_contains "$control_output" 'infinite for stopped at: 3'
assert_contains "$control_output" 'Admin branch reached by fallthrough'

string_output="$(go run ./module01_basics/blocks/01_go_basics/demo/05_strings_basics)"
assert_contains "$string_output" 'bytes round trip: Go语言'
assert_contains "$string_output" 'Atoi success: 42'
assert_contains "$string_output" 'Atoi error: true'

printf 'module01 demo contracts: PASS\n'
