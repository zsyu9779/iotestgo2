package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

// ValidationError 表示调用方可以修正的输入错误。
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// findUser 在每一层补充上下文，并用 %w 保留原始错误身份。
func findUser(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("validate user %q: %w", name, &ValidationError{
			Field: "username",
			Msg:   "must contain at least 3 characters",
		})
	}
	return fmt.Errorf("user %q: %w", name, ErrUserNotFound)
}

// recoverAtBoundary 只演示 defer 的这一项职责：在程序边界捕获 panic。
// 普通业务错误仍应通过 error 返回值处理。
func recoverAtBoundary(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}

func main() {
	notFoundErr := findUser("alice")
	fmt.Println(notFoundErr)
	fmt.Println("is not found:", errors.Is(notFoundErr, ErrUserNotFound))

	validationErr := findUser("ab")
	var validation *ValidationError
	if errors.As(validationErr, &validation) {
		fmt.Println("invalid field:", validation.Field)
	}

	if recovered := recoverAtBoundary(func() { panic("unexpected state") }); recovered != nil {
		fmt.Println("recovered at boundary:", recovered)
	}
}
