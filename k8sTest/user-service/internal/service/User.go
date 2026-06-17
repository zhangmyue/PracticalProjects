package service

import (
	"database/sql"
	"user-service/internal/model"
)

func GetUserByID(db *sql.DB, id string) (*model.User, error) {

	var user model.User

	err := db.QueryRow(
		"SELECT id,name FROM users WHERE id=?",
		id,
	).Scan(
		&user.ID,
		&user.Name,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
