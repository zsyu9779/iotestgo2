# Module 01 Standard Course Template Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rebuild `module01_basics` as a 09:30–16:00 instructor-and-learner course template with four teaching blocks, a Scorebook classroom lab, and a standalone Gitee-graded Task Manager homework package.

**Architecture:** Existing lesson directories move into four journey-oriented blocks so Git history remains useful. Each block owns a complete demo, a compilable starter, an opt-in failing exercise test, a passing solution, and learner instructions. The homework is a nested, Go 1.16-compatible module so it can be copied to Gitee independently while the root module remains on Go 1.25.

**Tech Stack:** Go standard library, Go `testing`, POSIX shell, Make, Gitee Go `.workflow` YAML, Markdown.

## Global Constraints

- Target learners are university students with Java experience and no prior Go requirement.
- Class runs 09:30–12:00 and 13:00–16:00, with a 60-minute lunch and two 10-minute short breaks.
- Net teaching and practice time is 310 minutes.
- Student hands-on time must be at least 50% of net class time.
- Task Manager is a one-week homework assignment, not the in-class capstone.
- The GitHub teacher repository contains demos, starters, public tests, solutions, and instructor material.
- The later Gitee learner repository receives only approved demos and `homework/task_manager/student_pack`.
- Task Manager must use only the Go standard library.
- Local grading and Gitee grading must both call `scripts/grade.sh`.
- Gitee CI has no push or PR trigger; the instructor selects a student branch and runs the pipeline manually.
- The student pack stays compatible with the officially documented Gitee `build@golang` Go 1.16 runtime; no homework requirement may need generics.
- Module 02–06 source code is out of scope.
- Do not make JSON persistence, third-party CLI packages, hidden tests, or plagiarism detection Week 1 requirements.
- Preserve unrelated existing user changes and do not fix repository-wide Module 05/06 failures in this plan.
- Every Block README uses this sequence: `学习结果`, `时间盒`, `前置知识`, `Java 对比`, `讲师 Demo`, `学员任务`, `验收命令`, `常见错误`, `三级提示`, `复盘问题`, `Bonus`.

---

## File Map

### Course entry and instructor material

- `module01_basics/README.md` — learner-facing course map and commands.
- `module01_basics/instructor/RUNBOOK.md` — minute-by-minute delivery plan.
- `module01_basics/instructor/DEMO_NOTES.md` — demo prompts, failure demonstrations, and rescue paths.
- `module01_basics/instructor/RUBRIC.md` — common lab scoring language.
- `module01_basics/assessments/{entry_quiz,exit_quiz,answer_key}.md` — class diagnostics.

### Teaching blocks

- `module01_basics/blocks/01_go_basics` — toolchain, values, control flow, functions, grade calculator lab.
- `module01_basics/blocks/02_collections` — arrays, slices, maps, strings, text statistics lab.
- `module01_basics/blocks/03_modeling` — pointers, structs, methods, student model lab.
- `module01_basics/blocks/04_functions_testing` — function values, closures, defer, testing lab.

Each block contains `README.md`, `demo/`, and `lab/{README.md,starter/,solution/}`.

### Integrated lab and homework

- `module01_basics/integrated_lab/scorebook` — in-class transfer exercise.
- `module01_basics/homework/task_manager/student_pack` — independently copyable learner repository.
- `module01_basics/homework/task_manager/teacher` — solution, rubric, release checklist, troubleshooting.

### Bonus and shared standards

- `module01_basics/bonus/{generics,data_structures,dark_corners,function_patterns}` — material removed from the core 310 minutes.
- `Makefile` — root Module 01 verification targets.
- `docs/course-module-standard.md` — condensed reusable six-week module standard.
- `Golang_Backend_Training_Syllabus.md` — aligns Module 01 with the actual one-day course.
- `docs/module01_basics_lesson_plan.md` and `docs/module01_02_lecture_cheatsheet.md` — become compatibility pointers to the new instructor entry.

---

### Task 1: Build Block 1 — Go Basics

**Files:**
- Move: `module01_basics/01_hello` → `module01_basics/blocks/01_go_basics/demo/01_hello`
- Move: `module01_basics/02_vars_types` → `module01_basics/blocks/01_go_basics/demo/02_vars_types`
- Move: `module01_basics/03_control_funcs` → `module01_basics/blocks/01_go_basics/demo/03_control_funcs`
- Create: `module01_basics/blocks/01_go_basics/README.md`
- Create: `module01_basics/blocks/01_go_basics/lab/README.md`
- Create: `module01_basics/blocks/01_go_basics/lab/starter/grade.go`
- Create: `module01_basics/blocks/01_go_basics/lab/starter/grade_exercise_test.go`
- Create: `module01_basics/blocks/01_go_basics/lab/solution/grade.go`
- Test: `module01_basics/blocks/01_go_basics/lab/solution/grade_test.go`

**Interfaces:**
- Produces: `grade.Grade(score int) (string, error)` and `grade.ErrScoreOutOfRange`.
- Consumes: Go standard library `errors` only.

- [ ] **Step 1: Move the three existing demonstrations into the new block**

Run:

```bash
mkdir -p module01_basics/blocks/01_go_basics/demo
git mv module01_basics/01_hello module01_basics/blocks/01_go_basics/demo/01_hello
git mv module01_basics/02_vars_types module01_basics/blocks/01_go_basics/demo/02_vars_types
git mv module01_basics/03_control_funcs module01_basics/blocks/01_go_basics/demo/03_control_funcs
```

Expected: all three original programs retain their individual `package main` directories and remain runnable.

- [ ] **Step 2: Write the solution behavior test first**

Create `lab/solution/grade_test.go`:

```go
package grade

import (
	"errors"
	"testing"
)

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "excellent lower bound", score: 90, want: "A"},
		{name: "good", score: 82, want: "B"},
		{name: "pass", score: 60, want: "D"},
		{name: "fail", score: 59, want: "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grade(tt.score)
			if err != nil {
				t.Fatalf("Grade(%d) returned error: %v", tt.score, err)
			}
			if got != tt.want {
				t.Fatalf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestGradeRejectsOutOfRange(t *testing.T) {
	for _, score := range []int{-1, 101} {
		if _, err := Grade(score); !errors.Is(err, ErrScoreOutOfRange) {
			t.Fatalf("Grade(%d) error = %v, want ErrScoreOutOfRange", score, err)
		}
	}
}
```

- [ ] **Step 3: Run the solution test to verify it fails**

Run:

```bash
go test ./module01_basics/blocks/01_go_basics/lab/solution
```

Expected: FAIL with `undefined: Grade` and `undefined: ErrScoreOutOfRange`.

- [ ] **Step 4: Implement the minimal passing solution**

Create `lab/solution/grade.go`:

```go
package grade

import "errors"

var ErrScoreOutOfRange = errors.New("score must be between 0 and 100")

func Grade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", ErrScoreOutOfRange
	}

	switch {
	case score >= 90:
		return "A", nil
	case score >= 80:
		return "B", nil
	case score >= 70:
		return "C", nil
	case score >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}
```

- [ ] **Step 5: Add a compilable starter and opt-in exercise test**

Create `lab/starter/grade.go` with the same error and signature, returning `"", nil`. Create `grade_exercise_test.go` by copying the solution test and prepending:

```go
//go:build exercise
// +build exercise

package grade
```

Expected: default root tests compile the starter; `go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter` fails until the learner implements `Grade`.

- [ ] **Step 6: Write the block and lab instructions**

`blocks/01_go_basics/README.md` must include these headings exactly:

```markdown
# Block 1：Go Basics
## 学习结果
## 时间盒：45 分钟
## 前置知识
## Java 对比
## 讲师 Demo
## 学员任务
## 验收命令
## 常见错误
## 三级提示
## 复盘问题
## Bonus
```

`lab/README.md` must state the A/B/C/D/F thresholds, the 0–100 invariant, the command `go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter`, and three hint levels without pasting the switch implementation.

- [ ] **Step 7: Verify Block 1**

Run:

```bash
find module01_basics/blocks/01_go_basics -type f -name '*.go' -exec gofmt -w {} +
go test ./module01_basics/blocks/01_go_basics/...
go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter
```

Expected: default tests PASS; exercise-tagged starter test FAILS with a grade mismatch.

- [ ] **Step 8: Commit Block 1**

```bash
git add module01_basics/blocks/01_go_basics
git commit -m "feat(course): build module 01 basics block"
```

---

### Task 2: Build Block 2 — Collections

**Files:**
- Move: `module01_basics/04_arrays_slices` → `module01_basics/blocks/02_collections/demo/04_arrays_slices`
- Move: `module01_basics/05_maps_strings` → `module01_basics/blocks/02_collections/demo/05_maps_strings`
- Create: `module01_basics/blocks/02_collections/README.md`
- Create: `module01_basics/blocks/02_collections/lab/README.md`
- Create: `module01_basics/blocks/02_collections/lab/starter/textstats.go`
- Create: `module01_basics/blocks/02_collections/lab/starter/textstats_exercise_test.go`
- Create: `module01_basics/blocks/02_collections/lab/solution/textstats.go`
- Test: `module01_basics/blocks/02_collections/lab/solution/textstats_test.go`

**Interfaces:**
- Produces: `textstats.Stats` and `textstats.Analyze(text string) Stats`.
- Consumes: `strings.Fields`, `strings.ToLower`, `utf8.RuneCountInString`.

- [ ] **Step 1: Move the existing demonstrations**

```bash
mkdir -p module01_basics/blocks/02_collections/demo
git mv module01_basics/04_arrays_slices module01_basics/blocks/02_collections/demo/04_arrays_slices
git mv module01_basics/05_maps_strings module01_basics/blocks/02_collections/demo/05_maps_strings
```

- [ ] **Step 2: Write the failing text analysis test**

Create `lab/solution/textstats_test.go`:

```go
package textstats

import "testing"

func TestAnalyze(t *testing.T) {
	got := Analyze("Go go 你好")
	if got.Bytes != 12 {
		t.Fatalf("Bytes = %d, want 12", got.Bytes)
	}
	if got.Runes != 8 {
		t.Fatalf("Runes = %d, want 8", got.Runes)
	}
	if got.Words != 3 {
		t.Fatalf("Words = %d, want 3", got.Words)
	}
	if got.Frequencies["go"] != 2 || got.Frequencies["你好"] != 1 {
		t.Fatalf("Frequencies = %#v", got.Frequencies)
	}
}

func TestAnalyzeEmptyText(t *testing.T) {
	got := Analyze("")
	if got.Bytes != 0 || got.Runes != 0 || got.Words != 0 || len(got.Frequencies) != 0 {
		t.Fatalf("Analyze empty = %#v", got)
	}
}
```

- [ ] **Step 3: Run the test and observe the missing interface**

Run `go test ./module01_basics/blocks/02_collections/lab/solution`.

Expected: FAIL with `undefined: Analyze`.

- [ ] **Step 4: Implement text analysis**

Create `lab/solution/textstats.go`:

```go
package textstats

import (
	"strings"
	"unicode/utf8"
)

type Stats struct {
	Bytes       int
	Runes       int
	Words       int
	Frequencies map[string]int
}

func Analyze(text string) Stats {
	words := strings.Fields(text)
	frequencies := make(map[string]int, len(words))
	for _, word := range words {
		frequencies[strings.ToLower(word)]++
	}
	return Stats{
		Bytes:       len(text),
		Runes:       utf8.RuneCountInString(text),
		Words:       len(words),
		Frequencies: frequencies,
	}
}
```

- [ ] **Step 5: Add starter, tagged test, and instructions**

The starter defines `Stats` and returns an empty `Stats{Frequencies: map[string]int{}}`. The tagged test is the solution test with both `exercise` build-tag lines. The lab instructions explicitly define whitespace tokenization and case-insensitive English words so punctuation behavior is not ambiguous. The Block README follows the global heading sequence and uses `## 时间盒：75 分钟`.

- [ ] **Step 6: Verify Block 2**

Run:

```bash
find module01_basics/blocks/02_collections -type f -name '*.go' -exec gofmt -w {} +
go test ./module01_basics/blocks/02_collections/...
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter
```

Expected: default tests PASS; starter exercise test FAILS on `Bytes`.

- [ ] **Step 7: Commit Block 2**

```bash
git add module01_basics/blocks/02_collections
git commit -m "feat(course): build module 01 collections block"
```

---

### Task 3: Build Block 3 — Modeling

**Files:**
- Move: `module01_basics/06_pointers` → `module01_basics/blocks/03_modeling/demo/06_pointers`
- Move: `module01_basics/07_structs_methods` → `module01_basics/blocks/03_modeling/demo/07_structs_methods`
- Create: `module01_basics/blocks/03_modeling/README.md`
- Create: `module01_basics/blocks/03_modeling/lab/README.md`
- Create: `module01_basics/blocks/03_modeling/lab/starter/student.go`
- Create: `module01_basics/blocks/03_modeling/lab/starter/student_exercise_test.go`
- Create: `module01_basics/blocks/03_modeling/lab/solution/student.go`
- Test: `module01_basics/blocks/03_modeling/lab/solution/student_test.go`

**Interfaces:**
- Produces: `student.New(id int, name string, score int) (*Student, error)`, `(*Student).Rename(name string) error`, `(*Student).UpdateScore(score int) error`, `Student.Snapshot() Student`.
- Produces errors: `ErrInvalidID`, `ErrInvalidName`, `ErrInvalidScore`.

- [ ] **Step 1: Move the existing pointer and struct demos**

```bash
mkdir -p module01_basics/blocks/03_modeling/demo
git mv module01_basics/06_pointers module01_basics/blocks/03_modeling/demo/06_pointers
git mv module01_basics/07_structs_methods module01_basics/blocks/03_modeling/demo/07_structs_methods
```

- [ ] **Step 2: Write constructor and mutation tests**

Create `lab/solution/student_test.go` with tests that require: positive ID, trimmed non-empty name, score 0–100, pointer-receiver mutation, and snapshot copy isolation.

```go
package student

import "testing"

func TestStudentLifecycle(t *testing.T) {
	s, err := New(1, " Alice ", 80)
	if err != nil {
		t.Fatal(err)
	}
	if s.Name != "Alice" {
		t.Fatalf("Name = %q, want Alice", s.Name)
	}
	if err := s.UpdateScore(95); err != nil {
		t.Fatal(err)
	}
	if s.Score != 95 {
		t.Fatalf("Score = %d, want 95", s.Score)
	}
	copy := s.Snapshot()
	copy.Name = "changed copy"
	if s.Name != "Alice" {
		t.Fatalf("snapshot mutation leaked to original: %#v", s)
	}
}
```

- [ ] **Step 3: Run the test and verify the missing model**

Run `go test ./module01_basics/blocks/03_modeling/lab/solution`.

Expected: FAIL with `undefined: New`.

- [ ] **Step 4: Implement the Student model**

Use this exact shape:

```go
type Student struct {
	ID    int
	Name  string
	Score int
}

func New(id int, name string, score int) (*Student, error)
func (s *Student) Rename(name string) error
func (s *Student) UpdateScore(score int) error
func (s Student) Snapshot() Student
```

Validation helpers trim names with `strings.TrimSpace`; score validation accepts both 0 and 100. `Snapshot` returns `s` by value.

- [ ] **Step 5: Add starter, tagged tests, and the receiver-choice lab guide**

The starter declares all three sentinel errors and returns `&Student{}, nil` from `New`; `Rename` and `UpdateScore` return `nil` without mutating, while `Snapshot` returns the receiver by value. This keeps the starter compilable and makes the tagged test fail on the first state assertion. The README follows the global heading sequence, uses `## 时间盒：65 分钟`, and asks learners to explain why `Rename` and `UpdateScore` use pointer receivers while `Snapshot` uses a value receiver.

- [ ] **Step 6: Verify and commit Block 3**

```bash
find module01_basics/blocks/03_modeling -type f -name '*.go' -exec gofmt -w {} +
go test ./module01_basics/blocks/03_modeling/...
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter
git add module01_basics/blocks/03_modeling
git commit -m "feat(course): build module 01 modeling block"
```

Expected: solution PASS; starter exercise FAIL; commit succeeds.

---

### Task 4: Build Block 4 — Functions and Testing

**Files:**
- Move: `module01_basics/09_advanced_functions` → `module01_basics/blocks/04_functions_testing/demo/09_advanced_functions`
- Create: `module01_basics/blocks/04_functions_testing/README.md`
- Create: `module01_basics/blocks/04_functions_testing/lab/README.md`
- Create: `module01_basics/blocks/04_functions_testing/lab/starter/scores.go`
- Create: `module01_basics/blocks/04_functions_testing/lab/starter/scores_exercise_test.go`
- Create: `module01_basics/blocks/04_functions_testing/lab/solution/scores.go`
- Test: `module01_basics/blocks/04_functions_testing/lab/solution/scores_test.go`

**Interfaces:**
- Produces: `scores.Filter(values []int, keep func(int) bool) []int`.
- Produces: `scores.AtLeast(min int) func(int) bool`.
- Produces: `scores.WithAudit(name string, record func(string), operation func())`.

- [ ] **Step 1: Move the advanced-functions directory into the block demo**

```bash
mkdir -p module01_basics/blocks/04_functions_testing/demo
git mv module01_basics/09_advanced_functions module01_basics/blocks/04_functions_testing/demo/09_advanced_functions
```

- [ ] **Step 2: Write function-value, closure, and defer tests**

```go
package scores

import (
	"reflect"
	"testing"
)

func TestFilterWithClosure(t *testing.T) {
	got := Filter([]int{59, 60, 75, 90}, AtLeast(60))
	want := []int{60, 75, 90}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Filter = %v, want %v", got, want)
	}
}

func TestWithAuditRecordsEndAfterOperation(t *testing.T) {
	var events []string
	WithAudit("average", func(event string) {
		events = append(events, event)
	}, func() {
		events = append(events, "operation")
	})
	want := []string{"start:average", "operation", "end:average"}
	if !reflect.DeepEqual(events, want) {
		t.Fatalf("events = %v, want %v", events, want)
	}
}
```

- [ ] **Step 3: Run the test and observe missing functions**

Run `go test ./module01_basics/blocks/04_functions_testing/lab/solution`.

Expected: FAIL with `undefined: Filter`.

- [ ] **Step 4: Implement the solution**

```go
package scores

func Filter(values []int, keep func(int) bool) []int {
	filtered := make([]int, 0, len(values))
	for _, value := range values {
		if keep(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func AtLeast(min int) func(int) bool {
	return func(value int) bool { return value >= min }
}

func WithAudit(name string, record func(string), operation func()) {
	record("start:" + name)
	defer record("end:" + name)
	operation()
}
```

- [ ] **Step 5: Add starter, tagged tests, and learner instructions**

The starter returns `nil` from `Filter`, returns a predicate that is always false, and calls only `operation` from `WithAudit`. The lab guide must have learners run one named test at a time with `-run` before the full package. The Block README follows the global heading sequence and uses `## 时间盒：45 分钟`.

- [ ] **Step 6: Verify and commit Block 4**

```bash
find module01_basics/blocks/04_functions_testing -type f -name '*.go' -exec gofmt -w {} +
go test ./module01_basics/blocks/04_functions_testing/...
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter
git add module01_basics/blocks/04_functions_testing
git commit -m "feat(course): build module 01 functions testing block"
```

Expected: solution PASS; starter exercise FAIL.

---

### Task 5: Build the Scorebook Integrated Classroom Lab

**Files:**
- Create: `module01_basics/integrated_lab/scorebook/README.md`
- Create: `module01_basics/integrated_lab/scorebook/starter/scorebook.go`
- Create: `module01_basics/integrated_lab/scorebook/starter/scorebook_exercise_test.go`
- Create: `module01_basics/integrated_lab/scorebook/solution/scorebook.go`
- Test: `module01_basics/integrated_lab/scorebook/solution/scorebook_test.go`

**Interfaces:**
- Produces: `scorebook.New() *Scorebook`.
- Produces: `(*Scorebook).Add(name string, score int) (Student, error)`.
- Produces: `(*Scorebook).Find(id int) (Student, error)`.
- Produces: `(*Scorebook).UpdateScore(id int, score int) error`.
- Produces: `(*Scorebook).Average() (float64, error)`.
- Produces: `(*Scorebook).CountByGrade() map[string]int`.

- [ ] **Step 1: Write the integrated behavior tests**

Tests must cover sequential IDs, returned-copy isolation, unknown IDs, empty average, average calculation, and grade counts. Use sentinel errors `ErrInvalidName`, `ErrInvalidScore`, `ErrStudentNotFound`, and `ErrEmptyScorebook` with `errors.Is`.

Core test example:

```go
func TestScorebookWorkflow(t *testing.T) {
	book := New()
	alice, err := book.Add("Alice", 90)
	if err != nil {
		t.Fatal(err)
	}
	bob, err := book.Add("Bob", 70)
	if err != nil {
		t.Fatal(err)
	}
	if alice.ID != 1 || bob.ID != 2 {
		t.Fatalf("IDs = %d, %d; want 1, 2", alice.ID, bob.ID)
	}
	if err := book.UpdateScore(bob.ID, 80); err != nil {
		t.Fatal(err)
	}
	average, err := book.Average()
	if err != nil || average != 85 {
		t.Fatalf("Average = %v, %v; want 85, nil", average, err)
	}
}
```

- [ ] **Step 2: Run the test to confirm the integrated lab is absent**

Run `go test ./module01_basics/integrated_lab/scorebook/solution`.

Expected: FAIL with missing package files or undefined symbols.

- [ ] **Step 3: Implement the minimal Scorebook solution**

Use `map[int]*Student` plus `nextID int`. `Add` stores an internal pointer but returns a value copy. `Find` returns a value copy. `Average` iterates the map and returns `ErrEmptyScorebook` for zero students. `CountByGrade` uses a private `grade(score int) string` helper with the Block 1 thresholds.

- [ ] **Step 4: Add the starter and exercise tests**

Starter includes all exported types, errors, and methods so it compiles. Methods return zero values or the relevant not-found/empty error. Tagged tests mirror the solution tests. Do not import any block solution; the lab verifies transfer rather than reuse.

- [ ] **Step 5: Write the 40-minute lab guide**

Divide the guide into four 10-minute checkpoints: model/add, find/update, average, grade counts/tests. Include a stopping rule after each checkpoint and three levels of hints.

- [ ] **Step 6: Verify and commit the integrated lab**

```bash
find module01_basics/integrated_lab/scorebook -type f -name '*.go' -exec gofmt -w {} +
go test ./module01_basics/integrated_lab/scorebook/...
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter
git add module01_basics/integrated_lab/scorebook
git commit -m "feat(course): add module 01 scorebook lab"
```

Expected: solution PASS; starter exercise FAIL.

---

### Task 6: Package the Task Manager Homework and Gitee Grader

**Files:**
- Move: `module01_basics/project_task_manager` → `module01_basics/homework/task_manager/teacher/solution`
- Create: `module01_basics/homework/task_manager/student_pack/README.md`
- Create: `module01_basics/homework/task_manager/student_pack/go.mod`
- Create: `module01_basics/homework/task_manager/student_pack/Makefile`
- Create: `module01_basics/homework/task_manager/student_pack/cmd/taskmanager/main.go`
- Create: `module01_basics/homework/task_manager/student_pack/taskmanager/task.go`
- Create: `module01_basics/homework/task_manager/student_pack/taskmanager/manager.go`
- Test: `module01_basics/homework/task_manager/student_pack/taskmanager/manager_test.go`
- Create: `module01_basics/homework/task_manager/student_pack/scripts/grade.sh`
- Create: `module01_basics/homework/task_manager/student_pack/.workflow/GradePipeline.yml`
- Create: `module01_basics/homework/task_manager/teacher/solution/go.mod`
- Create: `module01_basics/homework/task_manager/teacher/solution/cmd/taskmanager/main.go`
- Create: `module01_basics/homework/task_manager/teacher/solution/taskmanager/{task.go,manager.go,manager_test.go}`
- Create: `module01_basics/homework/task_manager/teacher/{RUBRIC,RELEASE_CHECKLIST,TROUBLESHOOTING}.md`

**Interfaces:**
- Produces: `taskmanager.NewManager() *Manager`.
- Produces: `Add(title string) (Task, error)`, `List() []Task`, `Complete(id int) (Task, error)`, `Delete(id int) error`.
- Produces errors: `ErrEmptyTitle`, `ErrTaskNotFound`, and starter-only `ErrNotImplemented`.

- [ ] **Step 1: Move the original homework so its history remains available**

```bash
mkdir -p module01_basics/homework/task_manager/teacher
git mv module01_basics/project_task_manager module01_basics/homework/task_manager/teacher/solution
```

- [ ] **Step 2: Write public homework tests in the student pack**

Create a standalone module with:

```go
module taskmanager

go 1.16
```

Public tests cover trimmed titles, sequential IDs, copy-safe `List`, completion, deletion, and unknown IDs. A key test is:

```go
func TestManagerLifecycle(t *testing.T) {
	m := NewManager()
	first, err := m.Add(" write tests ")
	if err != nil {
		t.Fatal(err)
	}
	if first.ID != 1 || first.Title != "write tests" || first.Completed {
		t.Fatalf("first task = %#v", first)
	}
	completed, err := m.Complete(first.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !completed.Completed {
		t.Fatal("task was not completed")
	}
	if err := m.Delete(first.ID); err != nil {
		t.Fatal(err)
	}
	if len(m.List()) != 0 {
		t.Fatalf("tasks after delete = %#v", m.List())
	}
}
```

- [ ] **Step 3: Add a compilable learner starter**

`task.go` defines:

```go
type Task struct {
	ID        int
	Title     string
	Completed bool
}
```

`manager.go` defines the final method signatures, initializes `nextID: 1`, and returns `ErrNotImplemented` from incomplete methods. `List` returns an empty non-nil slice so the CLI can start. The starter must compile and its public tests must fail at the first `Add` call with `ErrNotImplemented`.

- [ ] **Step 4: Verify the starter fails for the intended reason**

Run:

```bash
cd module01_basics/homework/task_manager/student_pack
go test ./...
```

Expected: FAIL because `Add` returns `ErrNotImplemented`, not because of compilation or environment errors.

- [ ] **Step 5: Refactor the teacher solution into the same interface**

The solution uses `[]*Task` internally, trims titles, assigns sequential IDs, returns copies from every public method, and never exposes the stored pointers. Copy the public tests unchanged into the solution module and add tests for empty titles and `ErrTaskNotFound`.

- [ ] **Step 6: Implement the learner and solution CLI shells**

Both CLIs present commands `add`, `list`, `complete`, `delete`, and `exit` through `bufio.Scanner`. The learner CLI contains complete parsing and printing but depends on the incomplete manager methods, keeping the Week 1 focus on business behavior. The solution CLI uses the same command grammar. CLI behavior is documented for manual review and is not coupled to hidden tests.

- [ ] **Step 7: Create the canonical grading script**

Create `student_pack/scripts/grade.sh`:

```sh
#!/usr/bin/env sh
set -eu

echo "[1/4] 检查 gofmt"
unformatted=$(find . -type f -name '*.go' -exec gofmt -l {} \;)
if [ -n "$unformatted" ]; then
  echo "以下文件未通过 gofmt："
  echo "$unformatted"
  exit 1
fi

echo "[2/4] 执行 go vet"
go vet ./...

echo "[3/4] 执行测试"
go test ./...

echo "[4/4] 构建 CLI"
go build ./...

echo "Task Manager 作业验收通过"
```

`student_pack/Makefile` contains only:

```make
.PHONY: grade
grade:
	@./scripts/grade.sh
```

Make the script executable.

- [ ] **Step 8: Add the manually triggered Gitee pipeline**

Create `.workflow/GradePipeline.yml` using the documented cloud Go plugin and no `triggers` section:

```yaml
version: '1.0'
name: task-manager-grade
displayName: Task Manager 作业验收
stages:
  - stage:
    name: grade
    displayName: 作业验收
    strategy: naturally
    trigger: manual
    steps:
      - step: build@golang
        name: grade_task_manager
        displayName: 格式、Vet、测试和构建
        golangVersion: '1.16'
        commands: |
          chmod +x ./scripts/grade.sh
          ./scripts/grade.sh
```

Document that the instructor selects the student branch in the Gitee UI before clicking Run. Do not add push, PR, or schedule triggers.

- [ ] **Step 9: Write Chinese assignment and teacher operations documentation**

`student_pack/README.md` includes: learning goals, required behavior, explicit exclusions, four milestones, local `make grade`, branch submission, expected CLI examples, and scoring dimensions. Teacher docs include the 100-point rubric, a release checklist that excludes `teacher/`, and fixes for permissions, wrong branches, old Go versions, and failed CI.

- [ ] **Step 10: Verify the grading red/green contract**

Run student pack `make grade`; expect FAIL on tests. Run from `teacher/solution`:

```bash
../../student_pack/scripts/grade.sh
```

Expected: PASS with `Task Manager 作业验收通过`.

- [ ] **Step 11: Commit the homework package**

```bash
git add module01_basics/homework/task_manager
git commit -m "feat(course): package task manager homework"
```

---

### Task 7: Add Course Navigation, Instructor Runbook, Assessments, and Verification

**Files:**
- Rewrite: `module01_basics/README.md`
- Create: `module01_basics/instructor/RUNBOOK.md`
- Create: `module01_basics/instructor/DEMO_NOTES.md`
- Create: `module01_basics/instructor/RUBRIC.md`
- Create: `module01_basics/assessments/entry_quiz.md`
- Create: `module01_basics/assessments/exit_quiz.md`
- Create: `module01_basics/assessments/answer_key.md`
- Create: `Makefile`

**Interfaces:**
- Produces root commands: `module01-verify`, `module01-lab-01` through `module01-lab-04`, `module01-integrated-lab`, and `module01-homework-solution`.

- [ ] **Step 1: Write the root Makefile targets**

Use this behavior:

```make
.PHONY: module01-verify module01-lab-01 module01-lab-02 module01-lab-03 module01-lab-04 module01-integrated-lab module01-homework-solution

module01-lab-01:
	go test ./module01_basics/blocks/01_go_basics/lab/solution

module01-lab-02:
	go test ./module01_basics/blocks/02_collections/lab/solution

module01-lab-03:
	go test ./module01_basics/blocks/03_modeling/lab/solution

module01-lab-04:
	go test ./module01_basics/blocks/04_functions_testing/lab/solution

module01-integrated-lab:
	go test ./module01_basics/integrated_lab/scorebook/solution

module01-homework-solution:
	cd module01_basics/homework/task_manager/teacher/solution && ../../student_pack/scripts/grade.sh

module01-verify:
	@test -z "$$(find module01_basics -type f -name '*.go' -exec gofmt -l {} \;)" || (echo "Module 01 存在未格式化 Go 文件" && exit 1)
	go vet ./module01_basics/...
	go test ./module01_basics/...
	$(MAKE) module01-homework-solution
```

- [ ] **Step 2: Rewrite the learner entry README**

The README contains the real schedule, a “before class / in class / after class” path, commands for every solution and exercise, links to all block lab guides, the Scorebook Lab, Task Manager assignment, and Bonus. It explicitly says the current GitHub repository contains teacher answers.

- [ ] **Step 3: Write the 09:30–16:00 runbook**

The runbook reproduces every time row from the design spec and, for each block, states: objective, instructor action, learner action, observable checkpoint, common delay, and exact cuttable content. The two short breaks are 10:35–10:45 and 14:50–15:00; lunch is 12:00–13:00.

It also includes this auditable hands-on budget, counting only minutes in which learners operate the toolchain, edit code, run tests, or complete quizzes:

| Time box | Learner hands-on minutes |
|---|---:|
| Opening and Entry Quiz | 10 |
| Block 1 | 25 |
| Block 2 | 40 |
| Block 3 | 35 |
| Block 4 | 25 |
| Scorebook integrated lab | 40 |
| Homework kickoff and Exit Quiz | 10 |
| **Total** | **185 / 310 (59.7%)** |

- [ ] **Step 4: Write demo notes and assessments**

Entry Quiz has 8 questions measuring Java-to-Go starting assumptions. Exit Quiz has 10 behavior-oriented questions covering value semantics, slice sharing, map lookup, rune vs byte, receiver choice, closures, defer, and reading a test failure. `answer_key.md` contains answers plus one-sentence teaching diagnoses.

- [ ] **Step 5: Verify navigation and commands**

Run:

```bash
make module01-lab-01
make module01-lab-02
make module01-lab-03
make module01-lab-04
make module01-integrated-lab
make module01-homework-solution
```

Expected: all commands PASS.

- [ ] **Step 6: Commit course navigation and operations**

```bash
git add Makefile module01_basics/README.md module01_basics/instructor module01_basics/assessments
git commit -m "docs(course): add module 01 delivery system"
```

---

### Task 8: Migrate Bonus Material, Condense the Standard, and Remove Duplicate Entrypoints

**Files:**
- Move: `module01_basics/08_data_structures` → `module01_basics/bonus/data_structures`
- Move: `module01_basics/10_generics_intro` → `module01_basics/bonus/generics`
- Move selected dark-corner files into `module01_basics/bonus/dark_corners`
- Move advanced pattern files into `module01_basics/bonus/function_patterns`
- Create: `module01_basics/bonus/README.md`
- Create: `docs/course-module-standard.md`
- Modify: `Golang_Backend_Training_Syllabus.md`
- Rewrite: `docs/module01_basics_lesson_plan.md`
- Modify Module 01 portion: `docs/module01_02_lecture_cheatsheet.md`

**Interfaces:**
- Produces a single learner entry (`module01_basics/README.md`) and a single instructor entry (`module01_basics/instructor/RUNBOOK.md`).
- Produces the reusable Module 02–06 standard in `docs/course-module-standard.md`.

- [ ] **Step 1: Move Bonus directories and files with history**

```bash
mkdir -p module01_basics/bonus/dark_corners/{range,map,string} module01_basics/bonus/function_patterns
git mv module01_basics/08_data_structures module01_basics/bonus/data_structures
git mv module01_basics/10_generics_intro module01_basics/bonus/generics
git mv module01_basics/blocks/01_go_basics/demo/03_control_funcs/range_dark_corner.go module01_basics/bonus/dark_corners/range/main.go
git mv module01_basics/blocks/02_collections/demo/05_maps_strings/map_dark_corner.go module01_basics/bonus/dark_corners/map/main.go
git mv module01_basics/blocks/02_collections/demo/05_maps_strings/string_dark_corner.go module01_basics/bonus/dark_corners/string/main.go
git mv module01_basics/blocks/02_collections/demo/05_maps_strings/string_bench_test.go module01_basics/bonus/dark_corners/string/string_bench_test.go
git mv module01_basics/blocks/04_functions_testing/demo/09_advanced_functions/configuration_patterns.md module01_basics/bonus/function_patterns/configuration_patterns.md
git mv module01_basics/blocks/04_functions_testing/demo/09_advanced_functions/curry_best_practice_test.go module01_basics/bonus/function_patterns/curry_best_practice_test.go
git mv module01_basics/blocks/04_functions_testing/demo/09_advanced_functions/patterns_comparison_test.go module01_basics/bonus/function_patterns/patterns_comparison_test.go
git mv module01_basics/blocks/04_functions_testing/demo/09_advanced_functions/main.go module01_basics/bonus/function_patterns/main.go
```

Create a compact `module01_basics/blocks/04_functions_testing/demo/09_advanced_functions/main.go` containing only function values, one closure, and defer order:

```go
package main

import "fmt"

func atLeast(min int) func(int) bool {
	return func(value int) bool {
		return value >= min
	}
}

func main() {
	passed := atLeast(60)
	fmt.Println("75 passed:", passed(75))

	defer fmt.Println("end")
	fmt.Println("start")
}
```

- [ ] **Step 2: Make every Bonus directory independently runnable**

Keep the three dark-corner programs in their explicit `dark_corners/{range,map,string}` directories. In each moved `main.go`, rename the single exported entrypoint (`RunRangeDarkCornerDemo`, `RunMapDarkCornerDemo`, or `RunStringDarkCornerDemo`) to `main`, so both `go run ./module01_basics/bonus/dark_corners/<name>` and `go test ./module01_basics/bonus/...` work. Ensure all files under `function_patterns` use `package main`; no test package rename should be needed because the original directory already passes `go test` before the move.

Run:

```bash
go run ./module01_basics/bonus/dark_corners/range
go run ./module01_basics/bonus/dark_corners/map
go run ./module01_basics/bonus/dark_corners/string
go test ./module01_basics/bonus/...
```

Expected: all three programs run and all Bonus packages PASS.

- [ ] **Step 3: Write the condensed course-module standard**

`docs/course-module-standard.md` must define:

```markdown
# 六周课程 Module 标准
## 标准目录
## 时间盒与讲练比例
## Block README 模板
## Demo、Lab、Homework 职责
## Starter、公开测试与 Solution 分离
## 本地验收与 Gitee 手动 CI
## 教师仓库到学员仓库的披露流程
## Module 完成定义
```

Copy exact constraints from the design: one learner entry, one instructor entry, ≥50% hands-on time, one canonical grade script, standard-library-only Week 1 homework, and explicit out-of-scope content.

- [ ] **Step 4: Align existing course documents**

Update the Module 01 syllabus from “8 lessons” to “one-day four-block workshop + one-week homework.” Replace `docs/module01_basics_lesson_plan.md` with a short compatibility page linking to `instructor/RUNBOOK.md`. Replace only the Module 01 portion of the combined cheatsheet with a link to `DEMO_NOTES.md`; preserve its Module 02 content.

- [ ] **Step 5: Run the complete Module 01 verification**

Run:

```bash
make module01-verify
git diff --check
find module01_basics -maxdepth 1 -type d | sort
```

Expected:

- `make module01-verify` exits 0.
- `git diff --check` prints nothing.
- Top-level Module 01 directories are only `assessments`, `blocks`, `bonus`, `homework`, `instructor`, and `integrated_lab` plus the module root.

- [ ] **Step 6: Confirm the intentionally failing student paths**

Run each starter exercise test and the student homework grader individually. Expected: all fail on behavioral assertions, never compilation, missing dependency, or permission errors.

- [ ] **Step 7: Commit the migration and standard**

```bash
git add -A module01_basics Golang_Backend_Training_Syllabus.md Makefile
git add -f docs/course-module-standard.md docs/module01_basics_lesson_plan.md docs/module01_02_lecture_cheatsheet.md
git commit -m "refactor(course): standardize module 01 delivery"
```

---

## Final Verification

- [ ] Run `make module01-verify`; expect exit 0.
- [ ] Run `go test ./module01_basics/...`; expect exit 0.
- [ ] Run `go vet ./module01_basics/...`; expect exit 0.
- [ ] Run `git diff --check`; expect no output.
- [ ] Run `git status --short`; inspect that only intended plan/execution changes remain.
- [ ] Open every Markdown link from `module01_basics/README.md` and `module01_basics/instructor/RUNBOOK.md` with an automated link/path check or an `rg`-driven file existence script.
- [ ] From `student_pack`, run `make grade`; expect a behavioral test failure.
- [ ] From `teacher/solution`, run `../../student_pack/scripts/grade.sh`; expect `Task Manager 作业验收通过`.
- [ ] Review `.workflow/GradePipeline.yml` against the current official Gitee task orchestration and Go build plugin documentation before exporting the student pack.
