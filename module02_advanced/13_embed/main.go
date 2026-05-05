// embed 包演示：编译时嵌入静态资源到二进制
//
// 使用 //go:embed 指令嵌入文件/目录
// 程序运行时直接从内存读取，无需外部文件依赖
//
// 运行：go run .
// 注意：需要 data/ 目录下有文件，否则编译失败
package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed data/hello.txt
var helloContent string // 嵌入单个文件为字符串

//go:embed data/hello.txt
var helloBytes []byte // 嵌入单个文件为字节切片

//go:embed data/*.txt
var txtFiles embed.FS // 嵌入多个文件（匹配 pattern）

//go:embed data
var dataDir embed.FS // 嵌入整个目录

func main() {
	fmt.Println("=== embed 内嵌资源 ===")
	fmt.Println()

	// 1. 嵌入为 string
	fmt.Println("--- 1. 嵌入为 string ---")
	fmt.Printf("  hello.txt: %s\n", helloContent)

	// 2. 嵌入为 []byte
	fmt.Println()
	fmt.Println("--- 2. 嵌入为 []byte ---")
	fmt.Printf("  hello.txt bytes: %s\n", string(helloBytes))

	// 3. 嵌入多个文件（embed.FS）
	fmt.Println()
	fmt.Println("--- 3. 嵌入多个文件 (embed.FS) ---")
	entries, err := txtFiles.ReadDir(".")
	if err != nil {
		fmt.Println("  需要创建 data/hello.txt 文件")
	} else {
		for _, e := range entries {
			fmt.Printf("  %s\n", e.Name())
		}
	}

	// 4. 嵌入整个目录
	fmt.Println()
	fmt.Println("--- 4. 嵌入目录 ---")
	if entries, err := fs.ReadDir(dataDir, "data"); err == nil {
		for _, e := range entries {
			fmt.Printf("  %s\n", e.Name())
		}
	}

	fmt.Println()
	fmt.Println("--- 使用场景 ---")
	fmt.Println("  1. 嵌入 HTML 模板文件（Web 应用）")
	fmt.Println("  2. 嵌入配置文件（默认配置）")
	fmt.Println("  3. 嵌入静态资源（CSS/JS/图片）")
	fmt.Println("  4. 嵌入 SQL 迁移脚本")
	fmt.Println("  5. 单文件部署（二进制包含所有依赖）")
	fmt.Println()
	fmt.Println("--- 约束 ---")
	fmt.Println("  1. //go:embed 后必须紧跟变量声明")
	fmt.Println("  2. 路径相对于当前目录（不能向上 ../）")
	fmt.Println("  3. 不能嵌入符号链接指向的目录")
	fmt.Println("  4. 空目录不会被嵌入")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: ClassLoader.getResourceAsStream() + 打包进 JAR")
	fmt.Println("  Go:   embed 编译时嵌入，性能更好，更安全")
}

// 为了编译通过（无 data/ 目录时），取消注释下面行：
// 实际使用时需要创建 data/hello.txt
