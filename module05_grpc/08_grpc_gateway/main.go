// 08 gRPC-Gateway：同时提供 gRPC 和 HTTP/JSON 接口
//
// 启动：go run main.go
// gRPC: grpcurl -plaintext -d '{"name":"Gopher"}' localhost:50051 hello.Greeter/SayHello
// HTTP: curl -X POST http://localhost:8080/v1/hello -d '{"name":"Gopher"}'
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	if name == "" {
		name = "World"
	}
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s! (from gRPC server)", name),
	}, nil
}

func main() {
	// ========== 1. 启动 gRPC Server ==========
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterGreeterServer(grpcSrv, &server{})
	reflection.Register(grpcSrv)

	go func() {
		log.Println("gRPC server listening on :50051")
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("gRPC serve failed: %v", err)
		}
	}()

	// ========== 2. 启动 HTTP Gateway（gRPC → HTTP/JSON 转换） ==========
	// 建立到 gRPC server 的连接（同进程内通过 localhost 通信）
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("gateway dial failed: %v", err)
	}

	// 创建 HTTP mux 并注册 gRPC-Gateway handler
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		}),
	)

	// 手动注册 HTTP 路由 → gRPC proxy
	// 效果：POST /v1/hello → gRPC SayHello
	err = gwmux.HandlePath("POST", "/v1/hello", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		// 1. 读取 HTTP Body → Proto Message
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req := &pb.HelloRequest{}
		if err := protojson.Unmarshal(body, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 2. 调用 gRPC 服务
		client := pb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		resp, err := client.SayHello(ctx, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. Proto Message → JSON → HTTP Response
		respJSON, _ := protojson.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	if err != nil {
		log.Fatalf("register path failed: %v", err)
	}

	httpSrv := &http.Server{Addr: ":8080", Handler: gwmux}
	go func() {
		log.Println("HTTP Gateway listening on :8080")
		log.Println("  POST http://localhost:8080/v1/hello")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP serve failed: %v", err)
		}
	}()

	// ========== 3. 优雅退出 ==========
	fmt.Println()
	fmt.Println("=== gRPC-Gateway 已启动 ===")
	fmt.Println("  gRPC:  localhost:50051")
	fmt.Println("  测试:  grpcurl -plaintext -d '{\"name\":\"Gopher\"}' localhost:50051 hello.Greeter/SayHello")
	fmt.Println()
	fmt.Println("  HTTP:  localhost:8080")
	fmt.Println("  测试:  curl -X POST http://localhost:8080/v1/hello -H 'Content-Type: application/json' -d '{\"name\":\"Gopher\"}'")
	fmt.Println()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	httpSrv.Shutdown(context.Background())
	grpcSrv.GracefulStop()
}
