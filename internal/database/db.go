package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(dbURL string) error {
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("БД недоступна: %w", err)
	}

	fmt.Println("!> Подключено к PostgreSQL!")
	return nil
}
