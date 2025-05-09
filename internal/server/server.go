package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет собой структуру, содержащую логгер и HTTP-сервер.
type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// NewServer создает новый экземпляр Server с заданными параметрами.
func NewServer(logger *log.Logger) *Server {
	// Создаем HTTP-роутер
	router := http.NewServeMux()

	// Регистрируем обработчики
	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/upload", handlers.ParseHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}
