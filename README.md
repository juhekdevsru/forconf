СТРУКТУРА ПРОЕКТА: 

cmd/server/ — точка входа в приложение (main.go)
config/ — конфигурационные файлы (config.json)
internal/
config/ — загрузка и парсинг конфига (config.go)
database/ — подключение к базе данных (db.go)
models/ — модели данных (user.go, message.go)
repository/ — работа с БД (user_repo.go, message_repo.go)
server/ — логика сервера (handlers.go, routes.go)
pkg/ — вспомогательные функции и утилиты (utils.go)
.gitignore — игнорируемые файлы
go.mod — модуль Go
go.sum — контрольные суммы зависимостей
README.md — документация проекта
users.json — временное хранилище пользователей

