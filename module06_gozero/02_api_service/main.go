// 02 API 服务开发：模仿 goctl 生成的 API 服务结构
//
// goctl 生成命令（可选安装 goctl 后运行）：
//   goctl api go -api user.api -dir .
//
// 生成后文件结构：
//   etc/user-api.yaml       配置文件
//   internal/config/config.go
//   internal/handler/        路由 Handler（自动生成，勿改）
//   internal/logic/          业务逻辑（核心，手写）
//   internal/svc/servicecontext.go 依赖注入容器
//   internal/types/types.go
//
// 启动：go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ========== Types（对应 goctl 生成的 types.go）==========

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,optional"`
}

type RegisterResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserInfoResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ========== Model（用户存储，生产环境用 MySQL）==========

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

// InMemoryUserStore 内存用户存储（教学用，实际项目用 sqlx Model）
type InMemoryUserStore struct {
	users  map[string]*User // username -> User
	nextID int64
}

func NewStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users:  make(map[string]*User),
		nextID: 1,
	}
}

func (s *InMemoryUserStore) Insert(username, password, email string) (*User, error) {
	if _, exists := s.users[username]; exists {
		return nil, fmt.Errorf("user already exists")
	}
	u := &User{ID: s.nextID, Username: username, Password: password, Email: email}
	s.users[username] = u
	s.nextID++
	return u, nil
}

func (s *InMemoryUserStore) FindByUsername(username string) (*User, error) {
	u, ok := s.users[username]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *InMemoryUserStore) FindByID(id int64) (*User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// ========== ServiceContext（依赖注入容器）==========

type ServiceContext struct {
	Store *InMemoryUserStore
	// 实际项目还有: Config, UserModel, JwtAuth, Redis, etc.
}

// ========== Handler（对应 goctl 生成的 handler 层，自动生成勿改）==========

// 每个 Handler 做：解析请求 → 调用 Logic → 写响应

func RegisterHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := &RegisterLogic{svcCtx: svcCtx}
		resp, err := l.Register(r, &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func LoginHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := &LoginLogic{svcCtx: svcCtx}
		resp, err := l.Login(r, &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UserInfoHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := &UserInfoLogic{svcCtx: svcCtx}
		resp, err := l.UserInfo(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// ========== Logic 层（核心业务逻辑，手写）==========

type RegisterLogic struct {
	svcCtx *ServiceContext
}

func (l *RegisterLogic) Register(r *http.Request, req *RegisterRequest) (*RegisterResponse, error) {
	// 1. 参数校验
	if len(req.Username) < 3 {
		return nil, fmt.Errorf("username too short (min 3 chars)")
	}
	if len(req.Password) < 6 {
		return nil, fmt.Errorf("password too short (min 6 chars)")
	}

	// 2. 密码加密（简化版，实际用 bcrypt）
	// hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// 3. 写入存储
	user, err := l.svcCtx.Store.Insert(req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}

	logx.Infof("新用户注册: id=%d username=%s", user.ID, user.Username)

	return &RegisterResponse{
		UserId:   user.ID,
		Username: user.Username,
		Token:    fmt.Sprintf("jwt-token-%d", user.ID), // 简化版 token
	}, nil
}

type LoginLogic struct {
	svcCtx *ServiceContext
}

func (l *LoginLogic) Login(r *http.Request, req *LoginRequest) (*LoginResponse, error) {
	user, err := l.svcCtx.Store.FindByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	if user.Password != req.Password {
		return nil, fmt.Errorf("invalid username or password")
	}

	logx.Infof("用户登录: id=%d username=%s", user.ID, user.Username)
	return &LoginResponse{Token: fmt.Sprintf("jwt-token-%d", user.ID)}, nil
}

type UserInfoLogic struct {
	svcCtx *ServiceContext
}

func (l *UserInfoLogic) UserInfo(r *http.Request) (*UserInfoResponse, error) {
	// 实际项目从 JWT 提取 user_id
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		return nil, fmt.Errorf("missing user_id")
	}

	var id int64
	fmt.Sscanf(userID, "%d", &id)
	user, err := l.svcCtx.Store.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &UserInfoResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// ========== Main ==========

func main() {
	// 依赖注入
	svcCtx := &ServiceContext{
		Store: NewStore(),
	}

	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8881,
	})

	// 注册路由（Handler → Logic → Model 分层）
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/api/v1/user/register", Handler: RegisterHandler(svcCtx)})
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/api/v1/user/login", Handler: LoginHandler(svcCtx)})
	server.AddRoute(rest.Route{Method: http.MethodGet, Path: "/api/v1/user/info", Handler: UserInfoHandler(svcCtx)})

	fmt.Println("=== go-zero API 服务已启动 ===")
	fmt.Println("  端口: 8881")
	fmt.Println()
	fmt.Println("  测试：")
	fmt.Println("    curl -X POST http://localhost:8881/api/v1/user/register \\")
	fmt.Println("      -H 'Content-Type: application/json' \\")
	fmt.Println("      -d '{\"username\":\"gopher\",\"password\":\"secret123\"}'")
	fmt.Println()
	fmt.Println("    curl -X POST http://localhost:8881/api/v1/user/login \\")
	fmt.Println("      -H 'Content-Type: application/json' \\")
	fmt.Println("      -d '{\"username\":\"gopher\",\"password\":\"secret123\"}'")
	fmt.Println()
	fmt.Println("    curl http://localhost:8881/api/v1/user/info?user_id=1")
	fmt.Println()
	fmt.Println("  go-zero 分层架构：")
	fmt.Println("    Handler（路由+解析） → Logic（业务逻辑） → Model/Store（数据访问）")
	fmt.Println("    ServiceContext 装载所有依赖，实现依赖注入")
	fmt.Println()

	server.Start()

	_ = json.Marshal
	_ = log.Println
}
