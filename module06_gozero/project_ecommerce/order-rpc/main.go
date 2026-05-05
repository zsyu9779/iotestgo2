// 电商项目 - Order RPC 服务
// 职责：订单 CRUD，调用 UserRpc 校验用户，事务写订单+订单项
//
// 启动：go run order-rpc/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// 订单模型
type Order struct {
	OrderID     string
	UserID      int64
	Items       []OrderItem
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
}

type OrderItem struct {
	ProductID int64
	Quantity  int32
	Price     float64
}

// Request/Response
type CreateOrderRequest struct {
	UserID int64
	Items  []OrderItemPb
	Amount float64
}
type OrderItemPb struct {
	ProductId int64
	Quantity  int32
}
type CreateOrderResponse struct {
	OrderID string
	Status  string
}

type GetOrderRequest struct {
	OrderID string
}
type GetOrderResponse struct {
	OrderID     string
	UserID      int64
	TotalAmount float64
	Status      string
	CreatedAt   string
}

// OrderServiceServer 接口
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
}

type orderServer struct {
	mu      sync.RWMutex
	orders  map[string]*Order
	nextID  int64
	userRpc *grpc.ClientConn // 调用 UserRpc
}

func newOrderServer(userRpcConn *grpc.ClientConn) *orderServer {
	return &orderServer{
		orders:  make(map[string]*Order),
		nextID:  1000,
		userRpc: userRpcConn,
	}
}

func (s *orderServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	// 1. 校验用户：调用 UserRpc（微服务间通信）
	userOK := s.checkUser(req.UserID)
	if !userOK {
		return nil, status.Errorf(codes.FailedPrecondition, "user %d is not valid or disabled", req.UserID)
	}

	// 2. 生成订单
	s.mu.Lock()
	orderID := fmt.Sprintf("ORD-%d", s.nextID)
	s.nextID++

	items := make([]OrderItem, len(req.Items))
	for i, it := range req.Items {
		items[i] = OrderItem{ProductID: it.ProductId, Quantity: it.Quantity, Price: 0}
	}

	order := &Order{
		OrderID:     orderID,
		UserID:      req.UserID,
		Items:       items,
		TotalAmount: req.Amount,
		Status:      "created",
		CreatedAt:   time.Now(),
	}
	s.orders[orderID] = order
	s.mu.Unlock()

	log.Printf("[OrderRpc] 订单创建: %s, userId=%d, amount=%.2f, items=%d", orderID, req.UserID, req.Amount, len(req.Items))
	return &CreateOrderResponse{OrderID: orderID, Status: "created"}, nil
}

func (s *orderServer) GetOrder(ctx context.Context, req *GetOrderRequest) (*GetOrderResponse, error) {
	s.mu.RLock()
	order, ok := s.orders[req.OrderID]
	s.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "order %s not found", req.OrderID)
	}
	return &GetOrderResponse{
		OrderID:     order.OrderID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt.Format(time.RFC3339),
	}, nil
}

// checkUser 调用 UserRpc 校验用户
func (s *orderServer) checkUser(userID int64) bool {
	if s.userRpc == nil {
		return true // 无 UserRpc 连接时默认通过
	}
	// 简化：直接用 conn 发送 unary RPC
	// 实际项目使用生成的 pb client
	return true // 教学简化
}

// ========== gRPC Service Descriptor ==========

var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
				req := &CreateOrderRequest{}
				if err := dec(req); err != nil {
					return nil, err
				}
				return srv.(OrderServiceServer).CreateOrder(ctx, req)
			},
		},
		{
			MethodName: "GetOrder",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
				req := &GetOrderRequest{}
				if err := dec(req); err != nil {
					return nil, err
				}
				return srv.(OrderServiceServer).GetOrder(ctx, req)
			},
		},
	},
}

func main() {
	// 连接 UserRpc 服务
	userRpcConn, err := grpc.NewClient("localhost:9091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("警告: 无法连接 UserRpc: %v（将跳过用户校验）", err)
		userRpcConn = nil
	}

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	s.RegisterService(&OrderService_ServiceDesc, newOrderServer(userRpcConn))
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down OrderRpc...")
		s.GracefulStop()
	}()

	fmt.Println("=== Order RPC 服务 已启动 ===")
	fmt.Println("  gRPC 端口: :9092")
	fmt.Println("  接口: CreateOrder, GetOrder")
	fmt.Println("  依赖: UserRpc (:9091)")
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	_ = context.Background
}
