package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 本示例介绍 Go 的底层文件 I/O，对应原项目中的 myio/myfile 部分。

func main() {
	// 准备文件操作目录。
	dir := "test_files"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // 教学示例结束后清理临时目录。

	filename := dir + "/example.txt"
	copyFilename := dir + "/example_copy.txt"

	fmt.Println("=== 1. Basic Write ===")
	basicWrite(filename)

	fmt.Println("\n=== 2. Basic Read ===")
	basicRead(filename)

	fmt.Println("\n=== 3. File Copy (Low-level) ===")
	fileCopy(filename, copyFilename)

	fmt.Println("\n=== 4. Buffered I/O (bufio) ===")
	bufferedCopy(filename, dir+"/example_buf_copy.txt")

	fmt.Println("\n=== 5. Seek and ReadAt ===")
	seekAndRead(filename)
}

func basicWrite(filename string) {
	// os.Create 创建文件，已存在时会截断文件。
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}
	defer f.Close()

	// 写入字符串。
	f.WriteString("Hello, World!\n")
	// 写入字节。
	f.Write([]byte("This is a low-level file I/O demo.\n"))

	fmt.Printf("Wrote content to %s\n", filename)
}

func basicRead(filename string) {
	// os.Open 以只读方式打开文件。
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer f.Close()

	// 读入缓冲区。
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			// 直接写入标准输出。
			os.Stdout.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
	}
}

func fileCopy(srcName, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// 使用缓冲区复制文件。
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := dst.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	fmt.Printf("Copied %s to %s\n", srcName, dstName)
}

func bufferedCopy(srcName, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// 使用 bufio 完成缓冲 I/O。
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)

	// io.Copy 内部也会使用缓冲，这里用 bufio 类型展示接口组合。
	n, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	writer.Flush() // 不要忘记刷新写缓冲。
	fmt.Printf("Buffered copied %d bytes to %s\n", n, dstName)
}

func seekAndRead(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 跳到偏移量 7，跳过 "Hello, "。
	_, err = f.Seek(7, io.SeekStart)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 5)
	io.ReadAtLeast(f, buf, 5)
	fmt.Printf("Read after seek(7): %s\n", string(buf)) // 预期读取到 "World"。
}
