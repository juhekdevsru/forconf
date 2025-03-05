package main

import (
	"chat_project/internal/config"
	"chat_project/internal/database"
	"fmt"
	"log"
)

func main() {
	fmt.Println("!>>> Запуск сервера...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("x Ошибка загрузки конфига:", err)
	}

	if err := database.ConnectDB(cfg.DatabaseURL); err != nil {
		log.Fatal("x Ошибка подключения к БД:", err)
	}

	fmt.Println("[ok] Подключение к PostgreSQL!")
	fmt.Println("[go] Сервер запущен")
}
