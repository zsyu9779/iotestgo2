# gRPC and go-zero Deep Improvement Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Upgrade Module 05 and Module 06 from good concept demos into high-confidence undergraduate teaching modules that show real RPC contracts, generated code workflows, testable business logic, and a credible microservice path.

**Architecture:** Keep small lesson demos for first exposure, then add a second layer of "standard practice" implementations. Module 05 will separate compute engine, gRPC transport, interceptors, gateway generation, and tests. Module 06 will keep simple concept demos while adding a generated-style go-zero ecommerce project that uses real proto clients, config files, service contexts, and a true API -> RPC -> RPC call chain.

**Tech Stack:** Go, gRPC-Go, Protocol Buffers, grpc-gateway, go-zero, goctl, Etcd, MySQL, Redis, Prometheus, Docker Compose, `grpcurl`, `curl`, `go test`.

---

## Official References Used for Direction

- go-zero official documentation emphasizes `goctl` code generation and built-in resilience features such as timeout control, rate limiting, circuit breaker, adaptive load shedding: <https://go-zero.dev/>
- go-zero getting started path lists Go, goctl, protoc, API DSL, and Hello World as the first learning steps: <https://go-zero.dev/getting-started/>
- gRPC interceptor guide lists logging, metadata handling, metrics, authentication, authorization, and policy enforcement as common interceptor use cases: <https://grpc.io/docs/guides/interceptors/>
- gRPC metadata guide defines metadata as key-value side-channel data for RPC calls: <https://grpc.io/docs/guides/metadata/>
- grpc-gateway documentation describes generated reverse proxies from protobuf services and `google.api.http` annotations: <https://github.com/grpc-ecosystem/grpc-gateway>
- grpc-gateway annotation tutorial shows adding `google/api/annotations.proto` and generating `*.gw.pb.go`: <https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/>

## Current gRPC Module Diagnosis

Strengths:
- Lessons 01-07 follow a clean progression: proto -> codegen -> unary -> streaming -> interceptors -> metadata -> status codes.
- `project_distributed_compute` already has server/client and demonstrates bidirectional streaming plus worker pool.
- Server comments are classroom-friendly.

Problems:
- `project_distributed_compute/server/main.go` mixes compute algorithm, stream transport, concurrency, logging, and process startup.
- Concurrent workers call `stream.Send` directly. In gRPC-Go, concurrent calls to `SendMsg` on the same stream are not a teaching pattern to normalize; serialize result sending through one goroutine.
- Compute logic is not unit-tested.
- `ComputeTask.operation` is a string, so invalid operations are only runtime errors.
- `08_grpc_gateway` manually maps HTTP to gRPC instead of demonstrating standard annotation-based gateway generation.
- `module05_grpc/08_grpc_gateway/proto/` contains compiled Mach-O binaries.

## Current go-zero Module Diagnosis

Strengths:
- The module explains API service, RPC service, Etcd, Cache-Aside, MQ, observability, and deployment in small digestible examples.
- `03_rpc_service` uses actual generated `userpb` client/server code and a real API -> RPC call.
- `project_ecommerce/docker-compose.yml` provides useful infrastructure.

Problems:
- Several lesson files "simulate" go-zero rather than use generated goctl structure. That is fine for first pass, but insufficient for a final microservice module.
- `project_ecommerce` hand-writes gRPC service descriptors instead of using proto-generated stubs.
- `order-api` does not really call `order-rpc`; it creates local fake responses.
- `order-rpc.checkUser` always returns true and does not invoke UserRpc.
- Product-RPC is listed in README but not implemented.
- Prometheus target ports in `prometheus.yml` do not match actual service ports.
- No go-zero standard directories such as `etc`, `internal/config`, `internal/logic`, `internal/svc`, `internal/types`.

---

## Target Module 05 Structure

**Create:**
- `module05_grpc/project_distributed_compute/internal/engine/engine.go`
- `module05_grpc/project_distributed_compute/internal/engine/engine_test.go`
- `module05_grpc/project_distributed_compute/internal/auth/auth.go`
- `module05_grpc/project_distributed_compute/internal/auth/auth_test.go`
- `module05_grpc/project_distributed_compute/internal/server/service.go`
- `module05_grpc/project_distributed_compute/internal/server/service_test.go`
- `module05_grpc/project_distributed_compute/README.md`
- `module05_grpc/08_grpc_gateway/proto/hello.gw.pb.go` (generated)

**Modify:**
- `module05_grpc/project_distributed_compute/proto/compute.proto`
- `module05_grpc/project_distributed_compute/proto/gen.sh`
- `module05_grpc/project_distributed_compute/server/main.go`
- `module05_grpc/project_distributed_compute/client/main.go`
- `module05_grpc/08_grpc_gateway/proto/hello.proto`
- `module05_grpc/08_grpc_gateway/proto/gen.sh`
- `module05_grpc/08_grpc_gateway/main.go`
- `module05_grpc/README.md`

## Target Module 06 Structure

**Keep as concept demos:**
- `module06_gozero/01_gozero_intro`
- `module06_gozero/02_api_service`
- `module06_gozero/03_rpc_service`
- `module06_gozero/04_etcd_discovery`
- `module06_gozero/05_mysql_cache`
- `module06_gozero/06_message_queue`
- `module06_gozero/07_observability`
- `module06_gozero/08_k8s_deploy`

**Create standard practice project:**
- `module06_gozero/project_ecommerce_standard/README.md`
- `module06_gozero/project_ecommerce_standard/api/order.api`
- `module06_gozero/project_ecommerce_standard/rpc/user/user.proto`
- `module06_gozero/project_ecommerce_standard/rpc/product/product.proto`
- `module06_gozero/project_ecommerce_standard/rpc/order/order.proto`
- `module06_gozero/project_ecommerce_standard/sql/schema.sql`
- `module06_gozero/project_ecommerce_standard/docker-compose.yml`
- `module06_gozero/project_ecommerce_standard/prometheus.yml`
- generated-style directories under:
  - `module06_gozero/project_ecommerce_standard/api/order/internal/...`
  - `module06_gozero/project_ecommerce_standard/rpc/user/internal/...`
  - `module06_gozero/project_ecommerce_standard/rpc/product/internal/...`
  - `module06_gozero/project_ecommerce_standard/rpc/order/internal/...`

---

### Task 1: Refactor Distributed Compute Engine Into Testable Core

**Files:**
- Create: `module05_grpc/project_distributed_compute/internal/engine/engine.go`
- Create: `module05_grpc/project_distributed_compute/internal/engine/engine_test.go`
- Modify: `module05_grpc/project_distributed_compute/server/main.go`

- [ ] **Step 1: Write engine tests**

```go
package engine

import "testing"

func TestCompute(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int64
		operation Operation
		want      float64
		wantError bool
	}{
		{name: "sum", numbers: []int64{1, 2, 3}, operation: OperationSum, want: 6},
		{name: "avg", numbers: []int64{2, 4, 6}, operation: OperationAvg, want: 4},
		{name: "max", numbers: []int64{2, 9, 1}, operation: OperationMax, want: 9},
		{name: "min", numbers: []int64{2, 9, 1}, operation: OperationMin, want: 1},
		{name: "median even", numbers: []int64{3, 1, 4, 2}, operation: OperationMedian, want: 2.5},
		{name: "stddev", numbers: []int64{2, 4, 4, 4, 5, 5, 7, 9}, operation: OperationStddev, want: 2},
		{name: "empty", numbers: nil, operation: OperationSum, wantError: true},
		{name: "unknown", numbers: []int64{1}, operation: Operation("p99"), wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compute(tt.operation, tt.numbers)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %.4f, got %.4f", tt.want, got)
			}
		})
	}
}
```

- [ ] **Step 2: Run test and verify failure**

```bash
go test ./module05_grpc/project_distributed_compute/internal/engine
```

Expected: FAIL because the package does not exist yet.

- [ ] **Step 3: Create engine implementation**

```go
package engine

import (
	"fmt"
	"math"
	"sort"
)

type Operation string

const (
	OperationSum    Operation = "sum"
	OperationAvg    Operation = "avg"
	OperationMax    Operation = "max"
	OperationMin    Operation = "min"
	OperationStddev Operation = "stddev"
	OperationMedian Operation = "median"
)

func ParseOperation(value string) Operation {
	return Operation(value)
}

func Compute(operation Operation, numbers []int64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}

	switch operation {
	case OperationSum:
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		return float64(sum), nil
	case OperationAvg:
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		return float64(sum) / float64(len(numbers)), nil
	case OperationMax:
		max := numbers[0]
		for _, n := range numbers[1:] {
			if n > max {
				max = n
			}
		}
		return float64(max), nil
	case OperationMin:
		min := numbers[0]
		for _, n := range numbers[1:] {
			if n < min {
				min = n
			}
		}
		return float64(min), nil
	case OperationStddev:
		var mean float64
		for _, n := range numbers {
			mean += float64(n)
		}
		mean /= float64(len(numbers))
		var variance float64
		for _, n := range numbers {
			diff := float64(n) - mean
			variance += diff * diff
		}
		return math.Sqrt(variance / float64(len(numbers))), nil
	case OperationMedian:
		sorted := make([]int64, len(numbers))
		copy(sorted, numbers)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
		mid := len(sorted) / 2
		if len(sorted)%2 == 0 {
			return float64(sorted[mid-1]+sorted[mid]) / 2, nil
		}
		return float64(sorted[mid]), nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
```

- [ ] **Step 4: Run test and verify pass**

```bash
go test ./module05_grpc/project_distributed_compute/internal/engine
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add module05_grpc/project_distributed_compute/internal/engine
git commit -m "refactor: extract distributed compute engine"
```

---

### Task 2: Serialize gRPC Stream Sending and Add Service Tests

**Files:**
- Create: `module05_grpc/project_distributed_compute/internal/server/service.go`
- Create: `module05_grpc/project_distributed_compute/internal/server/service_test.go`
- Modify: `module05_grpc/project_distributed_compute/server/main.go`

- [ ] **Step 1: Create service implementation**

```go
package server

import (
	"io"
	"log"
	"sync"

	"iotestgo/module05_grpc/project_distributed_compute/internal/engine"
	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"
)

type Service struct {
	pb.UnimplementedDistributedComputeServer
	WorkerCount int
}

func NewService(workerCount int) *Service {
	if workerCount <= 0 {
		workerCount = 1
	}
	return &Service{WorkerCount: workerCount}
}

func (s *Service) Process(stream pb.DistributedCompute_ProcessServer) error {
	tasksCh := make(chan *pb.ComputeTask, 100)
	resultsCh := make(chan *pb.ComputeResult, 100)

	var workers sync.WaitGroup
	for i := 0; i < s.WorkerCount; i++ {
		workers.Add(1)
		go func(workerID int) {
			defer workers.Done()
			for task := range tasksCh {
				resultsCh <- computeTask(task)
				log.Printf("[Worker-%d] task=%s op=%s", workerID, task.GetTaskId(), task.GetOperation())
			}
		}(i + 1)
	}

	sendDone := make(chan error, 1)
	go func() {
		for result := range resultsCh {
			if err := stream.Send(result); err != nil {
				sendDone <- err
				return
			}
		}
		sendDone <- nil
	}()

	for {
		task, err := stream.Recv()
		if err == io.EOF {
			close(tasksCh)
			workers.Wait()
			close(resultsCh)
			return <-sendDone
		}
		if err != nil {
			close(tasksCh)
			workers.Wait()
			close(resultsCh)
			<-sendDone
			return err
		}
		tasksCh <- task
	}
}

func computeTask(task *pb.ComputeTask) *pb.ComputeResult {
	value, err := engine.Compute(engine.ParseOperation(task.GetOperation()), task.GetNumbers())
	result := &pb.ComputeResult{
		TaskId:    task.GetTaskId(),
		Operation: task.GetOperation(),
	}
	if err != nil {
		result.Status = "error"
		result.Message = err.Error()
		return result
	}
	result.Status = "done"
	result.Value = value
	return result
}
```

- [ ] **Step 2: Update server main**

```go
grpcServer := grpc.NewServer()
pb.RegisterDistributedComputeServer(grpcServer, server.NewService(4))
```

Import:

```go
computeServer "iotestgo/module05_grpc/project_distributed_compute/internal/server"
```

Use alias `computeServer.NewService(4)` to avoid colliding with package `main`.

- [ ] **Step 3: Write service test with fake stream**

Create a fake stream that sends two tasks and expects two results:

```go
package server

import (
	"context"
	"io"
	"testing"

	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"
	"google.golang.org/grpc/metadata"
)

type fakeProcessStream struct {
	pb.DistributedCompute_ProcessServer
	ctx     context.Context
	inputs  []*pb.ComputeTask
	outputs []*pb.ComputeResult
	index   int
}

func (f *fakeProcessStream) Context() context.Context {
	if f.ctx == nil {
		return context.Background()
	}
	return f.ctx
}

func (f *fakeProcessStream) Recv() (*pb.ComputeTask, error) {
	if f.index >= len(f.inputs) {
		return nil, io.EOF
	}
	task := f.inputs[f.index]
	f.index++
	return task, nil
}

func (f *fakeProcessStream) Send(result *pb.ComputeResult) error {
	f.outputs = append(f.outputs, result)
	return nil
}

func (f *fakeProcessStream) SendHeader(metadata.MD) error { return nil }
func (f *fakeProcessStream) SetHeader(metadata.MD) error  { return nil }
func (f *fakeProcessStream) SetTrailer(metadata.MD)       {}

func TestServiceProcess(t *testing.T) {
	stream := &fakeProcessStream{
		inputs: []*pb.ComputeTask{
			{TaskId: "sum", Operation: "sum", Numbers: []int64{1, 2, 3}},
			{TaskId: "bad", Operation: "p99", Numbers: []int64{1, 2, 3}},
		},
	}
	err := NewService(2).Process(stream)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stream.outputs) != 2 {
		t.Fatalf("expected 2 outputs, got %d", len(stream.outputs))
	}
	statuses := map[string]string{}
	for _, result := range stream.outputs {
		statuses[result.GetTaskId()] = result.GetStatus()
	}
	if statuses["sum"] != "done" {
		t.Fatalf("expected sum done, got %q", statuses["sum"])
	}
	if statuses["bad"] != "error" {
		t.Fatalf("expected bad error, got %q", statuses["bad"])
	}
}
```

- [ ] **Step 4: Verify**

```bash
go test ./module05_grpc/project_distributed_compute/internal/...
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add module05_grpc/project_distributed_compute
git commit -m "refactor: make compute grpc service testable"
```

---

### Task 3: Add Compute Authentication Interceptor as Optional Lesson Layer

**Files:**
- Create: `module05_grpc/project_distributed_compute/internal/auth/auth.go`
- Create: `module05_grpc/project_distributed_compute/internal/auth/auth_test.go`
- Modify: `module05_grpc/project_distributed_compute/server/main.go`
- Modify: `module05_grpc/project_distributed_compute/client/main.go`

- [ ] **Step 1: Implement token validation**

```go
package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const DefaultToken = "valid-token-12345"

func UnaryInterceptor(token string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := ValidateIncomingContext(ctx, token); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func StreamInterceptor(token string) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := ValidateIncomingContext(stream.Context(), token); err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func ValidateIncomingContext(ctx context.Context, token string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "missing metadata")
	}
	values := md.Get("authorization")
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "missing authorization")
	}
	got := strings.TrimPrefix(values[0], "Bearer ")
	if got != token {
		return status.Error(codes.PermissionDenied, "invalid token")
	}
	return nil
}
```

- [ ] **Step 2: Test auth**

```go
package auth

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestValidateIncomingContext(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer valid-token-12345"))
	if err := ValidateIncomingContext(ctx, DefaultToken); err != nil {
		t.Fatalf("expected valid token, got %v", err)
	}

	err := ValidateIncomingContext(context.Background(), DefaultToken)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected unauthenticated, got %v", status.Code(err))
	}
}
```

- [ ] **Step 3: Wire auth into compute server**

```go
grpcServer := grpc.NewServer(
	grpc.StreamInterceptor(auth.StreamInterceptor(auth.DefaultToken)),
)
```

- [ ] **Step 4: Add client metadata**

```go
ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer valid-token-12345")
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
```

- [ ] **Step 5: Verify**

```bash
go test ./module05_grpc/project_distributed_compute/internal/...
```

Expected: PASS.

- [ ] **Step 6: Commit**

```bash
git add module05_grpc/project_distributed_compute
git commit -m "feat: add compute stream authentication interceptor"
```

---

### Task 4: Replace Manual Gateway Bridge With Annotation-Based grpc-gateway

**Files:**
- Modify: `module05_grpc/08_grpc_gateway/proto/hello.proto`
- Modify: `module05_grpc/08_grpc_gateway/proto/gen.sh`
- Create: `module05_grpc/08_grpc_gateway/proto/hellopb/hello.gw.pb.go` (generated)
- Modify: `module05_grpc/08_grpc_gateway/main.go`

- [ ] **Step 1: Update proto annotations**

```proto
syntax = "proto3";

package hello;

import "google/api/annotations.proto";

option go_package = "iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb";

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}

service Greeter {
  rpc SayHello(HelloRequest) returns (HelloResponse) {
    option (google.api.http) = {
      post: "/v1/hello"
      body: "*"
    };
  }
}
```

- [ ] **Step 2: Update generation script**

```bash
#!/usr/bin/env bash
set -euo pipefail

protoc -I . \
  -I "$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.29.0/third_party/googleapis" \
  --go_out=hellopb --go_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  --go-grpc_out=hellopb --go-grpc_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  --grpc-gateway_out=hellopb --grpc-gateway_opt=module=iotestgo/module05_grpc/08_grpc_gateway/proto/hellopb \
  hello.proto
```

- [ ] **Step 3: Generate gateway code**

```bash
cd module05_grpc/08_grpc_gateway/proto
chmod +x gen.sh
./gen.sh
```

Expected: `hellopb/hello.gw.pb.go` is created.

- [ ] **Step 4: Replace manual `HandlePath` in main**

Use generated registration:

```go
ctx := context.Background()
conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
if err != nil {
	log.Fatalf("gateway dial failed: %v", err)
}
defer conn.Close()

gwmux := runtime.NewServeMux(
	runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{UseProtoNames: true},
	}),
)
if err := pb.RegisterGreeterHandler(ctx, gwmux, conn); err != nil {
	log.Fatalf("register gateway failed: %v", err)
}
```

Remove manual `io.ReadAll`, `protojson.Unmarshal`, and `gwmux.HandlePath`.

- [ ] **Step 5: Verify gateway manually**

Terminal 1:

```bash
go run ./module05_grpc/08_grpc_gateway
```

Terminal 2:

```bash
curl -s -X POST http://localhost:8080/v1/hello -H 'Content-Type: application/json' -d '{"name":"Gopher"}'
```

Expected:

```json
{"message":"Hello, Gopher! (from gRPC server)"}
```

- [ ] **Step 6: Commit**

```bash
git add module05_grpc/08_grpc_gateway
git commit -m "feat: generate grpc gateway from annotations"
```

---

### Task 5: Create go-zero Standard Ecommerce IDL Layer

**Files:**
- Create: `module06_gozero/project_ecommerce_standard/README.md`
- Create: `module06_gozero/project_ecommerce_standard/api/order.api`
- Create: `module06_gozero/project_ecommerce_standard/rpc/user/user.proto`
- Create: `module06_gozero/project_ecommerce_standard/rpc/product/product.proto`
- Create: `module06_gozero/project_ecommerce_standard/rpc/order/order.proto`

- [ ] **Step 1: Create API definition**

```api
syntax = "v1"

info(
	title: "Ecommerce Order API"
	desc: "本科后端课程 go-zero 标准项目：订单 API 调用订单 RPC"
	author: "iotestgo2"
)

type CreateOrderReq {
	UserId int64 `json:"user_id"`
	Items []OrderItemReq `json:"items"`
}

type OrderItemReq {
	ProductId int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type CreateOrderResp {
	OrderId string `json:"order_id"`
	Status string `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}

type GetOrderReq {
	OrderId string `path:"order_id"`
}

type GetOrderResp {
	OrderId string `json:"order_id"`
	UserId int64 `json:"user_id"`
	Status string `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}

service order-api {
	@handler CreateOrder
	post /api/v1/orders (CreateOrderReq) returns (CreateOrderResp)

	@handler GetOrder
	get /api/v1/orders/:order_id (GetOrderReq) returns (GetOrderResp)
}
```

- [ ] **Step 2: Create User RPC proto**

```proto
syntax = "proto3";

package user;

option go_package = "./userpb";

message GetUserRequest {
  int64 user_id = 1;
}

message GetUserResponse {
  int64 user_id = 1;
  string username = 2;
  string email = 3;
  int32 status = 4;
}

message CheckUserStatusRequest {
  int64 user_id = 1;
}

message CheckUserStatusResponse {
  bool valid = 1;
  string reason = 2;
}

service User {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc CheckUserStatus(CheckUserStatusRequest) returns (CheckUserStatusResponse);
}
```

- [ ] **Step 3: Create Product RPC proto**

```proto
syntax = "proto3";

package product;

option go_package = "./productpb";

message GetProductRequest {
  int64 product_id = 1;
}

message GetProductResponse {
  int64 product_id = 1;
  string name = 2;
  int64 stock = 3;
  double price = 4;
}

message ReserveStockRequest {
  int64 product_id = 1;
  int64 quantity = 2;
}

message ReserveStockResponse {
  bool success = 1;
  string reason = 2;
}

service Product {
  rpc GetProduct(GetProductRequest) returns (GetProductResponse);
  rpc ReserveStock(ReserveStockRequest) returns (ReserveStockResponse);
}
```

- [ ] **Step 4: Create Order RPC proto**

```proto
syntax = "proto3";

package order;

option go_package = "./orderpb";

message OrderItem {
  int64 product_id = 1;
  int64 quantity = 2;
}

message CreateOrderRequest {
  int64 user_id = 1;
  repeated OrderItem items = 2;
}

message CreateOrderResponse {
  string order_id = 1;
  string status = 2;
  double total_amount = 3;
}

message GetOrderRequest {
  string order_id = 1;
}

message GetOrderResponse {
  string order_id = 1;
  int64 user_id = 2;
  string status = 3;
  double total_amount = 4;
}

service Order {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
  rpc GetOrder(GetOrderRequest) returns (GetOrderResponse);
}
```

- [ ] **Step 5: Create project README**

```markdown
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
```

- [ ] **Step 6: Commit**

```bash
git add module06_gozero/project_ecommerce_standard
git commit -m "feat: add go-zero ecommerce standard idl"
```

---

### Task 6: Generate go-zero Standard Project and Implement Real Call Chain

**Files:**
- Generated and modify under `module06_gozero/project_ecommerce_standard/api/order`
- Generated and modify under `module06_gozero/project_ecommerce_standard/rpc/user`
- Generated and modify under `module06_gozero/project_ecommerce_standard/rpc/product`
- Generated and modify under `module06_gozero/project_ecommerce_standard/rpc/order`

- [ ] **Step 1: Generate project skeletons**

```bash
cd module06_gozero/project_ecommerce_standard
goctl api go -api api/order.api -dir api/order
goctl rpc protoc rpc/user/user.proto --go_out=rpc/user --go-grpc_out=rpc/user --zrpc_out=rpc/user
goctl rpc protoc rpc/product/product.proto --go_out=rpc/product --go-grpc_out=rpc/product --zrpc_out=rpc/product
goctl rpc protoc rpc/order/order.proto --go_out=rpc/order --go-grpc_out=rpc/order --zrpc_out=rpc/order
```

Expected: generated `etc`, `internal/config`, `internal/logic`, `internal/server`, `internal/svc`, and client packages.

- [ ] **Step 2: Configure ports and Etcd keys**

Use:

```yaml
Name: user.rpc
ListenOn: 0.0.0.0:9101
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: user.rpc
```

```yaml
Name: product.rpc
ListenOn: 0.0.0.0:9102
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: product.rpc
```

```yaml
Name: order.rpc
ListenOn: 0.0.0.0:9103
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: order.rpc
UserRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: user.rpc
ProductRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: product.rpc
```

```yaml
Name: order-api
Host: 0.0.0.0
Port: 8890
OrderRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: order.rpc
```

- [ ] **Step 3: Implement User logic with in-memory repository**

In `getuserlogic.go`, return:

```go
users := map[int64]*user.GetUserResponse{
	1: {UserId: 1, Username: "gopher", Email: "gopher@example.com", Status: 1},
	2: {UserId: 2, Username: "alice", Email: "alice@example.com", Status: 1},
	3: {UserId: 3, Username: "disabled", Email: "disabled@example.com", Status: 2},
}
```

For missing user, return `status.Error(codes.NotFound, "user not found")`.

- [ ] **Step 4: Implement Product logic with stock checks**

Use:

```go
products := map[int64]*product.GetProductResponse{
	101: {ProductId: 101, Name: "Go Backend Book", Stock: 10, Price: 59.9},
	102: {ProductId: 102, Name: "Cloud Native Notebook", Stock: 5, Price: 29.9},
}
```

`ReserveStock` returns `Success=false` if quantity exceeds stock.

- [ ] **Step 5: Implement Order RPC logic**

Flow:

1. Call `UserRpc.CheckUserStatus`.
2. For each item, call `ProductRpc.GetProduct`.
3. For each item, call `ProductRpc.ReserveStock`.
4. Calculate `totalAmount`.
5. Save order to an in-memory map in `ServiceContext`.

Use deterministic order id:

```go
orderID := fmt.Sprintf("ORD-%d", time.Now().UnixNano())
```

- [ ] **Step 6: Implement Order API logic**

`CreateOrderLogic` maps API request to `orderclient.CreateOrderRequest`, calls `OrderRpc.CreateOrder`, and maps the response.

`GetOrderLogic` calls `OrderRpc.GetOrder` and maps the response.

- [ ] **Step 7: Verify manually**

Start infra:

```bash
cd module06_gozero/project_ecommerce_standard
docker compose up -d etcd
```

Start services in order:

```bash
go run ./rpc/user/user.go -f rpc/user/etc/user.yaml
go run ./rpc/product/product.go -f rpc/product/etc/product.yaml
go run ./rpc/order/order.go -f rpc/order/etc/order.yaml
go run ./api/order/order.go -f api/order/etc/order-api.yaml
```

Call:

```bash
curl -s -X POST http://localhost:8890/api/v1/orders \
  -H 'Content-Type: application/json' \
  -d '{"user_id":1,"items":[{"product_id":101,"quantity":2}]}'
```

Expected JSON contains:

```json
{"order_id":"ORD-","status":"created","total_amount":119.8}
```

The exact `order_id` suffix can vary.

- [ ] **Step 8: Commit**

```bash
git add module06_gozero/project_ecommerce_standard
git commit -m "feat: implement standard go-zero ecommerce call chain"
```

---

### Task 7: Fix Observability and Docker Compose for go-zero Standard Project

**Files:**
- Create: `module06_gozero/project_ecommerce_standard/docker-compose.yml`
- Create: `module06_gozero/project_ecommerce_standard/prometheus.yml`
- Modify generated service configs to expose metrics if needed.

- [ ] **Step 1: Create compose file**

```yaml
version: "3.8"

services:
  etcd:
    image: quay.io/coreos/etcd:v3.5.9
    command:
      - /usr/local/bin/etcd
      - --name=etcd0
      - --data-dir=/etcd-data
      - --listen-client-urls=http://0.0.0.0:2379
      - --advertise-client-urls=http://127.0.0.1:2379
    ports:
      - "2379:2379"

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root123
      MYSQL_DATABASE: ecommerce
    ports:
      - "3306:3306"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  prometheus:
    image: prom/prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
```

- [ ] **Step 2: Create prometheus targets matching actual ports**

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: "order-api"
    static_configs:
      - targets: ["host.docker.internal:8890"]
  - job_name: "user-rpc"
    static_configs:
      - targets: ["host.docker.internal:9101"]
  - job_name: "product-rpc"
    static_configs:
      - targets: ["host.docker.internal:9102"]
  - job_name: "order-rpc"
    static_configs:
      - targets: ["host.docker.internal:9103"]
```

- [ ] **Step 3: Add README verification commands**

```bash
docker compose up -d
curl http://localhost:9090/targets
```

- [ ] **Step 4: Commit**

```bash
git add module06_gozero/project_ecommerce_standard/docker-compose.yml module06_gozero/project_ecommerce_standard/prometheus.yml module06_gozero/project_ecommerce_standard/README.md
git commit -m "chore: add ecommerce standard infrastructure"
```

---

### Task 8: Add Module 05 and 06 README Teaching Maps

**Files:**
- Modify: `module05_grpc/README.md`
- Modify: `module06_gozero/README.md`

- [ ] **Step 1: Add Module 05 teaching map**

```markdown
## 建议讲课顺序

1. `01_protobuf_basics`: 只讲 IDL 和 message。
2. `02_codegen`: 讲 `protoc` 到 `*.pb.go` / `*_grpc.pb.go`。
3. `03_unary_rpc`: 讲最小 request/response。
4. `04_streaming_rpc`: 讲三种 stream。
5. `05_interceptors`: 讲横切能力。
6. `06_metadata_auth`: 讲 metadata 和认证。
7. `07_error_handling`: 讲 status code。
8. `08_grpc_gateway`: 讲 annotation-based HTTP bridge。
9. `project_distributed_compute`: 讲综合项目，重点是双向流 + worker pool + auth interceptor + testable engine。

## 生成代码规则

不要手改 `*.pb.go`、`*_grpc.pb.go`、`*.gw.pb.go`。修改 `.proto` 后运行对应 `gen.sh`。
```

- [ ] **Step 2: Add Module 06 two-track map**

```markdown
## 两条教学线

### Track A: 概念演示

`01_gozero_intro` 到 `08_k8s_deploy` 使用小文件解释概念，适合第一遍课堂讲解。

### Track B: 标准工程

`project_ecommerce_standard` 使用 goctl 风格目录和真实 API -> RPC -> RPC 调用链，适合作为综合实践或期末项目基础。
```

- [ ] **Step 3: Commit**

```bash
git add module05_grpc/README.md module06_gozero/README.md
git commit -m "docs: clarify grpc and go-zero teaching tracks"
```

---

## Self-Review

- **Spec coverage:** Covers deeper gRPC analysis, compute refactor, auth, gateway generation, go-zero standard project, true microservice call chain, observability correction, and README teaching maps.
- **占位检查:** No unresolved marker sections remain.
- **Type consistency:** `engine.Operation`, `server.Service`, and generated go-zero project names are consistent across tasks.
- **Source alignment:** Direction follows official gRPC/go-zero/grpc-gateway docs: interceptors for cross-cutting concerns, metadata for auth side-channel, grpc-gateway via annotations, go-zero via goctl-generated API/RPC structures.
