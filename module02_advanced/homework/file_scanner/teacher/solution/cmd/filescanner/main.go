package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"module02filescanner/scanner"
)

func main() {
	root := flag.String("root", ".", "要扫描的根目录")
	text := flag.String("text", "", "要匹配的文本")
	workers := flag.Int("workers", 4, "并发 worker 数量")
	timeout := flag.Duration("timeout", 10*time.Second, "扫描超时")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	result, err := scanner.Scan(ctx, *root, *text, *workers)
	if err != nil {
		fmt.Fprintln(os.Stderr, "scan:", err)
		os.Exit(1)
	}
	fmt.Printf("files=%d matching_lines=%d\n", result.Files, result.MatchingLines)
}
