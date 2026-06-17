package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"order-service/internal/model"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	userSvcAddr string
)

func InitDB() (*sql.DB, error) {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDB := os.Getenv("MYSQL_DB")
	userSvcAddr = os.Getenv("USER_SERVICE_URL")

	if mysqlUser == "" || mysqlPass == "" || mysqlHost == "" || mysqlDB == "" || userSvcAddr == "" {
		return nil, fmt.Errorf("缺少环境变量，请检查ConfigMap/Secret配置")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDB)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= 30; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("mysql connected")
			return db, nil
		}
		log.Printf("waiting mysql... (%d/30) err=%v\n", i, err)
		time.Sleep(1 * time.Second)
	}
	return nil, err
}

func GetUser(uid string) (*model.User, error) {
	client := http.Client{Timeout: 3 * time.Second}
	url := fmt.Sprintf("%s?id=%s", userSvcAddr, uid)
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("调用user-service失败: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var u model.User
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func AddOrder(user *model.User) (int64, error) {
	res, err := db.Exec("INSERT INTO orders(user_id) VALUES(?)", user.ID)
	if err != nil {
		return 0, fmt.Errorf("插入订单失败: %w", err)
	}
	orderID, _ := res.LastInsertId()
	return orderID, nil
}

func GetOrderByID(orderID int64) (*model.Order, error) {
	order := &model.Order{}
	err := db.QueryRow(`SELECT id,user_id,create_time FROM orders WHERE id=?`, orderID).
		Scan(&order.ID, &order.UserID, &order.CreateTime)
	if err != nil {
		return nil, fmt.Errorf("查询订单失败: %w", err)
	}
	return order, nil
}
