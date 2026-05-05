// 02 API 服务：从 .api 文件生成代码，手写 Logic 层业务
//
// goctl 生成命令：goctl api go -api user.api -dir .
//
// 生成后的文件结构：
//   etc/user-api.yaml       配置文件
//   internal/config/config.go
//   internal/handler/        路由 Handler（自动生成）
//   internal/logic/          业务逻辑（手写核心）
//   internal/svc/servicecontext.go
//   internal/types/types.go
//   user.go                  main 入口
//   user-api.go              自动生成的路由注册
package main

import "fmt"

func main() {
	fmt.Println("=== 02 API 服务开发 ===")
	fmt.Println()

	fmt.Println("--- goctl 代码生成流程 ---")
	fmt.Println("1. 编写 user.api 定义文件")
	fmt.Println("2. 运行：goctl api go -api user.api -dir .")
	fmt.Println("3. goctl 自动生成：")
	fmt.Println("   - types.go: RegisterRequest/LoginRequest 等结构体")
	fmt.Println("   - handler/: UserRegisterHandler/UserLoginHandler 等")
	fmt.Println("   - config.go: JWT Secret、数据库连接等配置")
	fmt.Println("   - servicecontext.go: ServiceContext（装载各类依赖）")
	fmt.Println()

	fmt.Println("--- Handler 层（自动生成，勿改动） ---")
	fmt.Println(`  func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          var req types.RegisterRequest
          if err := httpx.Parse(r, &req); err != nil {
              httpx.ErrorCtx(r.Context(), w, err)
              return
          }
          l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
          resp, err := l.UserRegister(&req)
          if err != nil {
              httpx.ErrorCtx(r.Context(), w, err)
          } else {
              httpx.OkJsonCtx(r.Context(), w, resp)
          }
      }
  }`)
	fmt.Println()

	fmt.Println("--- Logic 层（核心业务，手写） ---")
	fmt.Println(`  func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (*types.RegisterResponse, error) {
      // 1. 参数校验
      if len(req.Username) < 3 { return nil, errors.New("username too short") }

      // 2. 检查用户是否已存在
      exist, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
      if err != nil && err != model.ErrNotFound { return nil, err }
      if exist != nil { return nil, errors.New("user already exists") }

      // 3. 密码加密
      hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

      // 4. 写入数据库
      user := &model.User{Username: req.Username, Password: string(hashedPwd), Email: req.Email}
      result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
      if err != nil { return nil, err }

      // 5. 生成 JWT Token
      userId, _ := result.LastInsertId()
      token, _ := l.svcCtx.JwtAuth.GenerateToken(userId)

      return &types.RegisterResponse{UserId: userId, Username: req.Username, Token: token}, nil
  }`)
	fmt.Println()

	fmt.Println("--- ServiceContext（依赖注入容器） ---")
	fmt.Println(`  type ServiceContext struct {
      Config    config.Config
      UserModel model.UserModel   // 数据库模型
      JwtAuth   *utils.JwtAuth    // JWT 工具
  }
  func NewServiceContext(c config.Config) *ServiceContext {
      conn := sqlx.NewMysql(c.Mysql.Dsn)
      return &ServiceContext{
          Config:    c,
          UserModel: model.NewUserModel(conn),
          JwtAuth:   utils.NewJwtAuth(c.JwtAuth.AccessSecret, c.JwtAuth.AccessExpire),
      }
  }`)
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: @RestController → @Service → @Repository")
	fmt.Println("  go-zero: Handler → Logic → Model")
	fmt.Println("  同理的分层思路，go-zero 的 Logic 层更薄、更聚焦单次请求")
}
