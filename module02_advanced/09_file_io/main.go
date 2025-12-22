package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// This module covers low-level File I/O operations in Go.
// It corresponds to the 'myio/myfile' section from the original iotestgo.

func main() {
	// Prepare a directory for our file operations
	dir := "test_files"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // Cleanup

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
	// os.Create creates or truncates the named file.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}
	defer f.Close()

	// WriteString
	f.WriteString("Hello, World!\n")
	// Write bytes
	f.Write([]byte("This is a low-level file I/O demo.\n"))

	fmt.Printf("Wrote content to %s\n", filename)
}

func basicRead(filename string) {
	// os.Open opens the named file for reading.
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer f.Close()

	// Read into a buffer
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			// Write to stdout directly
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

	// Copy using a buffer
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

	// Use bufio for buffered I/O
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)

	// io.Copy uses buffers internally, but here we demonstrate manual buffered read/write logic or just io.Copy with buffered streams
	// For demonstration, let's use the standard io.Copy which works great with bufio types too
	n, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	writer.Flush() // Don't forget to flush the writer!
	fmt.Printf("Buffered copied %d bytes to %s\n", n, dstName)
}

func seekAndRead(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Seek to offset 7 (skip "Hello, ")
	_, err = f.Seek(7, io.SeekStart)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 5)
	io.ReadAtLeast(f, buf, 5)
	fmt.Printf("Read after seek(7): %s\n", string(buf)) // Should be "World"
}
