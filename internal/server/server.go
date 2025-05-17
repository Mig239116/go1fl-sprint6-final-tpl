package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	MyLogger *log.Logger
	MyServer *http.Server
}

func InitiateServer(myLogger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.IndexHanlder)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     myLogger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		MyLogger: myLogger,
		MyServer: httpServer,
	}
}
