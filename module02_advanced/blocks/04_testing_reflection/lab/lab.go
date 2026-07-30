package lab

import (
	"fmt"
	"reflect"
)

// User 是反射练习使用的业务结构体。
type User struct {
	ID   int
	Name string
}

// ReadFieldName 读取结构体字段并返回字符串值。
func ReadFieldName(value User, fieldName string) (string, error) {
	field := reflect.ValueOf(value).FieldByName(fieldName)
	if !field.IsValid() {
		return "", fmt.Errorf("字段 %q 不存在", fieldName)
	}
	if field.Kind() != reflect.String {
		return "", fmt.Errorf("字段 %q 不是字符串", fieldName)
	}
	return field.String(), nil
}
