package repository

import (
	"chat_project/internal/database"
	"database/sql"
	"errors"
	"fmt"
)

// ErrUserExists - ошибка, если пользователь с таким именем уже есть
var ErrUserExists = errors.New("пользователь с таким именем уже существует")

// CreateUser добавляет нового пользователя в базу, если имя свободно
func CreateUser(username, password string) (int, error) {
	var existingID int

	// Проверяем, есть ли пользователь с таким именем
	err := database.DB.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		username,
	).Scan(&existingID)

	if err != nil && err != sql.ErrNoRows {
		// Возвращаем ошибку, если что-то пошло не так с запросом
		return 0, fmt.Errorf("ошибка проверки пользователя: %w", err)
	}

	if err == nil {
		// Если err == nil, значит пользователь найден
		return 0, ErrUserExists
	}

	var userID int
	// Добавляем нового пользователя в таблицу
	err = database.DB.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		username, password,
	).Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	return userID, nil
}
