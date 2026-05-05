// 05 数据库与缓存：Cache-Aside 模式演示（用 sync.Map 模拟 Redis）
//
// 启动：go run main.go
//
// Cache-Aside 模式是 go-zero sqlx Model 内置的缓存一致性模式
package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache 模拟 Redis 缓存
type Cache struct {
	data sync.Map
}

func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.data.Load(key)
	if !ok {
		return "", false
	}
	return v.(string), true
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	fmt.Printf("    [Cache] SET %s = %s (TTL=%v)\n", key, value, ttl)
	c.data.Store(key, value)
}

func (c *Cache) Del(key string) {
	fmt.Printf("    [Cache] DEL %s\n", key)
	c.data.Delete(key)
}

// Database 模拟 MySQL 数据库
type Database struct {
	data map[int64]map[string]string
}

func NewDB() *Database {
	return &Database{
		data: map[int64]map[string]string{
			1: {"id": "1", "username": "gopher", "email": "gopher@example.com"},
			2: {"id": "2", "username": "alice", "email": "alice@example.com"},
		},
	}
}

func (db *Database) FindOne(id int64) (map[string]string, bool) {
	fmt.Printf("    [MySQL] SELECT * FROM users WHERE id=%d\n", id)
	time.Sleep(50 * time.Millisecond) // 模拟数据库查询耗时
	user, ok := db.data[id]
	return user, ok
}

func (db *Database) Update(id int64, field, value string) {
	fmt.Printf("    [MySQL] UPDATE users SET %s='%s' WHERE id=%d\n", field, value, id)
	if user, ok := db.data[id]; ok {
		user[field] = value
	}
}

// UserModel 带缓存的 Model（go-zero sqlx Model 的行为）
type UserModel struct {
	db    *Database
	cache *Cache
}

func NewUserModel(db *Database, cache *Cache) *UserModel {
	return &UserModel{db: db, cache: cache}
}

// FindOne 查用户（Cache-Aside 读流程）
func (m *UserModel) FindOne(id int64) (map[string]string, error) {
	cacheKey := fmt.Sprintf("cache:user:id:%d", id)

	// 1. 先查缓存
	fmt.Printf("  [Step 1] 查缓存 %s...\n", cacheKey)
	if v, ok := m.cache.Get(cacheKey); ok {
		fmt.Printf("    ✓ 缓存命中! 直接返回, 无需查 DB\n")
		return map[string]string{"cached": v}, nil
	}
	fmt.Printf("    ✗ 缓存未命中\n")

	// 2. 查数据库
	fmt.Printf("  [Step 2] 查数据库...\n")
	user, ok := m.db.FindOne(id)
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}

	// 3. 写回缓存
	fmt.Printf("  [Step 3] 写回缓存（下次命中）\n")
	m.cache.Set(cacheKey, user["username"], 5*time.Minute)

	return user, nil
}

// Update 更新用户（Cache-Aside 写流程）
func (m *UserModel) Update(id int64, field, value string) {
	// 1. 写数据库
	m.db.Update(id, field, value)

	// 2. 删除缓存（让下次读取重新查 DB 并写回缓存）
	cacheKey := fmt.Sprintf("cache:user:id:%d", id)
	m.cache.Del(cacheKey)
}

func main() {
	cache := &Cache{}
	db := NewDB()
	model := NewUserModel(db, cache)

	fmt.Println("=== Cache-Aside 模式演示 ===")
	fmt.Println()
	fmt.Println("Cache-Aside 是 go-zero sqlx Model 内置的缓存模式")
	fmt.Println()

	// ========== 第一次查：缓存未命中 → 查 DB → 写缓存 ==========
	fmt.Println("--- 第一次查询（缓存未命中）---")
	user, _ := model.FindOne(1)
	fmt.Printf("  结果: %v\n\n", user)

	// ========== 第二次查同一条：缓存命中 ==========
	fmt.Println("--- 第二次查询（缓存命中）---")
	user, _ = model.FindOne(1)
	fmt.Printf("  结果: %v\n\n", user)

	// ========== 更新：写 DB → 删缓存 ==========
	fmt.Println("--- 更新数据（写 DB → 删缓存）---")
	model.Update(1, "email", "newemail@example.com")
	fmt.Println()

	// ========== 更新后再查：缓存未命中 → 查 DB → 写缓存 ==========
	fmt.Println("--- 更新后查询（缓存已失效，重新查 DB）---")
	user, _ = model.FindOne(1)
	fmt.Printf("  结果: %v\n", user)

	fmt.Println()
	fmt.Println("=== 缓存三大问题及 go-zero 防护 ===")
	fmt.Println("  1. 穿透（查不存在的数据）：singleflight + 布隆过滤器")
	fmt.Println("  2. 击穿（热点 key 过期）：singleflight 合并并发请求")
	fmt.Println("  3. 雪崩（大量 key 同时过期）：TTL + 随机偏移（jitter）")
	fmt.Println()
	fmt.Println("=== go-zero 中的使用 ===")
	fmt.Println("  从 DDL 生成带缓存的 Model：")
	fmt.Println("    goctl model mysql ddl -src user.sql -dir internal/model -c")
	fmt.Println("    参数 -c 表示启用缓存模式（Cache-Aside）")
	fmt.Println()
	fmt.Println("  go-zero 的 sqlx Model 自动处理：")
	fmt.Println("    FindOne → 先查 Redis → 未命中查 MySQL → 写 Redis")
	fmt.Println("    Update/Delete → 自动删除对应 Redis key")
}
