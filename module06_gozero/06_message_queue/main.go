// 06 消息队列：go-queue 抽象 + Kafka 集成
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 06 消息队列 ===")
	fmt.Println()

	fmt.Println("--- go-queue 设计理念 ---")
	fmt.Println("  go-queue 是 go-zero 的消息队列抽象层：")
	fmt.Println("  1. 统一接口：Producer / Consumer 接口不绑定具体 MQ")
	fmt.Println("  2. 支持多种 MQ：Kafka、Beanstalk、NSQ、RabbitMQ")
	fmt.Println("  3. 内置重试、死信队列、消息去重")
	fmt.Println()

	fmt.Println("--- Kafka 集成示例 ---")
	fmt.Println()
	fmt.Println("配置文件 etc/order-api.yaml：")
	fmt.Println(`  Kafka:
    Addrs:
      - localhost:9092
    Topic: order.created`)
	fmt.Println()

	fmt.Println("生产者（订单创建后发送消息）：")
	fmt.Println(`  func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (*types.CreateOrderResponse, error) {
      // 1. 创建订单（写数据库）
      order, err := l.svcCtx.OrderModel.Insert(l.ctx, &model.Order{...})
      if err != nil { return nil, err }

      // 2. 发送消息到 Kafka
      msg := map[string]interface{}{
          "order_id": order.Id,
          "user_id":  req.UserId,
          "amount":   req.Amount,
          "time":     time.Now().Unix(),
      }
      msgBytes, _ := json.Marshal(msg)
      if err := l.svcCtx.KafkaPusher.Push(string(msgBytes)); err != nil {
          logx.Errorf("kafka push error: %v", err)
          // 消息发送失败不影响主流程（异步补偿）
      }

      return &types.CreateOrderResponse{OrderId: order.Id}, nil
  }`)
	fmt.Println()

	fmt.Println("消费者（库存服务消费订单消息）：")
	fmt.Println(`  func ConsumeOrderCreated(ctx context.Context, svcCtx *svc.ServiceContext) error {
      for {
          msg, err := svcCtx.KafkaConsumer.Consume(ctx)
          if err != nil { return err }

          var orderMsg OrderMessage
          json.Unmarshal(msg, &orderMsg)

          // 扣减库存
          if err := svcCtx.InventoryModel.Deduct(ctx, orderMsg.ProductId, orderMsg.Quantity); err != nil {
              // 处理失败：重试 or 死信队列 or 人工补偿
              logx.Errorf("deduct inventory failed: %v", err)
          }

          msg.Ack()  // 确认消费
      }
  }`)
	fmt.Println()

	fmt.Println("--- 消息队列在微服务中的作用 ---")
	fmt.Println("  1. 解耦：订单服务不需要直接调用库存服务")
	fmt.Println("  2. 削峰：高峰期消息堆积在 Kafka，消费者按自己速度处理")
	fmt.Println("  3. 异步：订单创建后立即返回，后续流程异步完成")
	fmt.Println("  4. 最终一致性：通过消息 + 本地事务表保证")
	fmt.Println()

	fmt.Println("--- 消息可靠性保证 ---")
	fmt.Println("  1. 生产端：本地事务表 + 定时任务补偿（Outbox 模式）")
	fmt.Println("  2. Broker 端：Kafka 多副本 + ISR 机制")
	fmt.Println("  3. 消费端：手动 Ack + 幂等消费（唯一 ID 去重）")
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: Spring Kafka / RocketMQTemplate")
	fmt.Println("  go-zero: go-queue 抽象 + Kafka 插件")
	fmt.Println("  go-zero 的 go-queue 接口更简洁，不依赖 Spring 的自动配置")

	_ = time.Now
}
