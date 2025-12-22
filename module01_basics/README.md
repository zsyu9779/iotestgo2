# Module 01: Go 语言基础

本模块包含 Go 语言的基础知识和核心概念，适合初学者入门学习。

## 目录结构

### 01_hello/
- **main.go**: 第一个 Go 程序，学习基本的程序结构和输出
- 学习内容：package main、import、fmt.Println()

### 02_vars_types/
- **main.go**: 变量声明和基本数据类型
- 学习内容：var、const、:= 短声明、int、string、bool、float64 等类型

### 03_control_funcs/
- **main.go**: 控制流和函数
- 学习内容：if/else、for 循环、switch、函数定义和调用

### 04_arrays_slices/
- **main.go**: 数组和切片
- 学习内容：数组声明、切片操作、append、len、cap

### 05_maps_strings/
- **main.go**: 映射和字符串操作
- 学习内容：map 声明和操作、字符串处理、rune

### 06_pointers/
- **main.go**: 指针
- 学习内容：指针声明、取地址(&)、解引用(*)、指针与函数

### 07_structs_methods/
- **main.go**: 结构体和方法
- 学习内容：结构体定义、方法接收器、值接收器 vs 指针接收器

### 08_data_structures/
- **main.go**: 数据结构实践
- 学习内容：自定义类型、复杂数据结构组合

### 09_advanced_functions/
- **main.go**: 高级函数特性
- 学习内容：函数变量、匿名函数、闭包、高阶函数、函数式编程、柯里化、延迟执行

### project_task_manager/
- **main.go**: 任务管理器项目
- **task_manager_test.go**: 单元测试
- 学习内容：综合应用前面所学知识，实现一个简单的任务管理系统

## 学习目标

1. 掌握 Go 语言的基本语法和结构
2. 理解变量、数据类型和内存管理
3. 熟练使用控制流和函数
4. 掌握数组、切片、映射等集合类型
5. 理解指针的概念和使用
6. 能够定义和使用结构体及方法
7. 掌握高级函数特性：函数变量、匿名函数、闭包、高阶函数
8. 理解函数式编程概念和模式
9. 完成一个综合性的小项目

## 运行方式

每个目录下的程序都可以通过以下命令运行：
```bash
cd 目录名
go run main.go
```

对于包含测试的项目：
```bash
cd project_task_manager
go test -v
```