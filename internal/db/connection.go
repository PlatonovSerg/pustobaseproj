package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// Подключение к SQLite (можно легко сменить на Postgres)
	DB, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	fmt.Println("Успешное подключение к SQLite")

	// Авто-миграция моделей
	if err := RunMigrations(); err != nil {
		log.Fatalf("Ошибка миграции БД: %v", err)
	}
}
