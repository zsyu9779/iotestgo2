// 扩展 error wrap 演示：fmt.Errorf("%w") + errors.Is/As 深入
//
// Go 1.13+ error wrapping 机制
package main

import (
	"errors"
	"fmt"
)

// 基础错误类型
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Msg)
}

// 错误包装链
var ErrDB = errors.New("database error")
var ErrNotFound = fmt.Errorf("not found: %w", ErrDB)

func validateUsername(name string) error {
	if len(name) < 3 {
		return &ValidationError{Field: "username", Msg: "too short"}
	}
	return nil
}

func findUser(name string) error {
	if err := validateUsername(name); err != nil {
		// 包装错误，附加上下文
		return fmt.Errorf("findUser(%s): %w", name, err)
	}
	// 模拟数据库查询失败
	return fmt.Errorf("findUser(%s): %w", name, ErrNotFound)
}

func main() {
	fmt.Println("=== Error Wrapping 深入 ===")
	fmt.Println()

	// --- 1. %w 包装 + errors.Is ---
	fmt.Println("--- 1. errors.Is ---")
	err := findUser("ab")
	fmt.Printf("  Error: %v\n", err)

	// errors.Is 沿错误链查找
	fmt.Printf("  Is ErrDB? %v\n", errors.Is(err, ErrDB))
	fmt.Printf("  Is ErrNotFound? %v\n", errors.Is(err, ErrNotFound))

	// --- 2. errors.As ---
	fmt.Println()
	fmt.Println("--- 2. errors.As ---")
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("  ValidationError found: field=%s, msg=%s\n", valErr.Field, valErr.Msg)
	} else {
		fmt.Println("  Not a ValidationError")
	}

	// --- 3. errors.Unwrap ---
	fmt.Println()
	fmt.Println("--- 3. errors.Unwrap ---")
	unwrapped := errors.Unwrap(err)
	fmt.Printf("  Unwrap 1: %v\n", unwrapped)
	unwrapped = errors.Unwrap(unwrapped)
	fmt.Printf("  Unwrap 2: %v\n", unwrapped)

	// --- 4. fmt.Errorf 多层包装 ---
	fmt.Println()
	fmt.Println("--- 4. 多层包装 ---")
	origErr := errors.New("original error")
	wrap1 := fmt.Errorf("wrap1: %w", origErr)
	wrap2 := fmt.Errorf("wrap2: %w", wrap1)
	wrap3 := fmt.Errorf("wrap3: %w", wrap2)

	fmt.Printf("  Is original? %v\n", errors.Is(wrap3, origErr))
	fmt.Printf("  错误链: %v\n", wrap3)

	fmt.Println()
	fmt.Println("--- 最佳实践 ---")
	fmt.Println("  1. 自定义错误类型 + errors.As 做类型判断")
	fmt.Println("  2. Sentinel Error (var ErrXxx) + errors.Is 做值判断")
	fmt.Println("  3. 每层包装用 %w 附加上下文，保留原始错误")
	fmt.Println("  4. 不要重复包装同一错误（避免信息冗余）")
	fmt.Println("  5. fmt.Errorf(\": %w\", err) 注意冒号和空格")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: throw new CustomException(msg, cause)")
	fmt.Println("  Go:   fmt.Errorf(\"context: %w\", cause)")
	fmt.Println("  Go 用多重返回值替代 try-catch，错误链更显式")
}
