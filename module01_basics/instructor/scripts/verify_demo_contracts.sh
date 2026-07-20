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

slice_output="$(go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices)"
assert_contains "$slice_output" 'copy count=3 dst=[10 20 30 0 0] src=[10 20 30]'

map_output="$(go run ./module01_basics/blocks/02_collections/demo/05_maps_strings)"
assert_contains "$map_output" 'nested map value: ready'

pointer_output="$(go run ./module01_basics/blocks/03_modeling/demo/06_pointers)"
assert_contains "$pointer_output" 'pointer field sugar: 3'

receiver_output="$(go run ./module01_basics/blocks/03_modeling/demo/09_copy_and_receivers)"
assert_contains "$receiver_output" 'value receiver mutation keeps: Alice'

printf 'module01 demo contracts: PASS\n'
