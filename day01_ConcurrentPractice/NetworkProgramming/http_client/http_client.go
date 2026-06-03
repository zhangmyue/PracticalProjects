package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("请输入要访问的网站代码：1.http://localhost:8090/hello 2.http://localhost:8090/headers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()
	if msg == "1" {
		sayHello()
	} else if msg == "2" {
		getHeaders()
	}

}

func sayHello() {
	resp, err := http.Get("http://localhost:8090/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 打印服务端返回的 hello
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body)) // 输出：hello
}

func getHeaders() {
	resp, err := http.Get("http://localhost:8090/headers")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取服务端返回的所有请求头
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
