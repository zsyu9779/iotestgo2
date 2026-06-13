# project_ecommerce_standard

这是 Module 06 的标准 go-zero 实践版，和 `project_ecommerce` 的概念演示版并存。

## 目标调用链

HTTP Client -> order-api -> order-rpc -> user-rpc + product-rpc

## 生成命令

```bash
goctl api go -api api/order.api -dir api/order
goctl rpc protoc rpc/user/user.proto --go_out=rpc/user --go-grpc_out=rpc/user --zrpc_out=rpc/user
goctl rpc protoc rpc/product/product.proto --go_out=rpc/product --go-grpc_out=rpc/product --zrpc_out=rpc/product
goctl rpc protoc rpc/order/order.proto --go_out=rpc/order --go-grpc_out=rpc/order --zrpc_out=rpc/order
```

## 基础设施验证

```bash
docker compose up -d
curl http://localhost:9090/targets
```

本目录的 compose 文件只启动 Etcd、MySQL、Redis、Prometheus。四个 Go 服务仍在宿主机启动，Prometheus 通过 `host.docker.internal` 抓取 `8890`、`9101`、`9102`、`9103`。
