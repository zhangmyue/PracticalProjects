package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 1. 连接服务端
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		log.Fatal("连接服务端失败：", err)
	}
	defer conn.Close() // 用完关闭

	// 2. 从控制台读取输入
	fmt.Print("请输入要发送的消息：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	// 3. 发送消息（必须加 \n，因为服务端按行读取）
	_, err = fmt.Fprintf(conn, msg+"\n")
	if err != nil {
		log.Fatal("发送失败：", err)
	}

	// 4. 读取服务端返回的大写消息
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("接收失败：", err)
	}

	// 5. 打印结果
	fmt.Println("服务端返回：", response)
}
