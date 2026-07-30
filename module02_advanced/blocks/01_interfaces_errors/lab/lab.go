package lab

import (
	"fmt"
	"math"
	"strconv"
)

// Shape 描述可以计算面积的图形。
type Shape interface {
	Area() float64
}

// Rectangle 表示矩形。
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 返回矩形面积。
func (r Rectangle) Area() float64 { return r.Width * r.Height }

// Circle 表示圆形。
type Circle struct {
	Radius float64
}

// Area 返回圆形面积。
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// TotalArea 汇总所有图形面积。
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// ValidationError 表示输入值未通过业务校验。
type ValidationError struct {
	Field string
	Msg   string
}

// Error 返回校验错误文本。
func (e *ValidationError) Error() string { return fmt.Sprintf("%s: %s", e.Field, e.Msg) }

// ParsePort 将字符串解析为合法 TCP 端口。
func ParsePort(value string) (int, error) {
	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parse port %q: %w", value, &ValidationError{Field: "port", Msg: "必须是整数"})
	}
	if port < 1 || port > 65535 {
		return 0, &ValidationError{Field: "port", Msg: "必须在 1 到 65535 之间"}
	}
	return port, nil
}
