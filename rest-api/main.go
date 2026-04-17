package main

import (
	"demo/handlers"
	"demo/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/health", handlers.HealthHandler)

	// /users (GET, POST)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListUsers(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /users/{id}
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUser(w, r)
		case http.MethodPut:
			handlers.UpdateUser(w, r)
		case http.MethodDelete:
			handlers.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8000")

	// Apply middleware
	log.Fatal(http.ListenAndServe(":8000", middleware.Logger(mux)))
}