package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Создаем HTTP сервер
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleRequest),
	}

	// Канал для ошибок сервера
	serverErrors := make(chan error, 1)

	// Запускаем сервер в горутине
	go func() {
		log.Println("Server listening on", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Канал для сигналов ОС
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Блокируем main, пока не получим сигнал или ошибку сервера
	select {
	case err := <-serverErrors:
		log.Fatalf("Server error: %v", err)

	case sig := <-shutdown:
		log.Printf("Received %v signal, starting graceful shutdown...", sig)

		// Даем 30 секунд на завершение работы
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Пытаемся gracefully shutdown сервер
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
			if err := server.Close(); err != nil {
				log.Fatalf("Forced shutdown failed: %v", err)
			}
		}
	}

	log.Println("Server gracefully stopped")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Имитация долгой обработки запроса
	if r.URL.Path == "/long" {
		time.Sleep(10 * time.Second)
	}
	w.Write([]byte("Hello, World!"))
}
