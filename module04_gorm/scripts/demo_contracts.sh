#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

assert_output() {
	local package="$1"
	local expected="$2"
	local output
	if ! output="$(go run "$package" 2>&1)"; then
		printf 'demo failed: %s\n%s\n' "$package" "$output" >&2
		exit 1
	fi
	if ! grep -Fq "$expected" <<<"$output"; then
		printf 'missing %q in output from %s\n%s\n' "$expected" "$package" "$output" >&2
		exit 1
	fi
}

assert_output ./module04_gorm/01_setup 'connection=ok'
assert_output ./module04_gorm/02_models_migrations 'migration=email'
assert_output ./module04_gorm/03_crud 'map_zero=0'
assert_output ./module04_gorm/04_queries_hooks/queries 'joins='
assert_output ./module04_gorm/04_queries_hooks/nplusone 'preload=2'
assert_output ./module04_gorm/04_queries_hooks/hooks 'blocked_negative'
assert_output ./module04_gorm/05_transactions 'second_step=rolled_back'
assert_output ./module04_gorm/06_raw_sql 'injection_safe=parameters'

echo 'module04 demo contracts: PASS'
