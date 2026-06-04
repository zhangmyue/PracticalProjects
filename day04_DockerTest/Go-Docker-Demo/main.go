package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "✅ Go 程序在 Docker 中运行成功！")
	})

	// 容器内部监听 8080 端口
	fmt.Println("服务启动在 :8080")
	http.ListenAndServe(":8080", nil)
}
