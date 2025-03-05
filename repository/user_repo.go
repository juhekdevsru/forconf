package repository

import (
	"chat_project/internal/database"
	"database/sql"
	"fmt"
)

func CreateUser(username string) (int, error) {
	var userID int
	query := `INSERT INTO users (username) VALUES ($1) RETURNING id`
	err := database.DB.QueryRow(query, username).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("не удалсоь создать пользователя")
		}
		return 0, err
	}

	return userID, nil
}
