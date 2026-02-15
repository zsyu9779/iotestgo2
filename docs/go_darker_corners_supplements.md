# Go“黑暗角落”补充教学点（基于 iotestgo2 + TonyBai 译文 Part2）

目标：在不改动现有章节结构与代码的前提下，整理一份可插入到 Module01/02 的“高价值陷阱点”清单，并给出对应可直接复制运行的示例代码/练习代码，供后续梳理与分发。

参考来源：
- iotestgo2 现有教案：`docs/module01_basics_lesson_plan.md`、`docs/module02_advanced_lesson_plan.md`
- 博客：<https://tonybai.com/2021/03/29/darker-corners-of-go-part2/>

## 放置建议（与现有教案对齐）

| 主题 | 建议插入位置 | 课堂形式 | 一句话结论 |
|---|---|---|---|
| string 的本质、不可变、拼接性能 | Module01 第5课 Map与字符串 | 演示 + Benchmark | 拼接请用 Builder/bytes.Buffer，别在循环里 `+=` |
| UTF-8、len/index vs range/rune | Module01 第5课 | 演示 + 小练习 | `len` 是字节数；`s[i]` 是字节；`range` 是 rune |
| map 迭代顺序、nil map、comma-ok | Module01 第5课 | 演示 + 问答 | map 顺序未定义；nil map 可读不可写；存在性用 comma-ok |
| map 值不可寻址（不能改 `m[k].Field`） | Module01 第5/7/8课（首次出现 struct+map） | 演示 + 练习 | 取出-修改-写回，或 map 存指针 |
| range 变量复用（取地址/闭包/协程） | Module01 第3课（for）+ Module02 第3课（goroutine） | 演示 + 修复模板 | 循环变量被复用；要复制一份或作为参数传入 |
| defer 细节（参数求值、作用域、命名返回） | Module01 第9课 + Module02 第2课 | 演示 + 问答 | defer 参数立即求值；defer 只随函数返回触发 |
| recover 只在 defer 生效；goroutine panic 会崩全局 | Module02 第2/3课 | 演示 + 修复模板 | 需要对 goroutine 做保护性 recover |
| 接口 typed-nil（接口不等于 nil） | Module02 第1课（interfaces） | 演示 + 规约 | 该返回 nil 就直接 return nil（接口） |
| 相等性：slice/map 不能 `==`；DeepEqual 慢且有坑 | Module02 第7课（testing） | 演示 + Benchmark | 关键路径手写 Equals；byte slice 用 bytes.Equal |
| 并发 map：读写会 fatal；sync.Map 仅特定场景 | Module02 第6课（并发安全） | 演示 | 常规用 `map+RWMutex`；sync.Map 有明确适用面 |
| log.Fatal/log.Panic 的副作用；time.LoadLocation 性能坑 | Module02 工程化补充（errors/os/time） | 演示 + Benchmark | Fatal 会 os.Exit，不跑 defer；LoadLocation 要缓存 |

下面给出每个主题的“教学点 + 可运行代码”。后续可以把它们拆进现有章节末尾的「黑暗角落」小节。

---

## 1) String：本质、不可变、拼接性能

### 教学点
- string 不是 `nil`：零值是空串 `""`。
- string 是“指针 + 长度”的值语义视图（底层字节不可随意改）。
- `len(s)` 是字节数，不是字符数。
- `s[i]` 取到的是第 i 个字节；`for range` 以 rune（Unicode code point）迭代。
- 循环中大量 `+=` 会反复分配新字符串；优先 `strings.Builder`（或 `bytes.Buffer`）。

### 演示代码：string 零值与不可变

```go
package main

import "fmt"

func main() {
	var s string
	fmt.Println(s == "")
	fmt.Println(len(s))
	// s = nil
}
```

### 演示代码：len / index vs range（UTF-8）

```go
package main

import "fmt"

func main() {
	s := "touché你好"

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

	for _, r := range s {
		fmt.Print(string(r))
	}
	fmt.Println()
}
```

### 基准测试代码：`+=` vs `strings.Builder`

文件建议：`string_concat_test.go`

```go
package main

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkStringPlusEqual(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strconv.Itoa(i)
	}
	_ = s
}

func BenchmarkStringsBuilder(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	_ = sb.String()
}
```

### 练习建议
- 给定字符串 `"你好Go"`：分别打印字节长度、rune 个数、以及每个 rune 的 `U+` 编码。
- 让学员把循环 `+=` 改成 Builder，并用基准测试比较差异。

---

## 2) Map：顺序、nil、不可寻址、并发

### 教学点
- map 迭代顺序未定义；不要依赖遍历顺序。
- nil map：`len` 和读取 OK，写入会 panic。
- 查键是否存在：用 `val, ok := m[k]`，不要拿零值猜。
- map 作为参数传递时，函数内改元素会影响外部；但函数内 `m = make(...)` 不会“替换”外部变量。
- map 的值不可寻址：不能 `&m[k]`，也不能直接改 `m[k].Field`（当值是 struct 时）。
- 并发读写 map 会直接 fatal；常规方案 `map + (RW)Mutex`；`sync.Map` 仅特定读多写少/键不相交场景。

### 演示代码：迭代顺序“看起来随机”

```go
package main

import "fmt"

func main() {
	m := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	for t := 0; t < 5; t++ {
		for k := range m {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	m[6] = 6
	m[7] = 7
	m[8] = 8

	for t := 0; t < 5; t++ {
		for k := range m {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}
```

### 演示代码：nil map 读写行为

```go
package main

import "fmt"

func main() {
	var m map[int]int
	fmt.Println(len(m))
	fmt.Println(m[10])
	// m[10] = 1

	m = make(map[int]int)
	m[10] = 11
	fmt.Println(m[10])
}
```

### 演示代码：map 传参的“替换”误解

```go
package main

import "fmt"

func fill(m map[int]int) {
	m[1] = 1
}

func replace(m map[int]int) {
	m = make(map[int]int)
	m[2] = 2
}

func main() {
	m := make(map[int]int)
	fill(m)
	replace(m)
	fmt.Println(m[1])
	fmt.Println(m[2])
}
```

### 演示代码：map 值不可寻址与正确修改方式

```go
package main

import "fmt"

type Item struct {
	Value string
}

func main() {
	m := map[int]Item{1: {Value: "one"}}

	tmp := m[1]
	tmp.Value = "two"
	m[1] = tmp

	fmt.Println(m[1].Value)
}
```

### 演示代码：并发读写 map 会 fatal（课堂只演示，避免学员误以为是偶现）

```go
package main

import (
	"math/rand"
	"time"
)

func rw(m map[int]int) {
	for i := 0; i < 1000; i++ {
		k := rand.Int()
		m[k] = m[k] + 1
		_ = m[k]
	}
}

func main() {
	m := make(map[int]int)
	for i := 0; i < 8; i++ {
		go rw(m)
	}
	time.Sleep(2 * time.Second)
}
```

### 修复模板：map + Mutex

```go
package main

import (
	"math/rand"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (s *SafeMap) RW() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := 0; i < 1000; i++ {
		k := rand.Int()
		s.m[k] = s.m[k] + 1
		_ = s.m[k]
	}
}

func main() {
	s := &SafeMap{m: make(map[int]int)}
	for i := 0; i < 8; i++ {
		go s.RW()
	}
	time.Sleep(2 * time.Second)
}
```

---

## 3) range/for：单变量是索引 + 迭代变量复用

### 教学点
- `for v := range slice`：v 是索引，不是元素值。
- 循环迭代变量会被复用：取地址/闭包捕获/启动 goroutine 都可能拿到同一个变量地址。
- 修复模板：在循环体内 `v := v` 复制一份，或把变量作为参数传入。

### 演示代码：单变量 range 是索引

```go
package main

import "fmt"

func main() {
	s := []string{"one", "two", "three"}

	for v := range s {
		fmt.Println(v)
	}

	for _, v := range s {
		fmt.Println(v)
	}
}
```

### 演示代码：取地址导致全是同一个值

```go
package main

import "fmt"

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println(*out[0], *out[1], *out[2])
	fmt.Println(out[0], out[1], out[2])
}
```

### 修复模板：循环体内复制一份

```go
package main

import "fmt"

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		i := i
		out = append(out, &i)
	}
	fmt.Println(*out[0], *out[1], *out[2])
	fmt.Println(out[0], out[1], out[2])
}
```

### 演示代码：goroutine 捕获循环变量

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Print(i)
		}()
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println()

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println()
}
```

---

## 4) defer / panic / recover：参数求值、作用域、goroutine 崩溃

### 教学点
- defer 的参数在 defer 语句处求值，而不是函数返回时求值。
- defer 触发点是“函数返回”，不是“离开代码块”。
- defer 按 LIFO 执行。
- recover 只在 defer 中生效。
- goroutine 内 panic 若不 recover，会把整个进程带崩。

### 演示代码：defer 参数求值

```go
package main

import "fmt"

func main() {
	s := "defer"
	defer fmt.Println(s)
	s = "original"
	fmt.Println(s)
}
```

### 演示代码：defer 不随代码块结束触发

```go
package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			defer func(i int) {
				fmt.Println(i)
			}(i)
		}
	}
	fmt.Println("exiting")
}
```

### 演示代码：recover 只能在 defer 里

```go
package main

import "fmt"

func panicky() {
	panic("boom")
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered", r)
		}
	}()

	panicky()
	fmt.Println("unreachable")
}
```

### 修复模板：给 goroutine 加保护（生产常用）

```go
package main

import "fmt"

func GoSafe(fn func()) {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("panic", r)
			}
		}()
		fn()
	}()
}

func main() {
	GoSafe(func() {
		panic("bad")
	})
}
```

---

## 5) Interface：typed-nil 与 nil 判断

### 教学点
- 接口值包含：静态类型、动态类型、值；只有“动态类型和值都为 nil”时接口才等于 nil。
- 规约：返回接口时，若没有值就 `return nil`，不要返回“带类型的 nil 指针”。
- 类型断言单返回会 panic；两返回更安全。

### 演示代码：typed-nil

```go
package main

import "fmt"

type Greeter interface {
	Greet()
}

type G struct{}

func (g *G) Greet() { fmt.Println("hi") }

func main() {
	var gi Greeter
	fmt.Println(gi == nil)

	var p *G
	gi = p
	fmt.Println(gi == nil)
}
```

### 演示代码：返回接口时的推荐写法

```go
package main

import "fmt"

type Greeter interface{ Greet() }

type G struct{}

func (g *G) Greet() { fmt.Println("hi") }

func NewGreeter(ok bool) Greeter {
	if !ok {
		return nil
	}
	return &G{}
}

func main() {
	fmt.Println(NewGreeter(false) == nil)
	fmt.Println(NewGreeter(true) == nil)
}
```

### 演示代码：类型断言单返回 vs 双返回

```go
package main

import "fmt"

type A interface{ M() }

type T struct{}

func (t *T) M() {}

func main() {
	var a A = &T{}
	_, ok := a.(string)
	fmt.Println(ok)
	// _ = a.(string)
}
```

---

## 6) 相等性：`==` 的边界、DeepEqual 的坑与性能

### 教学点
- slice/map 不能用 `==` 比较（只能和 nil 比）。
- struct 若包含不可比较字段（slice/map/func），整体也不可 `==`。
- `reflect.DeepEqual` 通用但慢，且语义有坑：nil slice 与 empty slice；未导出字段；NaN；func。
- `[]byte` 建议用 `bytes.Equal`。
- 关键路径建议手写 Equals，并用 Benchmark 验证。

### 演示代码：slice 不能 `==`

```go
package main

func main() {
	a := []int{1}
	b := []int{1}
	_, _ = a, b
	// _ = (a == b)
}
```

### 演示代码：bytes.Equal vs DeepEqual 的语义差异（nil vs empty）

```go
package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var a []byte
	b := []byte{}
	fmt.Println(bytes.Equal(a, b))
	fmt.Println(reflect.DeepEqual(a, b))
}
```

### 基准测试：手写 Equals vs DeepEqual

```go
package main

import (
	"reflect"
	"testing"
)

type S struct {
	A int
	B string
	C []int
}

func (s S) Equals(o S) bool {
	if s.A != o.A || s.B != o.B || len(s.C) != len(o.C) {
		return false
	}
	for i := 0; i < len(s.C); i++ {
		if s.C[i] != o.C[i] {
			return false
		}
	}
	return true
}

var s1 = S{A: 1, B: "x", C: []int{1, 2, 3}}
var s2 = S{A: 1, B: "x", C: []int{1, 2, 3}}

func BenchmarkManualEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s1.Equals(s2)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(s1, s2)
	}
}
```

---

## 7) 值/指针：不要迷信“指针更快”（逃逸与 GC）

### 教学点
- Go 参数永远是值传递；传指针只是“复制指针值”。
- 指针可能导致逃逸到堆，引入分配与 GC；值返回可能完全走栈更快。
- 结论以 Benchmark + pprof 为准；工程里通常语义优先。

### 基准测试：byValue vs byReference

```go
package main

import "testing"

type MyStruct struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byValue() MyStruct {
	return MyStruct{a: 1, b: 1, c: 1, d: "foo", e: "bar", f: "baz", g: 1, h: 1, i: 1}
}

func byReference() *MyStruct {
	return &MyStruct{a: 1, b: 1, c: 1, d: "foo", e: "bar", f: "baz", g: 1, h: 1, i: 1}
}

func BenchmarkByValue(b *testing.B) {
	var s MyStruct
	for i := 0; i < b.N; i++ {
		s = byValue()
	}
	_ = s
}

func BenchmarkByReference(b *testing.B) {
	var s *MyStruct
	for i := 0; i < b.N; i++ {
		s = byReference()
	}
	_ = s
}
```

---

## 8) log 与 time 的工程坑：Fatal/Panic、LoadLocation

### 教学点
- `log.Fatal(...)` = 打印 + `os.Exit(1)`：不会执行 defer（资源释放、Unlock、Flush 都会丢）。
- `log.Panic(...)` = 打印 + panic：会走 defer，但会把栈打出来。
- `time.LoadLocation` 每次调用可能读文件：要缓存 `*time.Location`。

### 演示代码：Fatal 不跑 defer

```go
package main

import (
	"log"
)

func main() {
	defer log.Println("defer")
	log.Fatal("fatal")
}
```

### 基准测试：LoadLocation 缓存

```go
package main

import (
	"testing"
	"time"
)

func BenchmarkLoadLocationEveryTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		_ = time.Now().In(loc)
	}
}

func BenchmarkLoadLocationCached(b *testing.B) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for i := 0; i < b.N; i++ {
		_ = time.Now().In(loc)
	}
}
```

---

## 附录 A：课堂组织方式（不改代码的前提下）

- 每课末尾加一个 5 分钟「黑暗角落」：只讲 1–2 个点，配 1 个“必错题”。
- 每个点给三件套：
  - 错误写法演示（可直接复制运行）
  - 正确写法修复模板
  - 一句话规约让学员带走

## 附录 B：后续如果要落到 iotestgo2 的文档里

- Module01 第5课追加：String/UTF-8/Builder、map nil/顺序/不可寻址。
- Module01 第3课追加：range 单变量是索引、迭代变量复用。
- Module01 第9课追加：defer 参数求值与作用域。
- Module02 第1课加强：typed-nil 的规约与常见返回接口场景。
- Module02 第2/3课加强：goroutine panic 的全局影响与 GoSafe 模板。
- Module02 第6课加强：map 并发与 sync.Map 适用面。
- Module02 第7课加强：DeepEqual vs 手写 Equals/bytes.Equal 的 benchmark。
- Module02 工程化彩蛋：log.Fatal、time.LoadLocation。
