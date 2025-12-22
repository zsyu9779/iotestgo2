package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// 演示 os/exec 的使用
func demoExec() {
	fmt.Println("--- 1. Exec Command Demo ---")

	// 1. 简单的命令执行
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		fmt.Printf("Error running date: %v\n", err)
	} else {
		fmt.Printf("> date\n%s", string(dateOut))
	}

	// 2. 带参数的命令
	echoCmd := exec.Command("echo", "Hello", "System", "Programming")
	echoOut, err := echoCmd.Output()
	if err != nil {
		fmt.Printf("Error running echo: %v\n", err)
	} else {
		fmt.Printf("> echo ...\n%s", string(echoOut))
	}

	// 3. 管道操作 (Grep)
	// 模拟: echo "hello grep\ngoodbye grep" | grep "hello"
	grepCmd := exec.Command("grep", "hello")
	
	// 获取 stdin 和 stdout 管道
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	
	grepCmd.Start() // 开始执行
	
	// 写入数据到 grep 的 stdin
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close() // 必须关闭 stdin，告诉 grep 输入结束
	
	// 读取 grep 的输出
	var outBuf bytes.Buffer
	outBuf.ReadFrom(grepOut)
	
	grepCmd.Wait() // 等待命令结束
	
	fmt.Printf("> grep hello\n%s", outBuf.String())
}

// 演示 os/signal 的使用
func demoSignals() {
	fmt.Println("\n--- 2. Signal Handling Demo ---")
	fmt.Println("Press Ctrl+C to trigger SIGINT...")

	// 创建一个接收信号的通道
	sigs := make(chan os.Signal, 1)
	
	// 注册要接收的信号
	// syscall.SIGINT: Ctrl+C
	// syscall.SIGTERM: 终止信号 (kill)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 使用 Context 控制退出
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 开启一个 goroutine 监听信号
	go func() {
		sig := <-sigs
		fmt.Printf("\nReceived signal: %v\n", sig)
		cancel() // 收到信号后取消 context
	}()

	// 模拟主程序运行
	fmt.Println("Program is running (will exit in 10s or on signal)...")
	select {
	case <-ctx.Done():
		fmt.Println("Program exiting...")
	}
}

func main() {
	demoExec()
	demoSignals()
}
