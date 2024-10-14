package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tes-sudo/online-learning-platform/user-service/handlers"
	"github.com/Tes-sudo/online-learning-platform/user-service/repository"
)

func main() {
	fmt.Println("Starting User Service...")

	// Initialize database
	repository.InitDB()

	// Set up HTTP server
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateUserHandler(w, r)
		case http.MethodGet:
			handlers.GetUserHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteUserHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// TODO: Add more routes as needed

	// Create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	log.Printf("Server started on %s", server.Addr)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Close database connection
	sqlDB, err := repository.DB.DB()
	if err != nil {
		log.Printf("Error getting underlying SQL DB: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}

	log.Println("Server gracefully stopped")
}
