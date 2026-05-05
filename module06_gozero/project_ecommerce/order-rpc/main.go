// 电商项目 - Order RPC 服务（内部 gRPC 服务）
//
// 职责：订单的业务逻辑和数据持久化
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 电商项目：Order RPC 服务 ===")
	fmt.Println()

	fmt.Println("--- 订单 DDL (order.sql) ---")
	fmt.Println(`  CREATE TABLE orders (
      id BIGINT AUTO_INCREMENT PRIMARY KEY,
      order_no VARCHAR(32) NOT NULL UNIQUE COMMENT '订单号',
      user_id BIGINT NOT NULL,
      total_amount DECIMAL(10,2) NOT NULL COMMENT '总金额',
      status TINYINT DEFAULT 1 COMMENT '1=待支付 2=已支付 3=已发货 4=已完成 5=已取消',
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      INDEX idx_user_id (user_id),
      INDEX idx_status (status)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

  CREATE TABLE order_items (
      id BIGINT AUTO_INCREMENT PRIMARY KEY,
      order_id BIGINT NOT NULL,
      product_id BIGINT NOT NULL,
      quantity INT NOT NULL,
      price DECIMAL(10,2) NOT NULL,
      INDEX idx_order_id (order_id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	fmt.Println()

	fmt.Println("--- Logic 层核心 ---")
	fmt.Println(`  func (l *CreateOrderLogic) CreateOrder(in *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
      // 1. 生成订单号
      orderNo := generateOrderNo()

      // 2. 写入订单表（事务）
      err := l.svcCtx.OrderModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
          order := &model.Order{
              OrderNo:     orderNo,
              UserId:      in.UserId,
              TotalAmount: in.Amount,
              Status:      1, // 待支付
          }
          result, err := l.svcCtx.OrderModel.Insert(ctx, order)
          if err != nil { return err }
          orderId, _ := result.LastInsertId()

          // 写入订单项
          for _, item := range in.Items {
              oi := &model.OrderItem{
                  OrderId:   orderId,
                  ProductId: item.ProductId,
                  Quantity:  item.Quantity,
                  Price:     item.Price,
              }
              if _, err := l.svcCtx.OrderItemModel.Insert(ctx, oi); err != nil {
                  return err
              }
          }
          return nil
      })
      if err != nil {
          return nil, status.Error(codes.Internal, err.Error())
      }

      // 3. 发送消息通知其他服务（如通知服务）
      l.pushOrderCreatedEvent(orderNo)

      return &orderpb.CreateOrderResponse{OrderId: orderNo, Status: "created"}, nil
  }`)

	fmt.Println()
	fmt.Println("--- 关键设计点 ---")
	fmt.Println("  1. 事务保证订单+订单项的一致性写入")
	fmt.Println("  2. 订单号生成策略（雪花算法 or UUID）")
	fmt.Println("  3. 异步事件通知（不阻塞主流程）")
	fmt.Println("  4. 幂等性：重复创建返回已有订单而非报错")

	_ = fmt.Sprint
}
