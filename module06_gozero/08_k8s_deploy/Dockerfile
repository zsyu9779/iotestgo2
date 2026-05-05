# go-zero 微服务 Dockerfile
# 生成命令：goctl docker -go user-api.go

# --- 多阶段构建 ---

# 阶段 1：编译
FROM golang:1.22-alpine AS builder

# 国内用户可替换为阿里云镜像
RUN apk add --no-cache make git

WORKDIR /app

# 缓存依赖层（利用 Docker 层缓存加速）
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/user-api .

# 阶段 2：运行（最小镜像）
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从编译阶段复制二进制文件
COPY --from=builder /app/user-api /app/user-api
COPY --from=builder /app/etc /app/etc

EXPOSE 8888

CMD ["./user-api"]
