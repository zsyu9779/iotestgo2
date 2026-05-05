// 03 RPC 服务：使用 zRPC（go-zero 对 gRPC 的封装）开发 RPC 服务
//
// 生成命令：goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("=== 03 RPC 服务开发 (zRPC) ===")
	fmt.Println()

	fmt.Println("--- zRPC vs 原生 gRPC ---")
	fmt.Println("  go-zero 的 zRPC 在 gRPC 基础上增加了：")
	fmt.Println("  1. 自动服务注册到 Etcd（无需手动写注册代码）")
	fmt.Println("  2. 内置拦截器（熔断、限流、超时控制、日志）")
	fmt.Println("  3. 配置文件驱动（.yaml 配置而非硬编码）")
	fmt.Println("  4. 与 API 服务无缝集成")
	fmt.Println()

	fmt.Println("--- 服务端实现（logic 层） ---")
	fmt.Println(`  func (l *GetUserLogic) GetUser(in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
      // 从数据库查询用户
      user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
      if err != nil {
          return nil, status.Error(codes.NotFound, "user not found")
      }
      return &userpb.GetUserResponse{
          User: &userpb.User{
              Id:       user.Id,
              Username: user.Username,
              Email:    user.Email,
              Status:   int32(user.Status),
          },
      }, nil
  }`)
	fmt.Println()

	fmt.Println("--- 客户端调用 RPC ---")
	fmt.Println(`  // API 服务的 ServiceContext 中注入 RPC Client：
  type ServiceContext struct {
      Config      config.Config
      UserRpc     userpb.UserRpcClient     // RPC 客户端
  }

  func NewServiceContext(c config.Config) *ServiceContext {
      // 通过 Etcd 发现 RPC 服务
      conn := zrpc.MustNewClient(zrpc.RpcClientConf{
          Etcd: discov.EtcdConf{
              Hosts: []string{"localhost:2379"},
              Key:   "user.rpc",
          },
      })
      return &ServiceContext{
          Config:  c,
          UserRpc: userpb.NewUserRpcClient(conn.Conn()),
      }
  }`)
	fmt.Println()

	fmt.Println("--- API 调用 RPC（首个微服务拆分） ---")
	fmt.Println(`  func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (*types.UserInfoResponse, error) {
      userId := l.ctx.Value("userId").(int64)  // 从 JWT 提取

      // 调用 UserRpc 服务获取用户信息
      userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userpb.GetUserRequest{UserId: userId})
      if err != nil {
          return nil, err
      }

      return &types.UserInfoResponse{
          UserId:   userResp.User.Id,
          Username: userResp.User.Username,
          Email:    userResp.User.Email,
          Status:   int(userResp.User.Status),
      }, nil
  }`)
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: @FeignClient(name=\"user-service\") + Eureka")
	fmt.Println("  go-zero: zRPC Client + Etcd")
	fmt.Println("  go-zero 的 RPC 调用不依赖注解，通过 ServiceContext 注入，更显式")

	_ = context.Background
	_ = fmt.Sprint
}
