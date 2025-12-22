package main

import (
	"fmt"
	"net"
)

// RunUDPServer starts a simple UDP echo server.
// It listens on the specified address (e.g., "127.0.0.1:9981") and echoes back any received data.
// Note: This function blocks indefinitely. In a real app, you'd want cancellation/context support.
func RunUDPServer(address string) error {
	// 1. 解析地址
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("error resolving address: %v", err)
	}

	// 2. 监听端口
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("error listening: %v", err)
	}
	defer conn.Close()

	fmt.Printf("UDP Server listening on %s\n", conn.LocalAddr().String())

	// 3. 循环读取数据
	buffer := make([]byte, 1024)
	for {
		// ReadFromUDP 返回读取的字节数、发送方的地址
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received from <%s>: %s\n", remoteAddr, message)

		// 4. 回复数据
		response := fmt.Sprintf("Echo: %s", message)
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			fmt.Println("Error writing:", err)
		}
	}
}
