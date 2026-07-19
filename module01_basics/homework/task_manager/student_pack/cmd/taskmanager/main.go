package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"taskmanager/taskmanager"
)

func main() {
	manager := taskmanager.NewManager()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Task Manager：输入 add <标题>、list、complete <ID>、delete <ID> 或 exit")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		args := strings.Fields(line)

		switch args[0] {
		case "add":
			if len(args) < 2 {
				fmt.Println("用法：add <标题>")
				continue
			}
			task, err := manager.Add(strings.Join(args[1:], " "))
			if err != nil {
				fmt.Printf("错误：%v\n", err)
				continue
			}
			fmt.Printf("已添加：%d [ ] %s\n", task.ID, task.Title)
		case "list":
			if len(args) != 1 {
				fmt.Println("用法：list")
				continue
			}
			tasks := manager.List()
			if len(tasks) == 0 {
				fmt.Println("暂无任务")
				continue
			}
			for _, task := range tasks {
				status := "[ ]"
				if task.Completed {
					status = "[x]"
				}
				fmt.Printf("%d %s %s\n", task.ID, status, task.Title)
			}
		case "complete":
			if len(args) != 2 {
				fmt.Println("用法：complete <ID>")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("错误：ID 必须是整数")
				continue
			}
			task, err := manager.Complete(id)
			if err != nil {
				fmt.Printf("错误：%v\n", err)
				continue
			}
			fmt.Printf("已完成：%d [x] %s\n", task.ID, task.Title)
		case "delete":
			if len(args) != 2 {
				fmt.Println("用法：delete <ID>")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("错误：ID 必须是整数")
				continue
			}
			if err := manager.Delete(id); err != nil {
				fmt.Printf("错误：%v\n", err)
				continue
			}
			fmt.Printf("已删除任务 %d\n", id)
		case "exit":
			if len(args) != 1 {
				fmt.Println("用法：exit")
				continue
			}
			fmt.Println("再见！")
			return
		default:
			fmt.Println("未知命令；请输入 add、list、complete、delete 或 exit")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入失败：%v\n", err)
	}
}
