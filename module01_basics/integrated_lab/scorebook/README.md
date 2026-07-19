# 成绩册综合实验

请在 `starter/scorebook.go` 中独立完成一个内存成绩册，把前四个区块中的条件分支、Map、Struct、指针、方法和测试串成一次完整练习。不要导入任何区块的 `solution`；本实验的目标是把已经学过的概念迁移到一个新问题。

保持以下公开接口不变：

```go
type Student struct {
	ID    int
	Name  string
	Score int
}

func New() *Scorebook
func (s *Scorebook) Add(name string, score int) (Student, error)
func (s *Scorebook) Find(id int) (Student, error)
func (s *Scorebook) UpdateScore(id int, score int) error
func (s *Scorebook) Average() (float64, error)
func (s *Scorebook) CountByGrade() map[string]int
```

`Scorebook` 内部使用 `map[int]*Student` 保存学生，并用 `nextID int` 生成从 1 开始的连续 ID。`Add` 和 `Find` 都返回 `Student` 值副本；调用者修改返回值时，成绩册中的记录不能变化。

姓名先用 `strings.TrimSpace` 去除首尾空白，结果为空时返回 `ErrInvalidName`。成绩必须在 0–100（含两端）之间，否则返回 `ErrInvalidScore`；失败的 `Add` 不消耗 ID，失败的 `UpdateScore` 不改变原成绩。找不到 ID 时返回 `ErrStudentNotFound`，空成绩册求平均值时返回 `0, ErrEmptyScorebook`。测试使用 `errors.Is` 判断这些哨兵错误。

等级规则为 A：90–100、B：80–89、C：70–79、D：60–69、F：0–59。`CountByGrade` 只需为实际出现的等级累计数量。

## 时间盒：40 分钟

### 检查点 1（0–10 分钟）：模型与 Add

为 `Scorebook` 添加 `students map[int]*Student` 和 `nextID int`，让 `New` 初始化可写入的 Map，并实现 `Add` 的姓名/成绩校验、连续 ID 与值副本返回。先运行：

```bash
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter -run '^(TestScorebookWorkflow|TestAddRejectsInvalidFieldsWithoutConsumingID)$'
```

**停止规则：**到第 10 分钟停止扩展 `Add`。只有当两个学生得到 ID 1 和 2、姓名被清理、无效输入返回约定错误且不消耗 ID 时才进入下一检查点；否则先使用提示或请同伴定位当前第一个失败断言。

### 检查点 2（10–20 分钟）：Find 与 UpdateScore

实现按 ID 查找和更新。`Find` 返回解引用后的值副本；`UpdateScore` 修改 Map 中保存的同一个 `*Student`，并在写入前处理未知 ID 与无效成绩。运行：

```bash
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter -run '^(TestScorebookWorkflow|TestFindAndUpdateReturnStudentNotFound|TestUpdateScoreRejectsInvalidScoresWithoutMutation)$'
```

**停止规则：**到第 20 分钟停止处理查找和更新。只有当未知 ID 返回 `ErrStudentNotFound`、无效更新保留旧成绩，并且修改 `Add`/`Find` 的返回值不会污染内部记录时才继续；否则保留失败输出并请求一次针对该断言的帮助。

### 检查点 3（20–30 分钟）：Average

遍历 Map 中的学生指针累加成绩。空 Map 返回 `0, ErrEmptyScorebook`；非空时先把总分或人数转换为 `float64`，避免整数除法丢失小数。运行：

```bash
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter -run '^TestAverage'
```

**停止规则：**到第 30 分钟停止修改平均值逻辑。`80` 与 `81` 的平均值必须是 `80.5`，空成绩册必须能被 `errors.Is(err, ErrEmptyScorebook)` 识别；若未达到，先只检查空集合分支和类型转换，不提前实现等级统计。

### 检查点 4（30–40 分钟）：等级统计与全量测试

编写私有 `grade(score int) string`，按从高到低的边界返回 A、B、C、D 或 F。`CountByGrade` 创建计数 Map，遍历所有学生并对相应等级加一。先运行边界测试，再运行整个练习包：

```bash
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter -run '^TestCountByGradeCoversEveryBoundary$'
go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter
```

**停止规则：**第 40 分钟立即停止编码并保留测试输出。全量测试通过即完成课堂实验；若仍失败，记录第一个失败测试、实际值和期望值，作为课后继续定位的起点，不在课堂内增加新功能。

## 三级提示

1. **一级提示（方向）：**每次只解决当前检查点的第一个失败。想清楚哪些方法要读 Map、哪些方法要改 Map，以及错误分支必须发生在状态变化之前。
2. **二级提示（结构）：**`New` 令 `nextID` 从 1 开始；`Add` 校验成功后创建 `*Student`、按 ID 存入 Map、再递增 ID，最后返回 `*student`。`Find` 也通过 `return *student, nil` 形成副本。
3. **三级提示（关键表达式）：**平均值使用 `float64(total) / float64(len(s.students))`；等级判断按 `score >= 90`、`>= 80`、`>= 70`、`>= 60` 排列；计数语句可以写成 `counts[grade(student.Score)]++`。
