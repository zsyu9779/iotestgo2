package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 本示例介绍 Go 的底层文件 I/O，对应原项目中的 myio/myfile 部分。

func main() {
	// 使用系统临时目录，避免硬编码路径和污染工作区。
	dir, err := os.MkdirTemp("", "module02-file-io-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // 教学示例结束后清理临时目录。

	filename := filepath.Join(dir, "example.txt")
	copyFilename := filepath.Join(dir, "example_copy.txt")

	fmt.Println("=== 1. Basic Write ===")
	basicWrite(filename)

	fmt.Println("\n=== 2. Basic Read ===")
	basicRead(filename)

	fmt.Println("\n=== 3. File Copy (Low-level) ===")
	fileCopy(filename, copyFilename)

	fmt.Println("\n=== 4. Buffered I/O (bufio) ===")
	bufferedCopy(filename, filepath.Join(dir, "example_buf_copy.txt"))

	fmt.Println("\n=== 5. Seek and ReadAt ===")
	seekAndRead(filename)

	fmt.Println("\n=== 6. Scanner limit and atomic replace ===")
	showScannerLimit()
	atomicReplace(filepath.Join(dir, "settings.txt"), []byte("version=2\n"))
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
	if _, err := io.WriteString(f, "Hello, World!\n"); err != nil {
		panic(err)
	}
	// 写入字节。
	if err := writeAll(f, []byte("This is a low-level file I/O demo.\n")); err != nil {
		panic(err)
	}

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

		if err := writeAll(dst, buf[:n]); err != nil {
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

func writeAll(writer io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := writer.Write(data)
		if err != nil {
			return err
		}
		if n == 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
	}
	return nil
}

func showScannerLimit() {
	longLine := bytes.Repeat([]byte("x"), 70*1024)
	scanner := bufio.NewScanner(bytes.NewReader(longLine))
	if scanner.Scan() {
		fmt.Println("unexpected: default Scanner accepted long token")
	}
	fmt.Println("Default Scanner error:", scanner.Err())

	scanner = bufio.NewScanner(bytes.NewReader(longLine))
	scanner.Buffer(make([]byte, 1024), 128*1024)
	fmt.Println("Expanded Scanner accepted token:", scanner.Scan())
}

func atomicReplace(filename string, data []byte) {
	temp, err := os.CreateTemp(filepath.Dir(filename), ".settings-*")
	if err != nil {
		panic(err)
	}
	tempName := temp.Name()
	defer os.Remove(tempName)

	if err := writeAll(temp, data); err != nil {
		temp.Close()
		panic(err)
	}
	if err := temp.Sync(); err != nil {
		temp.Close()
		panic(err)
	}
	if err := temp.Close(); err != nil {
		panic(err)
	}
	if err := os.Rename(tempName, filename); err != nil {
		panic(err)
	}
	fmt.Println("Atomically replaced:", filepath.Base(filename))
}
