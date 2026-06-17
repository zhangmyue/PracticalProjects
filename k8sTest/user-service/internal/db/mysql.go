package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := "root:root@tcp(mysql:3306)/test?parseTime=true"

	db, err := sql.Open("mysql", dsn)
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
