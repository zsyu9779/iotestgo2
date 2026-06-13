# Module 04: GORM 数据库操作 - 教师备课教案

**适用对象**: 已理解 HTTP API 分层，并具备基础 SQL 概念的本科生
**总课时**: 预计 8 课时
**教学目标**: 让学生掌握 Go 后端中数据库连接、模型设计、CRUD、关联查询、迁移、事务和测试的基本实践。

> 教学简化: 示例中的硬编码 DSN、演示数据库和自动迁移只服务课堂说明。正式项目应优先使用 `.env.example` 对应的环境变量配置，并由迁移工具管理结构变更。

## 每节课固定结构

1. 一句话定位
2. 运行一个最小 Demo
3. 解释核心机制
4. 修改一个小功能
5. 指出一个生产边界

## 第 1 课: GORM 初始化与连接

**源码路径**: `module04_gorm/01_setup`
**演示命令**:

```bash
docker compose -f module06_gozero/project_ecommerce/docker-compose.yml up -d mysql
MYSQL_DSN='root:root123@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local' go run ./module04_gorm/01_setup
```

**讲解重点**:
- DSN、Driver、连接池的含义
- `gorm.Open` 和底层 `database/sql` 的关系
- 为什么连接失败通常不是 ORM 问题

**练习**:
- 对比 `.env.example` 中的课堂默认 DSN，说明为什么演示命令要匹配 Docker Compose 的 MySQL 密码和数据库名。

**生产边界**:
- 不要在代码中写死数据库密码；课堂示例之后应迁移到 `.env` 或部署平台密钥。

## 第 2 课: 模型与关联

**源码路径**: `module04_gorm/02_models_relations`
**演示命令**:

```bash
go run ./module04_gorm/02_models_relations
```

**讲解重点**:
- Struct tag 和数据库字段映射
- 一对一、一对多、多对多关系
- 模型设计与业务概念的一致性

**练习**:
- 给用户增加 `Profile` 字段，并建立一对一关联。

**生产边界**:
- 真实项目要避免模型无限膨胀，应区分持久化模型、请求 DTO 和响应 View。

## 第 3 课: CRUD 基础

**源码路径**: `module04_gorm/03_crud`
**演示命令**:

```bash
go run ./module04_gorm/03_crud
```

**讲解重点**:
- Create、First、Find、Updates、Delete
- 条件查询和错误处理
- 零值更新陷阱

**练习**:
- 增加按邮箱查询用户的函数，并处理记录不存在。

**生产边界**:
- 真实写接口要考虑幂等性、唯一约束、输入校验和审计字段。

## 第 4 课: 高级查询与预加载

**源码路径**: `module04_gorm/04_queries_preload`
**演示命令**:

```bash
go run ./module04_gorm/04_queries_preload
```

**讲解重点**:
- `Preload` 解决关联读取
- N+1 查询问题
- Hooks 的触发时机和风险

**练习**:
- 为文章列表增加作者信息预加载，并观察 SQL 数量变化。

**生产边界**:
- 预加载不是越多越好，真实接口要结合分页、字段裁剪和慢查询分析。

## 第 5 课: 数据库迁移

**源码路径**: `module04_gorm/05_migrations`
**演示命令**:

```bash
go run ./module04_gorm/05_migrations
```

**讲解重点**:
- `AutoMigrate` 能做什么、不能做什么
- 表结构变更的风险
- 迁移脚本和版本管理思想

**练习**:
- 给模型增加一个可空字段，观察迁移后的表结构。

**生产边界**:
- 生产库不要依赖随应用启动自动迁移，应使用可审计、可回滚的迁移流程。

## 第 6 课: 事务

**源码路径**: `module04_gorm/06_transactions`
**演示命令**:

```bash
go run ./module04_gorm/06_transactions
```

**讲解重点**:
- 事务边界和一致性
- rollback、commit、savepoint
- 业务服务层为什么适合组织事务

**练习**:
- 实现“创建订单并扣库存”的事务伪代码。

**生产边界**:
- 跨服务事务不能简单套数据库事务，要引出最终一致性和消息补偿。

## 第 7 课: 数据库测试

**源码路径**: `module04_gorm/07_testing_mysql`
**演示命令**:

```bash
go test -v ./module04_gorm/07_testing_mysql
```

**讲解重点**:
- 集成测试和 sqlmock 的差异
- 测试数据准备与清理
- 为什么数据库测试要稳定、可重复

**练习**:
- 为一次查询增加 sqlmock 断言，验证 SQL 和参数。

**生产边界**:
- CI 中应区分快速单元测试和依赖容器的集成测试。

## 第 8 课: 博客 API 综合项目

**源码路径**: `module04_gorm/project_blog_api`
**演示命令**:

```bash
make run-blog-api
```

**讲解重点**:
- handler、service、repository、model 的协作
- 创建文章、评论和查询列表的业务链路
- GORM 在 Web 项目中的位置

**练习**:
- 为文章列表增加分页参数，并在 service 层做默认值处理。

**生产边界**:
- 博客项目是分层样例，不包含完整认证、权限、缓存和迁移流水线。
