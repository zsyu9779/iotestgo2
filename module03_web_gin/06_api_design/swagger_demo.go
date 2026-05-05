// Swagger 文档生成演示
//
// 安装 swag：
//   go install github.com/swaggo/swag/cmd/swag@latest
//
// 生成文档：
//   swag init
//
// 访问文档：
//   http://localhost:8080/swagger/index.html
package main

// @title           iotestgo2 API
// @version         1.0
// @description     教学演示 API 文档
// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

import (
	"fmt"
	"net/http"
)

// 以下为 swag 注释 + 代码结构演示，实际运行需要配合 gin-swagger

// @Summary      用户注册
// @Description  创建新用户账号
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body body RegisterRequest true "注册信息"
// @Success      200  {object}  RegisterResponse
// @Failure      400  {object}  ErrorResponse
// @Router       /api/v1/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	// 实际实现省略
}

// @Summary      用户登录
// @Description  使用用户名密码登录获取 Token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body body LoginRequest true "登录信息"
// @Success      200  {object}  LoginResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /api/v1/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	// 实际实现省略
}

type RegisterRequest struct {
	Username string `json:"username" example:"gopher"`
	Password string `json:"password" example:"123456"`
	Email    string `json:"email" example:"go@example.com"`
}

type RegisterResponse struct {
	UserID int64  `json:"user_id" example:"1"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

type LoginRequest struct {
	Username string `json:"username" example:"gopher"`
	Password string `json:"password" example:"123456"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

type ErrorResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

func main() {
	fmt.Println("=== Swagger 文档集成 ===")
	fmt.Println()
	fmt.Println("1. 安装 swag: go install github.com/swaggo/swag/cmd/swag@latest")
	fmt.Println("2. 编写 swag 注释（@title, @Summary, @Param 等）")
	fmt.Println("3. 运行 swag init 生成 docs/ 目录")
	fmt.Println("4. 在路由中添加 Swagger UI Handler：")
	fmt.Println()
	fmt.Println(`   import (
       _ "your-project/docs"
       swaggerFiles "github.com/swaggo/files"
       ginSwagger "github.com/swaggo/gin-swagger"
   )

   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))`)
	fmt.Println()
	fmt.Println("5. 访问 http://localhost:8080/swagger/index.html")
	fmt.Println()
	fmt.Println("主要 swag 注释标签：")
	fmt.Println("  @Summary      - 接口简介")
	fmt.Println("  @Description  - 详细描述")
	fmt.Println("  @Tags         - 分组标签")
	fmt.Println("  @Accept       - 请求格式 (json/xml)")
	fmt.Println("  @Produce      - 响应格式")
	fmt.Println("  @Param        - 参数描述")
	fmt.Println("  @Success      - 成功响应")
	fmt.Println("  @Failure      - 失败响应")
	fmt.Println("  @Router       - 路由路径 [method]")
	fmt.Println("  @Security     - 安全方案")
}
