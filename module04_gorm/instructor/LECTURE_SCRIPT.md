# Module 04：GORM 授课逐字讲稿

> 适用对象：已经完成 Go 基础、Gin 与基础 SQL，具备 Java/JDBC 或 JPA 阅读经验的学生  
> 当前版本：6 节核心课，每节 75 分钟；博客综合 Lab 作为课后作业启动  
> 使用方式：普通文字可直接照读；`【操作】` 是投屏操作；`【预期】` 是结果检查；`【提问】` 后停 3–5 秒；`【板书】` 写到白板；`【学员操作】` 是学生动手时间。

## 一、授课边界

本讲稿只使用当前工作区实际保留的课程内容：

1. 连接、配置与连接池。
2. 模型、关系、软删除与迁移。
3. CRUD 与零值更新。
4. 条件查询、JOIN、Preload 与 Hooks。
5. 事务、回滚与 SavePoint。
6. Raw SQL、Exec 与 ORM 选型。
7. 博客综合 Lab 作业启动。

课堂中的 MySQL、固定数据、`AutoMigrate` 和本地 DSN 都是教学方案。生产系统还必须单独设计版本化迁移、密钥管理、最小权限、TLS、监控、备份与恢复。

## 二、课前准备

### 1. 创建课堂数据库

```sql
CREATE DATABASE gorm_demo
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;
```

为了减少课堂准备步骤，Module 04 使用代码内固定的本机教学连接：

```text
root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local
```

无需导出环境变量。本机账号或密码不一致时，课前统一修改 `module04_gorm/internal/classroomdb/db.go` 中的 `DSN`。强调该固定连接仅用于本机教学，生产项目不得在代码中保存真实凭证。

### 2. 环境检查

```bash
go version
make module04-verify
make module04-env-check
```

【预期】

- 离线检查输出 `module04 verify: PASS`。
- 环境检查输出 MySQL 版本、数据库名、字符集和 `ddl=ok`。
- 当前账号至少具备课堂库的 `CREATE`、`ALTER`、`INDEX`、`SELECT`、`INSERT`、`UPDATE`、`DELETE` 权限。

### 3. 数据安全规则

所有课堂表以 `m04_` 开头。如果需要清理，只能逐个确认并删除这些课堂表。不要执行清空整个数据库、删除其他 schema 或递归删除工作区的命令。

### 4. 终端布局

- 终端 A：运行 Go Demo 和测试。
- 终端 B：连接 MySQL，执行只读检查语句。
- 编辑器：左侧打开当前 Demo，右侧打开对应 Starter 测试。

### 5. 每节固定节奏

| 时间 | 内容 |
|---|---|
| 0–15 分钟 | 概念、业务问题与 Java 对比 |
| 15–35 分钟 | 先预测，再运行可观察 Demo |
| 35–65 分钟 | 学员完成 Starter Lab |
| 65–75 分钟 | 自动验收、错误复盘与迁移问题 |

---

## 三、课程开场与摸底（第 1 课前 5 分钟）

【讲师逐字说】

“前一个模块里，我们已经能用 Gin 接收请求、校验参数并返回 JSON。今天开始解决下一层问题：数据怎样可靠地落到数据库，怎样查询，怎样修改，以及多个写操作怎样保持一致。”

“GORM 的价值不是让我们从此不用懂 SQL。恰恰相反，使用 ORM 时更需要知道它会生成什么 SQL、何时开启事务、零值是否被忽略、关联数据如何加载。今天每个 Demo 都先预测 SQL 和最终数据，再运行验证。”

“请记住一条课堂纪律：程序没有报错，不代表数据一定正确。凡是写操作，我们都要看最终状态；凡是事务，我们都要确认失败后数据是否恢复。”

【板书】

```text
HTTP Handler
    ↓
Service / business rule
    ↓
Repository / GORM
    ↓
database/sql pool
    ↓
MySQL
```

【提问】

1. “`gorm.Open` 返回的对象和 `database/sql.DB` 是同一个层次吗？”
2. “转账扣款成功、加款失败时，最终数据应该是什么状态？”
3. “把用户输入直接拼进 SQL 会有什么风险？”

【讲师收束】

“这三个问题分别对应连接管理、事务和参数化查询，也是今天六节课的主线。”

---

## 四、第 1 课：连接、配置与连接池（75 分钟）

### 0–15 分钟：GORM 与 `sql.DB`

【讲师逐字说】

“先区分两个对象。`*gorm.DB` 是 ORM 会话入口，负责模型映射、查询构造、回调和事务 API。底层的 `*sql.DB` 负责连接池。名字里虽然有 DB，但 `sql.DB` 不是一条固定连接，而是并发安全的连接池句柄。”

“如果有 Java 经验，可以把 `sql.DB` 类比成 DataSource 或 HikariCP，把 `gorm.DB` 类比成更轻量的 ORM 操作入口。类比只帮助定位，具体生命周期仍以 Go 和 GORM API 为准。”

“连接池三个参数今天必须讲清：`MaxOpenConns` 限制最大打开连接数；`MaxIdleConns` 限制空闲连接数；`ConnMaxLifetime` 限制连接最长复用时间。最大连接不是越大越好，因为应用连接数最终会压到 MySQL。”

【板书】

```text
gorm.Open(...) → *gorm.DB
db.DB()        → *sql.DB（连接池）

MaxOpenConns >= MaxIdleConns >= 0
MaxOpenConns > 0
```

【讲师逐字说】

“`gorm.Open` 成功也不能代替完整环境检查。课堂代码会调用 `Ping`，主动验证数据库地址、账号、密码和网络是否真的可用。”

### 15–35 分钟：连接 Demo

【操作】打开 `module04_gorm/01_setup/main.go`，先只展示 `run()`。

【讲师逐字说】

“请先不要运行。大家沿代码找出六个阶段：引用课堂 DSN、打开 GORM、取得连接池、配置连接池、Ping、迁移并写入一条教学数据。”

【操作】依次高亮：

```go
classroomdb.DSN
gorm.Open(mysql.Open(classroomdb.DSN), &gorm.Config{})
db.DB()
SetMaxOpenConns / SetMaxIdleConns / SetConnMaxLifetime
sqlDB.Ping()
db.AutoMigrate(&Product{})
```

【提问】“为什么课堂 Demo 可以固定本机连接，而生产系统不能照搬？”

【讲师收束】

“课堂环境固定是为了快速演示；生产凭证必须由安全配置系统注入，并按环境隔离、定期轮换。这里的固定 DSN 不能复制到真实项目。”

【操作】需要演示连接失败时，临时把 `classroomdb.DSN` 中的端口改为 `3307`，运行：

```bash
go run ./module04_gorm/01_setup
```

【预期】出现连接失败信息，并以非零状态退出。演示后立即把端口恢复为 `3306`。

【操作】再运行正常路径：

```bash
go run ./module04_gorm/01_setup
```

【预期】输出类似：

```text
connection=ok max_open=10 product=M04-L01-A001 price=100
```

【讲师逐字说】

“`FirstOrCreate` 让 Demo 可以重复运行。这里重复可运行是课堂便利，不等于所有业务写入都应该这样处理。真实业务还要区分创建、更新和幂等语义。”

【操作】终端 B 做只读检查：

```sql
SELECT id, code, price FROM m04_l01_products;
```

### 35–65 分钟：Pool Starter Lab

【操作】投影 Starter 和公开测试：

- `module04_gorm/01_setup/lab/starter/lab.go`
- `module04_gorm/01_setup/lab/starter/lab_exercise_test.go`

【讲师逐字说】

“现在请实现 `ValidatePool(maxOpen, maxIdle int) error`。合法条件有三个：最大打开连接必须大于零；最大空闲连接不能小于零；最大空闲连接不能大于最大打开连接。合法返回 nil，非法返回包含参数值的错误。”

“先运行一次看到 RED，再实现。不要改测试，不要打开 Solution。”

【学员操作】

```bash
go test -tags=exercise ./module04_gorm/01_setup/lab/starter
```

【三级提示】

1. 先把三个非法条件写成布尔表达式。
2. 非法条件之间使用逻辑或。
3. 用 `fmt.Errorf` 返回 `max_open` 与 `max_idle`，合法路径返回 nil。

【验收】同一命令由 RED 变 GREEN。

### 65–75 分钟：复盘

【讲师逐字说】

“这一节带走四句话。第一，`gorm.DB` 负责 ORM，`sql.DB` 负责连接池。第二，`sql.DB` 应长期复用，不要每个请求创建一次。第三，连接参数必须满足约束，也要结合 MySQL 上限和实例数量统一计算。第四，打开对象不等于数据库可用，课堂启动检查要 Ping。”

【复盘问题】

1. “十个应用实例，每个 `MaxOpenConns` 是 50，理论上可能给数据库带来多少连接？”
2. “为什么 DSN 不应该提交到 Git？”
3. “`defer sqlDB.Close()` 应放在每个 HTTP handler 里吗？”

【参考答案】“理论上最多约 500 条应用连接；DSN 包含凭证和环境信息；连接池应在进程生命周期内复用，在应用退出时关闭。”

---

## 五、第 2 课：模型、关系、软删除与迁移（75 分钟）

### 0–15 分钟：模型和关系

【讲师逐字说】

“GORM 模型不是只把表字段翻译成 Go 字段。它还表达主键、索引、外键字段、关联和软删除行为。结构体 Tag 是字符串配置，编译器不会替我们验证所有数据库语义，所以迁移和测试仍然重要。”

【板书】

```text
Author 1 ─── N Post
Post   N ─── N Tag
              ↓
       m04_l02_post_tags
```

【讲师逐字说】

“一对多由 `Post.AuthorID` 保存外键，`Author.Posts` 表达集合。多对多需要中间表，当前模型通过 `many2many:m04_l02_post_tags` 指定。”

“`gorm.DeletedAt` 让普通 Delete 变成软删除：记录仍在表中，只是写入删除时间，普通查询默认过滤。`Unscoped` 会绕开这个过滤，也能执行硬删除，所以必须谨慎使用。”

“迁移方面，课堂用 `AutoMigrate` 展示 V1 增加 Email 到 V2。请不要把它讲成生产迁移万能工具。生产环境要有版本号、审查、回滚或前滚方案、数据回填和发布顺序。”

### 15–35 分钟：关系与迁移 Demo

【操作】打开 `module04_gorm/02_models_migrations/main.go`，先看 `Author`、`Post`、`Tag`、`ProfileV1`、`ProfileV2`。

【提问】“同一个 `m04_l02_profiles` 表，V1 和 V2 的区别是什么？”

【收束】“V2 新增 Email，用来演示 schema 演进。”

【操作】运行：

```bash
go run ./module04_gorm/02_models_migrations
```

【预期】

```text
relations=ok author=M04 Alice tags=2 migration=email soft_delete=ok hard_delete=ok
```

【操作】终端 B 检查表与列：

```sql
SHOW TABLES LIKE 'm04_l02%';
SHOW COLUMNS FROM m04_l02_profiles LIKE 'email';
SELECT * FROM m04_l02_post_tags LIMIT 10;
```

【讲师逐字说】

“代码先迁移 V1，再迁移 V2，然后用 `HasColumn` 验证 Email 真正存在。接着创建作者和标签，创建文章时写入关系，使用两个 `Preload` 装载作者和标签。”

“删除部分先软删除文章，再用 `Unscoped` 找回。硬删除前先清除标签关联，最后硬删除文章。注意，共享 Tag 本身没有被删除，只清除了关联。”

【提问】“为什么删除一篇文章时不能顺手删除所有关联 Tag？”

【收束】“Tag 可能被其他文章共享。删除关联和删除实体是两件事。”

### 35–65 分钟：Relation Starter Lab

【讲师逐字说】

“Starter 用一个简化的 `Relation` 表达作者、文章和标签三个键。请实现校验：三个 ID 必须全部非零，任意一个为零都返回错误。这个练习不假装替代数据库外键，它训练的是在进入持久化前明确输入契约。”

【学员操作】

```bash
go test -tags=exercise ./module04_gorm/02_models_migrations/lab/starter
```

【三级提示】

1. 读取 `r.AuthorID`、`r.PostID`、`r.TagID`。
2. 任意一个为零就是非法。
3. 返回错误时说明是 relation key 非法，合法时返回 nil。

【验收】测试转绿，并补一个 TagID 为零的本地测试。

### 65–75 分钟：复盘

【讲师逐字说】

“今天不要混淆四件事：Go 字段、数据库列、外键关系和对象关联。`AuthorID` 是存储外键，`Author` 是关联对象，`Posts` 是集合视图，中间表承载多对多。”

“软删除提高恢复能力，但也意味着唯一索引、统计查询、存储增长和清理策略都需要重新设计。`Unscoped` 不是常规查询开关，而是越过默认安全边界的工具。”

【复盘问题】

1. “普通 Delete 和 Unscoped Delete 分别改变什么？”
2. “清除文章—标签关联会删除 Tag 表中的共享标签吗？”
3. “为什么 AutoMigrate 不能替代生产版本迁移？”

---

## 六、第 3 课：CRUD 与零值更新（75 分钟）

### 0–15 分钟：CRUD 语义

【讲师逐字说】

“CRUD 的难点不在记住 Create、Find、Update、Delete 四个名字，而在理解每个 API 对空结果、零值、条件和批量行为的语义。”

“`First` 期望一条记录，没有记录时返回 `gorm.ErrRecordNotFound`。`Find` 查询切片时，没有记录通常返回空切片和 nil error。业务层必须根据场景选择，不能把所有查询都写成同一种。”

“本节最重要的坑是结构体零值。GORM 用结构体做 `Updates` 时，默认只更新非零字段。数字 0、布尔 false、空字符串都可能被跳过。用 map 更新时，键明确出现，零值也会写入。”

【板书】

```text
Updates(Item{Stock: 0})
→ struct 零值默认被忽略

Updates(map[string]any{"stock": 0})
→ 键明确存在，写入 0
```

【提问】“库存从 100 清零，是‘没传库存’还是‘明确设置为 0’？”

【收束】“在更新契约里必须区分缺失与显式零值。HTTP 层常用指针或专门的 Patch DTO 表达这种差异。”

### 15–35 分钟：CRUD Demo

【操作】打开 `module04_gorm/03_crud/main.go`，运行前停在：

```go
db.Model(&first).Updates(Item{Stock: 0})
```

【提问】“执行以后，数据库 Stock 是 0 还是原来的 100？”

【操作】运行：

```bash
go run ./module04_gorm/03_crud
```

【预期】

```text
batch=2 first=M04-L03-PEN struct_zero=100 map_zero=0 delete=ok
```

【讲师逐字说】

“`struct_zero=100` 证明结构体的 Stock 0 被忽略；随后 map 明确包含 `stock` 键，所以 `map_zero=0`。不要只凭 API 名字猜数据库结果，必须知道 GORM 的字段选择规则。”

【操作】依次讲解：

1. `Create(&items)` 批量创建。
2. `First` 读取单条。
3. `Find` 读取集合。
4. 结构体 Updates 与 map Updates。
5. `Delete(&first)` 删除当前记录。

【操作】终端 B 验证：

```sql
SELECT sku, name, stock, price
FROM m04_l03_items
WHERE sku LIKE 'M04-L03-%';
```

### 35–65 分钟：StockUpdate Starter Lab

【讲师逐字说】

“请实现 `StockUpdate(stock int) map[string]any`。无论 stock 是不是零，都返回包含 `stock` 键的 map。目标是让调用方能把 0 明确交给 GORM。”

【学员操作】

```bash
go test -tags=exercise ./module04_gorm/03_crud/lab/starter
```

【验收】`StockUpdate(0)` 返回的 map 中存在 `stock`，值的动态类型和值都符合测试。

【加练】让函数同时接收 `active bool`，返回包含 `stock` 与 `active` 两个键的 map，验证 false 不丢失。

【三级提示】

1. 返回值不能是 nil。
2. map 的键使用数据库字段名 `stock`。
3. 直接返回 `map[string]any{"stock": stock}`。

### 65–75 分钟：复盘

【讲师逐字说】

“请带走三条更新原则。第一，更新前先定义字段缺失和零值的业务语义。第二，结构体更新适合只写非零字段，map 或 Select 适合显式字段更新。第三，任何由用户决定的查询条件都应参数化，不拼字符串。”

【复盘问题】

1. “查询一条必须存在的订单，应该优先 First 还是 Find？”
2. “把 `Active` 从 true 改成 false 时，结构体 Updates 有什么风险？”
3. “批量 Update 或 Delete 为什么必须检查 Where 条件？”

---

## 七、第 4 课：条件查询、JOIN、Preload 与 Hooks（75 分钟）

### 0–15 分钟：查询结果与关联装配

【讲师逐字说】

“这一节不追求把所有查询都写成 ORM 链，而是根据结果形状选择工具。查询完整模型和关联对象时，可以使用 Where 与 Preload；需要投影成报表行时，可以用 Table、Select、Joins 和 Scan。”

“`Preload` 的职责是把关联结果装配到模型字段。`JOIN` 的职责是让数据库按连接条件组合行。两者可能解决不同问题，不要看到有关联就机械地只选其中一个。”

【板书】

```text
模型 + 关联集合 → Preload
投影、筛选、报表 → JOIN + Select + Scan

外部值 → ? 参数
表名/列名 → 只能来自受控代码，不接受用户任意输入
```

【讲师逐字说】

“Hooks 是模型生命周期回调。适合做短小、确定、可测试的规范化、ID 生成和底线校验。不适合发送邮件、调用远程服务、执行长事务或悄悄改变业务单位。”

### 15–27 分钟：查询 Demo

【操作】打开 `module04_gorm/04_queries_hooks/queries/main.go`，让学生先找两种结果形状：`[]Category` 与 `[]JoinedProduct`。

【操作】运行：

```bash
go run ./module04_gorm/04_queries_hooks/queries
```

【预期】

```text
where=ok preload_active=1 joins=2
```

【讲师逐字说】

“第一段查询把分类装进 `[]Category`，并且只预加载 active 产品。第二段查询不需要完整模型，而是把产品名和分类名投影到 `JoinedProduct`。”

【操作】高亮：

```go
Preload("Products", "active = ?", true)
Where("name LIKE ?", "M04%")
Table("m04_l04_products AS p")
Select(...)
Joins(...)
Where("p.price >= ?", 5000)
Scan(&joined)
```

【提问】“价格 5000 是外部值时，为什么用问号参数，而不是 `fmt.Sprintf` 拼进 SQL？”

【收束】“参数和值分离，交给驱动安全绑定，也能避免引号和类型转换错误。”

### 27–35 分钟：Hooks Demo

【操作】打开 `module04_gorm/04_queries_hooks/hooks/main.go`，先展示 `BeforeCreate` 与 `BeforeUpdate`，再运行：

```bash
go run ./module04_gorm/04_queries_hooks/hooks
```

【预期】

```text
before_create=id_generated name="M04 Hook Product" before_update=blocked_negative
```

【讲师逐字说】

“创建前，Hook 去掉首尾空格、检查名字和价格、生成 UUID。更新前，Hook 阻止负价格。错误从 Hook 返回后，GORM 写操作失败，调用方必须继续处理这个错误。”

“这里的 Hook 没有把元改成分，也没有调用库存服务。单位转换和跨服务调用属于显式业务逻辑，藏进 Hook 会让写入行为难以预测和测试。”

### 35–65 分钟：查询与 Hook 实操

【讲师逐字说】

“本节不使用已经不在当前课程范围内的旧 Starter。请直接在当前查询 Demo 上完成两个小改动。”

【学员任务 A】增加一个 `minPrice` 变量，值为 6000，让 JOIN 查询继续通过 `?` 参数筛选；运行后只返回符合条件的投影行，并打印第一行产品名。

【学员任务 B】给 `BeforeCreate` 增加名称长度校验：Trim 后少于 3 个字符返回错误。新增一个本地测试或临时调用，证明非法名称不会写入。

【验收命令】

```bash
go run ./module04_gorm/04_queries_hooks/queries
go run ./module04_gorm/04_queries_hooks/hooks
```

【数据库只读验收】

```sql
SELECT p.name, p.price, p.active, c.name AS category
FROM m04_l04_products AS p
JOIN m04_l04_categories AS c ON c.id = p.category_id
ORDER BY p.id;
```

【三级提示】

1. `minPrice` 只替换参数值，不拼接 SQL。
2. 名称长度校验必须在 Trim 之后。
3. Hook 返回 error 后，调用方要检查 `db.Create(...).Error`。

### 65–75 分钟：复盘

【讲师逐字说】

“选择查询方式前，先画出想要的结果形状。需要完整模型和关联集合，就让模型承载；需要少量投影列，就定义专门结果结构体并 Scan。查询越复杂，越要观察实际 SQL 和数据库执行计划，而不是只看链式 API 是否优雅。”

“Hooks 只做局部、确定的模型约束。重要业务步骤应该在 Service 或明确事务中出现，让读代码的人看得见顺序和失败路径。”

【复盘问题】

1. “Preload 后得到的结果放在哪里？”
2. “为什么报表查询适合单独定义 JoinedProduct？”
3. “哪些逻辑不应该藏进 BeforeCreate？”

---

## 八、第 5 课：事务、回滚与 SavePoint（75 分钟）

### 0–15 分钟：原子性和边界

【讲师逐字说】

“事务的核心不是调用一个 API，而是定义哪些数据库变化必须一起成功或一起失败。转账最经典：扣款和加款是一个业务动作，任何一步失败，两边余额都必须回到操作前。”

【板书】

```text
BEGIN
  SELECT source
  SELECT target
  UPDATE source - amount
  UPDATE target + amount
COMMIT

任一步 error → ROLLBACK
```

【讲师逐字说】

“GORM 闭包事务 `db.Transaction(func(tx *gorm.DB) error {...})` 的规则很清晰：闭包返回 nil 就提交，返回 error 就回滚。因此事务内所有查询和更新都必须使用 tx，不能不小心继续使用外层 db。”

“SavePoint 是事务中的局部标记。回滚到 SavePoint 只撤销标记之后的操作，外层事务仍可继续。它不是完全独立的新事务，也不要把它机械等同于所有框架的嵌套事务传播。”

【提问】“数据库事务能不能自动撤回已经发送出去的短信？”

【收束】“不能。数据库事务只管理数据库资源。跨系统一致性需要 outbox、消息、补偿或工作流等额外设计。”

### 15–35 分钟：转账与 SavePoint Demo

【操作】打开 `module04_gorm/05_transactions/main.go`，先让学生预测最终 Alice 余额。

初始余额：Alice 100，Bob 50。依次执行：

1. 成功转账 30。
2. 尝试转账 1000，余额不足。
3. 尝试转账 10，在第二步前注入失败。
4. Alice 临时加 5，然后回滚到 SavePoint。

【提问】“四组操作完成后，Alice 应该是多少？”

【操作】运行：

```bash
go run ./module04_gorm/05_transactions
```

【预期】

```text
success=ok insufficient=rolled_back second_step=rolled_back savepoint=ok alice=70
```

【讲师逐字说】

“只有第一笔成功，所以 Alice 从 100 变 70。余额不足没有写入；第二步注入失败时，第一条扣款也被回滚；SavePoint 之后加的 5 被局部回滚。最终状态验证比‘函数返回了 error’更重要。”

【操作】终端 B 查询最终状态：

```sql
SELECT owner, balance
FROM m04_l05_wallets
ORDER BY owner;
```

【操作】高亮事务内的 `tx.Where`、`tx.Model` 和错误返回，强调不使用外层 `db`。

### 35–65 分钟：Transfer Starter Lab

【讲师逐字说】

“Starter 用纯函数模拟事务最终状态。成功时 from 减少、to 增加；金额小于等于零或余额不足时返回 error，并保持两个原值不变。请先保证失败不产生半成品状态。”

【学员操作】

```bash
go test -tags=exercise ./module04_gorm/05_transactions/lab/starter
```

【验收】

- `Transfer(100, 50, 30)` 得到 70、80、nil。
- `Transfer(10, 50, 30)` 返回 error，余额仍为 10、50。
- 学生补充 amount 为 0 和负数的测试。

【三级提示】

1. 先判断 amount 是否大于零、from 是否足够。
2. 错误路径直接返回原始 from 和 to。
3. 只有通过校验后才计算两个新余额。

### 65–75 分钟：复盘

【讲师逐字说】

“判断事务正确至少看三件事：边界是否覆盖完整业务动作；事务内是否始终使用 tx；失败后数据库最终状态是否保持不变量。程序返回错误只是证据之一，不是全部证据。”

【复盘问题】

1. “闭包返回 error 时 GORM 做什么？”
2. “为什么事务中混用 db 和 tx 会破坏原子性？”
3. “扣款成功、发送消息失败，数据库事务能自动解决吗？”

---

## 九、第 6 课：Raw SQL、Exec 与 ORM 选型（75 分钟）

### 0–15 分钟：何时使用 Raw SQL

【讲师逐字说】

“ORM 和 Raw SQL 不是非此即彼。简单模型 CRUD 和关联装配可以优先用 ORM；复杂聚合、窗口函数、精确投影、批量更新或数据库特性，可以使用参数化 Raw SQL。选择标准是正确、清晰、可测试和可观察，不是谁的代码更短。”

“`Raw` 通常用于返回行的查询，配合 `Scan`；`Exec` 用于不返回结果集的更新、删除或 DDL，并检查 `RowsAffected`。无论哪种方式，外部输入必须作为参数传递。”

【板书】

```text
错误："... WHERE name = '" + input + "'"
正确："... WHERE name = ?", input

Raw  + Scan → 查询结果
Exec        → Error + RowsAffected
```

【讲师逐字说】

“参数化只能保护值。表名、列名和排序方向通常不能作为普通占位符传递；如果业务允许动态选择，必须由服务端白名单映射，不能直接使用用户输入。”

### 15–35 分钟：Raw SQL Demo

【操作】打开 `module04_gorm/06_raw_sql/main.go`，让学生找出四种 SQL 用法：

1. 问号参数查询单个用户。
2. 聚合和分组。
3. `Exec` 批量更新。
4. `sql.Named` 命名参数。

【操作】运行：

```bash
go run ./module04_gorm/06_raw_sql
```

【预期】输出类似：

```text
raw=M04 Alice aggregate_groups=1 exec_rows=1 named=M04 Alice injection_safe=parameters
```

`exec_rows` 可能受重复运行后的当前数据状态影响；讲解重点是必须读取 `RowsAffected`，不要把固定行数当成永久业务保证。

【讲师逐字说】

“第一条 Raw 把名字作为独立参数。聚合查询将结果扫描到 `Summary`，字段别名 `avg_age` 映射到 `AvgAge`。Exec 把未成年课堂用户状态改为 0。命名参数提高复杂语句的可读性。”

“索引创建使用固定、受控的索引名，并先用 Migrator 检查是否存在。生产 DDL 应进入版本化迁移，而不是由普通应用请求路径临时执行。”

【操作】终端 B 只读检查：

```sql
SELECT name, age, status
FROM m04_l06_users
WHERE name LIKE 'M04%'
ORDER BY id;

SHOW INDEX FROM m04_l06_users
WHERE Key_name = 'idx_m04_l06_status';
```

### 35–65 分钟：参数化 Starter Lab

【讲师逐字说】

“请实现 `FindByName(name string) (string, []any)`。SQL 必须固定为带一个问号占位符的查询，name 原样放入参数切片。我们使用带引号和注释符号的字符串测试，但绝不执行拼接版危险 SQL。”

【学员操作】

```bash
go test -tags=exercise ./module04_gorm/06_raw_sql/lab/starter
```

测试输入：

```text
Alice' OR 1=1 --
```

【讲师逐字说】

“这个输入应该只作为一个普通名字参数。不要为了演示风险而把它拼接成 SQL 后去数据库执行。安全教学不需要真的执行危险语句。”

【验收】

- query 精确包含 `WHERE name = ?`。
- args 长度为 1。
- args[0] 等于原始 name。
- query 字符串中不包含用户输入。

【三级提示】

1. SQL 文本是固定常量。
2. 参数切片类型是 `[]any`。
3. 返回 `query, []any{name}`。

### 65–75 分钟：复盘与课程收口

【讲师逐字说】

“完成这个模块后，我们不应该变成只会写 GORM 链的人，而应该能回答四个问题：底层连接池是否健康；模型和关系是否表达正确；写操作失败后数据是否一致；最终执行的 SQL 是否安全且符合预期。”

“ORM 适合减少重复映射，Raw SQL 适合精确表达数据库能力。两者都必须参数化、检查 error、设置 context，并通过测试和数据库最终状态验证。”

【复盘问题】

1. “Raw 和 Exec 的结果形状有什么不同？”
2. “为什么动态列名不能直接来自用户输入？”
3. “复杂 SQL 应该完全放弃测试吗？”

【参考收束】“不能。可以测试参数构造、结果扫描和业务行为，并用真实数据库验证方言与执行结果。”

---

## 十、博客综合 Lab：作业启动讲稿（建议 35–45 分钟）

> 本环节在六节核心课后进行。只展示 Starter、目标行为和验收方式，不打开或复制 Solution。

### 1. 任务定位（5 分钟）

【讲师逐字说】

“最后把六节课组合到博客数据层。我们有文章、评论和标签。文章对评论是一对多，文章对标签是多对多。创建文章时，文章、标签关联和首条评论必须在同一事务。删除时软删除文章和评论、清除关联，但保留共享标签。”

【板书】

```text
Post 1 ─── N Comment
Post N ─── N Tag
              ↓
      m04_blog_post_tags
```

### 2. 只看 Starter 和测试（8 分钟）

【操作】打开：

- `module04_gorm/integrated_lab/blog_api/starter/starter.go`
- `module04_gorm/integrated_lab/blog_api/starter/starter_exercise_test.go`

【讲师逐字说】

“TODO 分为四组。第一，给 Post 增加 Comments 和 Tags 关系。第二，NormalizeTags 做 Trim、小写、去空、去重。第三，创建流程必须调用事务 runner。第四，删除流程也必须调用事务 runner。”

“公开测试只验证可公开契约，不代表全部生产要求。请按检查点逐个变绿，不要一次写完所有代码再调试。”

### 3. 建立 RED（3 分钟）

【操作】

```bash
go test -tags=exercise ./module04_gorm/integrated_lab/blog_api/starter/...
```

【预期】Starter 初始失败。失败原因应该是明确的行为断言，而不是缺包、数据库连接或编译环境错误。

### 4. 实现顺序提示（12 分钟）

【讲师逐字说】

“推荐顺序不是按文件从上往下，而是按风险从小到大。”

【板书】

```text
1. NormalizeTags 纯函数
2. Post 关系 Tag
3. 创建事务边界
4. 删除事务边界
5. HTTP 状态码
6. 真实 MySQL 验收
```

【讲师逐字说】

“NormalizeTags 最容易快速测试。模型关系完成后再迁移。事务先保证边界，再填具体 Repository 步骤。删除共享 Tag 是错误行为，只清除文章的关联。HTTP 层要把参数错误、未找到和内部错误映射成稳定状态码。”

### 5. 最终行为演示（5 分钟）

教师可运行 Solution 验证目标，但不要向学生展示实现代码：

```bash
go test ./module04_gorm/integrated_lab/blog_api/solution/...
make run-blog-api
```

另一个终端：

```bash
curl -X POST http://localhost:8091/posts \
  -H 'Content-Type: application/json' \
  -d '{"title":"GORM","content":"relations","tags":["go","gorm"]}'

curl http://localhost:8091/posts
```

【讲师逐字说】

“你们要复制的是行为，不是教师代码。提交前必须能解释：关系怎样表达，创建为什么需要事务，删除为什么保留 Tag，外部输入怎样参数化。”

### 6. 作业验收说明（5 分钟）

【讲师逐字说】

“提交前先运行 Starter 测试，再运行自己新增的测试。不要把真实数据库密码、教师 Solution 或本地数据库文件提交到项目中。出现数据库错误时，先区分是连接、迁移、SQL 还是业务状态，不要通过删除整个数据库来逃避定位。”

---

## 十一、现场故障速查

### 1. 固定课堂 DSN 与本机配置不一致

检查 `module04_gorm/internal/classroomdb/db.go` 中的 `DSN`：

```text
root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local
```

按课堂 MySQL 的实际账号、密码和端口修改；不要填入生产或个人真实凭证。

### 2. `connect: connection refused`

确认 MySQL 服务正在运行并监听预期端口：

```bash
lsof -nP -iTCP:3306 -sTCP:LISTEN
```

然后检查代码内 DSN 的地址和端口。

### 3. `Unknown database 'gorm_demo'`

说明数据库尚未创建或 DSN 指向错误实例。回到课前 SQL 创建课堂库，不要让应用自动创建任意数据库。

### 4. `Access denied`

确认用户名、密码、Host 授权和课堂库权限。固定的 `root:password` 只适用于隔离的本机课堂环境；真实系统应使用最小权限账号和外部密钥配置。

### 5. 表或索引冲突

先用只读语句确认对象：

```sql
SHOW TABLES LIKE 'm04\_%';
```

只处理确认属于本课程的 `m04_*` 对象，不清空其他数据。

### 6. 零值没有更新

检查是否使用结构体 Updates。若业务需要显式写入 0、false 或空字符串，使用 map、`Select` 或明确 Patch DTO，并补测试证明字段缺失和零值的区别。

### 7. 事务函数返回 error，但数据仍被改了

检查事务闭包中是否混用了外层 `db` 与事务 `tx`；再查询最终状态，不要只读日志。

### 8. 参数化查询仍报语法错误

确认占位符数量与参数数量相同，表名和列名来自受控代码，MySQL 方言与当前驱动匹配。不要退回字符串拼接。

### 9. Hook 行为难以解释

把跨服务调用、长耗时任务和业务编排移出 Hook。Hook 只保留局部规范化、ID 生成和底线校验。

---

## 十二、教师课后验收

离线检查：

```bash
make module04-verify
```

MySQL 环境和当前六个 Demo：

```bash
make module04-env-check
go run ./module04_gorm/01_setup
go run ./module04_gorm/02_models_migrations
go run ./module04_gorm/03_crud
go run ./module04_gorm/04_queries_hooks/queries
go run ./module04_gorm/04_queries_hooks/hooks
go run ./module04_gorm/05_transactions
go run ./module04_gorm/06_raw_sql
```

六个 Starter 的显式验收：

```bash
go test -tags=exercise ./module04_gorm/01_setup/lab/starter
go test -tags=exercise ./module04_gorm/02_models_migrations/lab/starter
go test -tags=exercise ./module04_gorm/03_crud/lab/starter
go test -tags=exercise ./module04_gorm/05_transactions/lab/starter
go test -tags=exercise ./module04_gorm/06_raw_sql/lab/starter
go test -tags=exercise ./module04_gorm/integrated_lab/blog_api/starter/...
```

学生应能口头回答：

1. `gorm.DB` 与 `sql.DB` 分别负责什么？
2. 一对多和多对多分别怎样表达？
3. 软删除与硬删除有什么区别？
4. 为什么结构体 Updates 可能跳过零值？
5. Preload、JOIN 和投影结构体分别适合什么结果？
6. 哪些逻辑适合 Hook，哪些不适合？
7. 为什么事务内必须使用 tx，并验证最终状态？
8. SavePoint 回滚和整个事务回滚有什么区别？
9. Raw 与 Exec 分别用于什么？
10. 为什么用户输入必须作为参数，而不能拼进 SQL？
