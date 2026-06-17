package service

import (
	"encoding/json"
	"order-service/internal/db"
)

func CreateOrder(uid string) (string, error) {
	user, err := db.GetUser(uid)
	if err != nil {
		return "", err
	}
	orderID, err := db.AddOrder(user)
	if err != nil {
		return "", err
	}
	respMap := map[string]interface{}{
		"code":     0,
		"msg":      "订单创建成功",
		"order_id": orderID,
		"user":     user,
	}
	data, _ := json.Marshal(respMap)
	return string(data), nil
}

func GetOrder(orderID int64) (string, error) {
	order, err := db.GetOrderByID(orderID)
	if err != nil {
		return "", err
	}
	data, _ := json.Marshal(order)
	return string(data), nil
}
