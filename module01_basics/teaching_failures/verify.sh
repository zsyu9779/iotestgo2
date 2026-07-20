#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "$script_dir/../.." && pwd)"
cd "$repo_root"

expect_failure() {
    local command_kind="$1"
    local package_path="$2"
    local expected_pattern="$3"
    local output
    local status

    set +e
    if [[ "$command_kind" == "test" ]]; then
        output="$(go test "$package_path" 2>&1)"
        status=$?
    else
        output="$(go run "$package_path" 2>&1)"
        status=$?
    fi
    set -e

    if [[ $status -eq 0 ]]; then
        printf 'expected failure but command passed: %s %s\n' "$command_kind" "$package_path" >&2
        exit 1
    fi
    if ! grep -Eq "$expected_pattern" <<<"$output"; then
        printf 'failure diagnostic mismatch: %s\n%s\n' "$package_path" "$output" >&2
        exit 1
    fi
    printf 'expected failure: %s\n' "$package_path"
}

expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/package_short_decl 'outside function body|syntax error'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/no_new_variable 'no new variables'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/defined_type_assignment 'cannot use raw'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/slice_comparison 'slice can only be compared to nil|invalid operation.*left == right'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/map_struct_field 'cannot assign to struct field'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/final_fallthrough 'cannot fallthrough final case'
expect_failure run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_map_write 'assignment to entry in nil map'
expect_failure run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_pointer 'nil pointer dereference|invalid memory address'

printf 'module01 teaching failures: PASS\n'
