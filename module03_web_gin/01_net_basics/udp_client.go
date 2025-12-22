package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// RunUDPClient starts a simple interactive UDP client.
// It connects to the server at serverAddress (e.g., "127.0.0.1:9981").
func RunUDPClient(serverAddress string) error {
	// 1. 准备目标地址
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		return fmt.Errorf("error resolving address: %v", err)
	}

	// 2. 建立连接 (DialUDP)
	// UDP 是无连接的，DialUDP 只是初始化一个 socket 并记录目标地址
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return fmt.Errorf("error dialing: %v", err)
	}
	defer conn.Close()

	fmt.Printf("UDP Client started. Target: %s\n", serverAddr)
	fmt.Println("Type message and press Enter (type 'exit' to quit):")

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) == "exit" {
			break
		}

		// 3. 发送数据
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing:", err)
			continue
		}

		// 4. 接收响应
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		fmt.Printf("Response from <%s>: %s\n", remoteAddr, string(buffer[:n]))
		fmt.Print("> ")
	}
	return nil
}
