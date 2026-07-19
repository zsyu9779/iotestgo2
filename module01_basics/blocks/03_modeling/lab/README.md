# 学生模型实验

请在 `starter/student.go` 中完成一个维护自身不变量的学生模型。

`Student` 必须保持以下状态：

- `ID` 是正整数。
- `Name` 去除首尾空白后非空，模型中保存去除空白后的值。
- `Score` 位于 0–100，0 和 100 都是有效成绩。

保持以下公开接口不变：

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

校验失败时返回已有的哨兵错误：非正 ID 返回 `ErrInvalidID`，去除首尾空白后为空的姓名返回 `ErrInvalidName`，范围外的成绩返回 `ErrInvalidScore`。`Rename` 或 `UpdateScore` 失败时，不得改变原来的有效状态。

`Rename` 和 `UpdateScore` 要修改调用方持有的对象，因此使用指针接收者。`Snapshot` 只读取当前状态并返回一份值副本，因此使用值接收者并返回 `Student`。请准备结合测试中的“修改快照不影响原对象”现象，解释这两个选择。

从仓库根目录先运行测试，观察初始实现如何在第一个状态断言处失败：

```bash
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter
```

遵循 RED、GREEN、REFACTOR 的循环：一次关注一个失败，写最小实现使它通过，再整理重复的校验逻辑。完成后重复运行命令，直到全部通过。

## 三级提示

1. 一级：构造函数按 ID、姓名、成绩的顺序校验；姓名可以先交给 `strings.TrimSpace`。
2. 二级：修改方法先完成校验，再给接收者字段赋值，这样失败时原状态保持不变。
3. 三级：把姓名和成绩校验提取为未导出的辅助函数；`Snapshot` 直接返回值接收者 `s`。
