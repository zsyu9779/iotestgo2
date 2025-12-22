# Module 04: GORM 数据库操作

本模块包含 Go 语言中使用 GORM 进行数据库操作的全部内容，从基础连接到完整的项目实践。

## 目录结构

### 01_setup/
- **main.go**: GORM 初始化和连接
- 学习内容：数据库连接、GORM 配置、连接池设置

### 02_models_relations/
- **main.go**: 模型定义和关系
- 学习内容：结构体标签、一对一、一对多、多对多关系

### 03_crud/
- **main.go**: 增删改查操作
- 学习内容：Create、Read、Update、Delete 操作、条件查询

### 04_queries_preload/
- **main.go**: 高级查询和预加载
- 学习内容：复杂查询、预加载关联数据、查询构建器

### 05_migrations/
- **main.go**: 数据库迁移
- 学习内容：自动迁移、手动迁移、版本控制

### 06_transactions/
- **main.go**: 事务处理
- 学习内容：事务管理、回滚、保存点、事务隔离

### 07_testing_mysql/
- **main_test.go**: 数据库测试
- 学习内容：测试数据库设置、测试数据准备、集成测试

### project_blog_api/
完整的博客 API 项目，包含：
- **internal/handler/post_handler.go**: 文章处理器
- **internal/model/post.go**: 文章模型
- **internal/repository/post_repo.go**: 文章数据存储
- **internal/service/post_service.go**: 文章业务逻辑
- **main.go**: 主程序

## 学习目标

1. 掌握 GORM 的基本配置和连接管理
2. 熟练定义数据模型和关系
3. 掌握完整的 CRUD 操作
4. 能够编写复杂的查询语句
5. 理解和使用数据库迁移
6. 掌握事务处理和数据一致性
7. 编写数据库集成测试
8. 完成一个完整的博客 API 项目

## 运行方式

每个目录下的程序都可以通过以下命令运行：
```bash
cd 目录名
go run main.go
```

对于测试：
```bash
cd 07_testing_mysql/
go test -v
```

运行博客项目：
```bash
cd project_blog_api/
go run main.go
```

## 数据库配置

大多数示例需要配置数据库连接，通常通过环境变量：
```bash
export DB_DSN="user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
```

或者修改代码中的连接字符串。

## 依赖安装

需要安装 GORM 和数据库驱动：
```bash
go mod tidy
```

对于 MySQL：
```bash
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
```