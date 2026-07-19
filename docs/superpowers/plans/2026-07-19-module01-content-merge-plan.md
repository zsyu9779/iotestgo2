# Module 01 Content Merge Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Incrementally merge the early `/Users/zhangshiyu/iotestgo` teaching material into the current `module01_basics` course without replacing its existing structure, while promoting `dark_corners` into core teaching touchpoints.

**Architecture:** Keep the current four Blocks, Labs, Scorebook, Homework, and verification commands as the canonical spine. Add focused runnable demos and teaching sections beside existing demos; link dark-corner material from the relevant Block and instructor documents while preserving the original source paths. Update assessments and runbook guidance after the content additions are stable.

**Tech Stack:** Go 1.25 root module for current course examples; Go 1.16 isolated Task Manager module; Markdown; standard library only for teaching demos.

## Global Constraints

- Do not delete or wholesale replace existing `module01_basics` demos, Labs, Scorebook, Task Manager, or verification flow.
- Treat `/Users/zhangshiyu/iotestgo` as a source of teaching material; clean inaccurate comments, stale APIs, and misleading Java analogies before reuse.
- Promote `module01_basics/bonus/dark_corners` through core Block links and instructor sequencing, but preserve those original files and paths.
- Keep interfaces, reflection, concurrency, Cgo, files, networking, generics, and complex function patterns out of the Module 01 core.
- Every new Go example must be formatted, vettable, runnable from the documented command, and independent of Module 02.
- Preserve the existing student/teacher disclosure boundary.

---

### Task 1: Expand Block 1 language demos and guide

**Files:**

- Create `module01_basics/blocks/01_go_basics/demo/04_zero_values/main.go`.
- Create `module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go`.
- Create `module01_basics/blocks/01_go_basics/demo/06_control_flow_edges/main.go`.
- Modify `module01_basics/blocks/01_go_basics/README.md`.
- Modify `module01_basics/instructor/DEMO_NOTES.md` and `RUNBOOK.md`.

- [ ] Add observable checks for zero values, String `Fields`/`TrimSpace`/`range`, multi-value `case`, expressionless `switch`, and actual `fallthrough` output.
- [ ] Implement focused deterministic demos from `myvar`, `mygolang`, and `myapi` without replacing the current demos.
- [ ] Document `for` forms, `break`, `continue`, multi-value `case`, `switch {}`, `fallthrough` limits, and optional `goto` recognition.
- [ ] Mark core, deep-dive, and cuttable material in instructor notes.
- [ ] Verify with `gofmt`, `go vet ./module01_basics/blocks/01_go_basics/...`, `go test ./module01_basics/blocks/01_go_basics/...`, and each new `go run` command.
- [ ] Commit as `feat: enrich module01 basics content`.

### Task 2: Expand Block 2 collections and promote map/string dark corners

**Files:**

- Create `module01_basics/blocks/02_collections/demo/06_slice_map_edges/main.go`.
- Create `module01_basics/blocks/02_collections/demo/07_string_utf8_edges/main.go`.
- Modify existing Array/Slice and Map/String demos, Block 2 README/Lab README, instructor notes/Runbook, and `bonus/README.md`.

- [ ] Add observable checks for Array range copies, capacity-sensitive Slice append, nil Slice versus nil Map, comma-ok, Map value write-back, `range string`, and common `strings` functions.
- [ ] Implement deterministic demos from `mycollection`, `myapi/mystr.go`, and `myapi/mystring2.go`; keep `unsafe.Pointer` out of core.
- [ ] Link `dark_corners/map` and `dark_corners/string` from Block 2 with phenomenon, cause, recommendation, and version/scope annotations.
- [ ] Verify with formatting, vet, normal tests, exercise-tagged TextStats tests, and both new demo commands.
- [ ] Commit as `feat: enrich module01 collections content`.

### Task 3: Expand Block 3 modeling semantics

**Files:**

- Create `module01_basics/blocks/03_modeling/demo/08_struct_zero_values/main.go`.
- Create `module01_basics/blocks/03_modeling/demo/09_copy_and_receivers/main.go`.
- Modify existing pointer/struct demos, Block 3 README/Lab README, and instructor notes/Runbook.

- [ ] Add observable checks for struct zero values, value versus pointer mutation, Snapshot isolation, pointer receiver mutation, nil pointers, and promoted fields.
- [ ] Implement focused demos from `myoop/mystruct` and `mymethod`.
- [ ] Correct pointer-parameter versus pointer-receiver terminology and replace inheritance language with composition language.
- [ ] Keep Embedding optional and do not change the Student Lab contract.
- [ ] Verify with formatting, vet, normal tests, and exercise-tagged Student tests.
- [ ] Commit as `feat: enrich module01 modeling content`.

### Task 4: Expand Block 4 functions, defer, and learner-authored testing

**Files:**

- Create `module01_basics/blocks/04_functions_testing/demo/10_function_forms/main.go`.
- Create `module01_basics/blocks/04_functions_testing/demo/11_defer_edges/main.go`.
- Add one focused exercise-tagged learner-authored test under `module01_basics/blocks/04_functions_testing/lab/starter/`.
- Modify Block 4 README/Lab README and instructor notes/Runbook.

- [ ] Write and run the new focused test first, confirming it fails for the intended behavior.
- [ ] Implement the smallest starter change needed without changing public signatures or weakening existing tests.
- [ ] Add demos for named returns, variadic functions, function types, higher-order functions, closure state, defer LIFO, argument evaluation, and closure capture from `myfunc`/`mydefer`.
- [ ] Explain `_test.go`, `TestXxx`, `testing.T`, `t.Run`, table-driven tests, and `-run`; keep `panic/recover` in Module 02.
- [ ] Verify with formatting, vet, normal tests, and exercise-tagged Block 4 tests.
- [ ] Commit as `feat: enrich module01 functions and testing`.

### Task 5: Integrate dark corners, assessments, and course navigation

**Files:**

- Modify `module01_basics/README.md`.
- Modify `module01_basics/assessments/entry_quiz.md` and `exit_quiz.md`.
- Modify `module01_basics/instructor/RUBRIC.md`, `DEMO_NOTES.md`, and `RUNBOOK.md`.
- Modify `module01_basics/bonus/README.md`.

- [ ] Add outcome-oriented questions for zero values, String basics, `range`, multi-value `case`, `fallthrough`, nil Map, Slice aliasing, and basic test authoring.
- [ ] Keep existing hands-on blocks and stop rules; mark new material as core, deep-dive, or cuttable detours.
- [ ] Make new demos discoverable from the canonical Module README and explain that dark corners are core-linked while other Bonus families remain instructor-selected.
- [ ] Add rubric evidence for prediction and explanation without requiring implementation details.
- [ ] Validate all new Markdown links and check the student/teacher disclosure boundary.
- [ ] Commit as `docs: integrate module01 content navigation and assessments`.

### Task 6: Full verification and handoff

**Files:** Verify all files changed by Tasks 1–5.

- [ ] Run `gofmt -l $(git ls-files 'module01_basics/**/*.go')` and `git diff --check`; expect no output/errors.
- [ ] Run `make module01-verify`; expect `go vet`, Module 01 tests, and Task Manager teacher verification to pass.
- [ ] Run all five exercise-tagged starter test commands from the Module README.
- [ ] Run every new demo command from the repository root and inspect deterministic output.
- [ ] Review changed paths, disclosure boundaries, and commit history.
- [ ] Commit only verification-driven fixes.
