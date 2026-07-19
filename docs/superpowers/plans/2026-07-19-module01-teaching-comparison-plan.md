# Module 01 Teaching Project Comparison Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Compare `/Users/zhangshiyu/iotestgo2/module01_basics` with `/Users/zhangshiyu/iotestgo` and produce an evidence-based inventory of missing Go language semantics, teaching examples, and migration priorities.

**Architecture:** Treat the old project as a semantic example archive and the new project as a paced course system. Build a topic-to-topic matrix, distinguish “missing”, “present but shallow”, “intentionally deferred”, and “should not migrate”, then turn the findings into a reviewable research document with concrete example designs.

**Tech Stack:** Markdown, Go source inspection, `rg`, `go test`, `go run`, Git.

## Global Constraints

- The old project is source material, not a replacement for the new course structure.
- Preserve the new project’s Block/Lab/assessment workflow and Go 1.16-compatible student scope where applicable.
- Prefer examples that expose an observable result or compile-time rule over lists of API names.
- Mark dangerous, version-sensitive, or production-discouraged behavior as dark corners instead of presenting it as normal style.
- Every proposed migration must identify its teaching role, placement, observable output, and recommended depth.

---

### Task 1: Establish the comparison baseline

**Files:**
- Inspect: `/Users/zhangshiyu/iotestgo2/module01_basics/README.md`
- Inspect: `/Users/zhangshiyu/iotestgo2/module01_basics/blocks/*/README.md`
- Inspect: `/Users/zhangshiyu/iotestgo2/module01_basics/assessments/*`
- Inspect: `/Users/zhangshiyu/iotestgo2/module01_basics/instructor/*`
- Inspect: `/Users/zhangshiyu/iotestgo/module01_basics` equivalent materials when present

- [ ] Record the new project’s four Blocks, core time box, labs, assessments, bonus materials, and stated version constraints.
- [ ] Record the old project’s topic entry points from `main.go`: variables, basic syntax, collections, functions, defer, errors, structs, methods, interfaces, reflection, APIs, concurrency, I/O, networking, systems, and cgo.
- [ ] Classify the old project’s material as core Module 01, later-block material, bonus/dark-corner material, or outside this course.

### Task 2: Build the semantic coverage matrix

**Files:**
- Create: `docs/research/module01-old-new-teaching-comparison.md`

- [ ] For each topic, record old-project evidence, new-project location, coverage status, and migration recommendation.
- [ ] Use the following statuses exactly: `完整保留`, `已覆盖但变薄`, `缺失`, `应延后`, `不建议迁移`.
- [ ] Separate syntax coverage from semantic coverage. For example, a `switch` example counts as syntax coverage, but multiple-value cases, implicit termination, expressionless switch, and explicit `fallthrough` are separate semantic checks.

### Task 3: Extract the old project’s teaching mechanisms

**Files:**
- Inspect: `/Users/zhangshiyu/iotestgo/myvar/myvar.go`
- Inspect: `/Users/zhangshiyu/iotestgo/mygolang/mygolang.go`
- Inspect: `/Users/zhangshiyu/iotestgo/mycollection/*`
- Inspect: `/Users/zhangshiyu/iotestgo/myfunc/*`
- Inspect: `/Users/zhangshiyu/iotestgo/myexception/myexception.go`
- Inspect: `/Users/zhangshiyu/iotestgo/myoop/*`
- Inspect: `/Users/zhangshiyu/iotestgo/myapi/*`

- [ ] Identify examples whose value comes from a surprising output, not from their surface topic label.
- [ ] Identify examples that contrast two almost-identical snippets, such as array range value versus index assignment, value receiver versus pointer receiver, and `defer` argument evaluation versus closure evaluation.
- [ ] Identify examples that show compiler restrictions or panic boundaries, then classify whether they are suitable for core teaching or a guided dark corner.

### Task 4: Produce the missing-essence backlog

**Files:**
- Modify: `docs/research/module01-old-new-teaching-comparison.md`

- [ ] Add a high-priority list for Module 01: constants and `iota`, implicit ConstSpec repetition, blank identifier, type definitions, explicit conversions, parse errors, control-flow labels, error construction, and function/closure semantics.
- [ ] Add a medium-priority list for later Blocks: custom sort interfaces, map-of-map initialization, map value write-back, method sets, interface/type switch, embedding, and reflection boundaries.
- [ ] Add a defer/bonus list: panic/recover, unsafe pointer arithmetic, reflection mutability, and version-sensitive range closure behavior.
- [ ] For each item, provide a minimal code shape and the behavior students should predict before running it.

### Task 5: Review migration quality

**Files:**
- Inspect: `docs/research/module01-old-new-teaching-comparison.md`
- Inspect: `/Users/zhangshiyu/iotestgo2/module01_basics/blocks/01_go_basics/demo/*`

- [ ] Confirm the `const`/`iota` case explicitly explains that omitted declarations repeat the previous complete expression, rather than saying that `iota` is “interrupted”.
- [ ] Confirm every recommended demo prints the values needed to observe the rule.
- [ ] Confirm no recommendation duplicates a new demo unnecessarily; merge into the existing teaching flow where the user has requested a less fragmented structure.
- [ ] Run `git diff --check` and `go test ./module01_basics/...` after any implementation follow-up.

### Task 6: Finalize the research handoff

- [ ] Add an executive summary, comparison matrix, omission list, priority roadmap, and explicit “do not migrate wholesale” section.
- [ ] State which conclusions are directly evidenced by source files and which are design recommendations.
- [ ] Record the document path and leave code changes separate from the research commit unless implementation is explicitly requested.
