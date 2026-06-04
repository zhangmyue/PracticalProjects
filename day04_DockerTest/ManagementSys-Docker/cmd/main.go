package main

import (
	"ManagementSys-Docker/internal/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	svc := service.NewStudentService()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("欢迎使用学生管理系统 (CLI版)")
	for {
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "add":
			if len(args) != 3 {
				fmt.Println("用法: add 名字 年龄")
				continue
			}
			age, _ := strconv.Atoi(args[2])
			student, err := svc.AddStudent(args[1], age)
			if err != nil {
				fmt.Println("添加失败", err)
				continue
			}
			fmt.Println("添加成功:", student)

		case "list":
			stus, err := svc.ListStudents()
			if err != nil {
				fmt.Println("查看失败", err)
			}
			for _, s := range stus {
				fmt.Println(s)
			}
		case "get":
			if len(args) != 2 {
				fmt.Println("用法: get ID")
				continue
			}
			id, _ := strconv.Atoi(args[1])
			s, err := svc.GetStudentByID(id)
			if err != nil {
				fmt.Printf("学生不存在, err=%v", err)
			} else {
				fmt.Println(s)
			}
		case "delete":
			if len(args) != 2 {
				fmt.Println("用法: delete ID")
				continue
			}
			id, _ := strconv.Atoi(args[1])
			_, err := svc.DeleteStudentByID(id)
			if err != nil {
				fmt.Printf("学生不存在, err=%v", err)
			} else {
				fmt.Println("删除成功")
			}
		case "update":
			if len(args) != 4 {
				fmt.Println("用法: update ID 名字 年龄")
				continue
			}
			id, _ := strconv.Atoi(args[1])
			age, _ := strconv.Atoi(args[3])
			_, err := svc.UpdateStudentByID(id, args[2], age)
			if err != nil {
				fmt.Printf("学生不存在, err=%v", err)
			} else {
				fmt.Println("更新成功")
			}
		case "exit":
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("未知命令:", args[0])
		}
	}
}
