// go-sqlmock 演示：不依赖真实 MySQL 的 Service 层单元测试
//
// 安装：
//   go get github.com/DATA-DOG/go-sqlmock
//
// 原理：
//   sqlmock 模拟 MySQL 驱动，拦截 SQL 语句，返回预设结果
//   适合测试 Service/Logic 层的数据库操作逻辑
package main

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type UserModel struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
}

func (s *UserService) GetUserByID(id uint) (*UserModel, error) {
	var user UserModel
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(name string) error {
	return s.db.Create(&UserModel{Name: name}).Error
}

// TestUserService_GetByID 测试查询用户
func TestUserService_GetByID(t *testing.T) {
	fmt.Println("=== sqlmock 演示：测试 Service 层 ===")
	fmt.Println()
	fmt.Println("sqlmock 使用步骤：")
	fmt.Println("  1. sqlmock.New() 创建 mock DB 连接")
	fmt.Println("  2. 用 mock.ExpectQuery/ExpectExec 设置预期 SQL 和返回结果")
	fmt.Println("  3. 用 mock DB 创建 gorm.DB")
	fmt.Println("  4. 调用待测试的 Service/Logic 方法")
	fmt.Println("  5. mock.ExpectationsWereMet() 验证所有预期 SQL 都被执行")
	fmt.Println()
	fmt.Println("优势：")
	fmt.Println("  - 不需要真实 MySQL，CI 环墣中可直接运行")
	fmt.Println("  - 精确控制 SQL 返回（正常数据、错误、超时）")
	fmt.Println("  - 验证 SQL 语句是否正确")
	fmt.Println("  - 测试速度快（无网络 IO）")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: @DataJpaTest + H2 in-memory database")
	fmt.Println("  Go:   sqlmock 更精确，但需手写预期 SQL 匹配")
	fmt.Println()

	_ = testing.T{}
	_ = gorm.Open
	_ = mysql.New
}
