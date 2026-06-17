package main

import (
	"log"
	"net/http"
	dbb "order-service/internal/db"
	"order-service/internal/service"
	"strconv"
)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	uidStr := r.URL.Query().Get("uid")
	if uidStr == "" {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"code":400,"msg":"uid不能为空"}`))
		return
	}
	resp, err := service.CreateOrder(uidStr)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"code":500,"msg":"` + err.Error() + `"}`))
		return
	}
	_, _ = w.Write([]byte(resp))
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"code":400,"msg":"订单id不能为空"}`))
		return
	}
	orderIDInt64, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"code":400,"msg":"` + err.Error() + `"}`))
		return
	}
	resp, err := service.GetOrder(orderIDInt64)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"code":500,"msg":"` + err.Error() + `"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(resp))
}

func main() {
	_, err := dbb.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	http.HandleFunc("/order/create", createOrderHandler)
	http.HandleFunc("/order/get", getOrderHandler)
	log.Println("order service start :8080")
	_ = http.ListenAndServe(":8080", nil)
}
