package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	dbb "user-service/internal/db"
	"user-service/internal/service"
)

var db *sql.DB

func UserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	user, err := service.GetUserByID(db, id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func main() {

	var err error

	db, err = dbb.InitDB()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/user", UserHandler)

	http.ListenAndServe(":8080", nil)
}
