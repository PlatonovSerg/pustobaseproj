package main

import (
	"log"
	"net/http"
	"os"

	"pustobaseproject/internal/db"
	"pustobaseproject/internal/domain/players"
	"pustobaseproject/internal/routes"
	"pustobaseproject/pkg/utils"

	//"pustobaseproject/internal/domain/players"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Загружаем .env
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// 2. Инициализируем шифрование через Fernet
	fernetKey := os.Getenv("FERNET_KEY")
	_, err = utils.NewEncryptionService(fernetKey)
	if err != nil {
		log.Fatalf("Ошибка инициализации шифрования: %v", err)
	}

	// 3. Подключаем базу данных
	db.InitDB()

	// 4. Создаём репозиторий и сервис игрока
	repo := players.NewGormRepository(db.DB)
	service := players.NewService(repo)

	// 5. Создаём роутер и подключаем маршруты
	r := routes.SetupRouter(service)
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
