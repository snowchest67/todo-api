package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/snowchest67/todo-api/internal/handler"
	"github.com/snowchest67/todo-api/internal/repository"
)

 
func main() {

	ctx := context.Background()

	connString := "postgres://todo_user:secret@localhost:5432/todo_db?sslmode=disable"

	repo, err := repository.NewPostgresRepo(ctx, connString)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer repo.Close()

	todoHandler := handler.NewTodoHandler(repo)

	log.Println("Connected to PostgreSQL")

	http.HandleFunc("/", handler.RootHandler)

	http.HandleFunc("/health", handler.HealthHandler)

	http.Handle("/todos", todoHandler)

	http.Handle("/todos/", todoHandler)

	srv := &http.Server{Addr: ":8080"}
	go func() { // так как srv.ListenAndServe() блокирует поток, чтобы реализовать Shutting down выносим отдельно
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctxShut, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctxShut)
}