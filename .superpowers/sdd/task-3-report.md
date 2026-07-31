# Module 01 Task 3 Report

## Scope

- Added the Collection demos for slice `copy(dst, src)` and per-level nested Map initialization.
- Added stable output-contract assertions for both teaching results.
- Updated the Block 2 README, Demo Notes, and Runbook so `copy` remains core content and nested Map initialization remains deep-dive, cuttable content.

## RED → GREEN Evidence

1. Added the two output assertions to `verify_demo_contracts.sh`.
2. Ran `make module01-demo-contracts`; it failed as expected because `copy count=3 dst=[10 20 30 0 0] src=[10 20 30]` was absent.
3. Added the minimal demo implementations and documentation updates.
4. Ran `gofmt`, `make module01-demo-contracts`, `go test ./module01_basics/blocks/02_collections/...`, and `git diff --check`; all passed.

## Changed Files

- `module01_basics/blocks/02_collections/demo/04_arrays_slices/main.go`
- `module01_basics/blocks/02_collections/demo/05_maps_strings/main.go`
- `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- `module01_basics/blocks/02_collections/README.md`
- `module01_basics/instructor/DEMO_NOTES.md`
- `module01_basics/instructor/RUNBOOK.md`

## Concerns

None. The nested Map result is deterministic because the demo prints a direct lookup rather than relying on Map iteration order.
