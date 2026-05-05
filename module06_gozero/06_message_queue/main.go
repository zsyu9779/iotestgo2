// 06 消息队列：用 Go Channel 模拟 Producer/Consumer 模式
//
// 启动：go run main.go
//
// 消息队列核心概念：
//   Producer（生产者） → Broker（消息队列/Channel） → Consumer（消费者）
//   - 解耦：订单服务不需要直接调用库存服务
//   - 削峰：高峰期消息堆积，消费者按自己速度处理
//   - 异步：发送后立即返回，不等待处理
package main

import (
	"fmt"
	"sync"
	"time"
)

// Message 消息结构
type Message struct {
	ID        string
	Topic     string
	Data      interface{}
	Timestamp time.Time
}

// Producer 生产者
type Producer struct {
	name  string
	queue chan<- Message
}

func (p *Producer) Send(topic string, data interface{}) {
	msg := Message{
		ID:        fmt.Sprintf("%s-%d", p.name, time.Now().UnixNano()),
		Topic:     topic,
		Data:      data,
		Timestamp: time.Now(),
	}
	p.queue <- msg
	fmt.Printf("  [%s] 生产消息 → topic=%s, id=%s, data=%v\n", p.name, topic, msg.ID, data)
}

// Consumer 消费者
type Consumer struct {
	name  string
	queue <-chan Message
	wg    *sync.WaitGroup
}

func (c *Consumer) Start() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for msg := range c.queue {
			fmt.Printf("  [%s] 消费消息 ← topic=%s, id=%s, data=%v\n", c.name, msg.Topic, msg.ID, msg.Data)
			// 模拟处理耗时
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("  [%s] 消费完成\n", c.name)
	}()
}

func main() {
	fmt.Println("=== 消息队列演示（Go Channel 版） ===")
	fmt.Println()

	// 创建消息队列（缓冲 channel，容量 10）
	queue := make(chan Message, 10)
	var wg sync.WaitGroup

	// ========== 1. 单生产者 → 单消费者 ==========
	fmt.Println("--- 1. 单生产者 → 单消费者 ---")
	p1 := &Producer{name: "order-service", queue: queue}
	c1 := &Consumer{name: "inventory-consumer", queue: queue, wg: &wg}
	c1.Start()

	p1.Send("order.created", "订单 #1001 已创建")
	p1.Send("order.created", "订单 #1002 已创建")
	p1.Send("order.created", "订单 #1003 已创建")

	time.Sleep(2 * time.Second)

	// ========== 2. 演示新的消费者 — 每个消息被一个消费者消费 ==========
	fmt.Println()
	fmt.Println("--- 2. 多消费者竞争消费 ---")
	// 注意：Channel 模式是点对点，消息被消费后其他消费者无法再消费
	// 实际 Kafka 是按 partition 分配消费者，每个 group 消费一次
	// 这里演示多个 goroutine 竞争消费同一个 channel

	queue2 := make(chan Message, 10)
	c2a := &Consumer{name: "consumer-a", queue: queue2, wg: &wg}
	c2b := &Consumer{name: "consumer-b", queue: queue2, wg: &wg}
	p2 := &Producer{name: "order-service", queue: queue2}

	c2a.Start()
	c2b.Start()

	for i := 1; i <= 4; i++ {
		p2.Send("order.created", fmt.Sprintf("订单 #%d", 2000+i))
	}
	close(queue2)
	wg.Wait()
	fmt.Println()

	fmt.Println("=== 消息队列的作用 ===")
	fmt.Println("  1. 解耦：订单服务不需要知道库存服务的存在")
	fmt.Println("  2. 削峰：消息缓冲在 Queue 中，消费者按自己速度处理")
	fmt.Println("  3. 异步：生产者发送后立即返回，不等待消费结果")
	fmt.Println("  4. 可靠性：Kafka 多副本 + ISR 保证消息不丢")
	fmt.Println()
	fmt.Println("=== go-zero 的 go-queue 抽象 ===")
	fmt.Println("  go-queue 封装了 Producer / Consumer 接口")
	fmt.Println("  支持 Kafka、Beanstalk、NSQ、RabbitMQ")
	fmt.Println("  内置重试、死信队列、消息去重")
	fmt.Println()
	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: Spring Kafka / RocketMQTemplate + @KafkaListener")
	fmt.Println("  go-zero: go-queue + goctl 生成 Consumer/Producer 模板")
}
