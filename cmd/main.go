package main

import (
	//"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	logger.Println("Запуск сервера на порту 8080")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
