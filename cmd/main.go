package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MyApp: ", log.LstdFlags|log.Lshortfile)
	logger.Println("Инициализация сервера...")
	srv := server.InitiateServer(logger)
	logger.Println("Запуск сервера на :8080")
	if err := srv.MyServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка сервера: %v", err)
	}
}
