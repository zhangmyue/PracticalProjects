package model

import "time"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID         int64     `json:"id"`
	UserID     int       `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
}
