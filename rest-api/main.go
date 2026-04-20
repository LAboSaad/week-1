package main

import (
	"demo/handlers"
	"demo/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// mux := http.NewServeMux()
	r := mux.NewRouter()
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

	// Health
	// mux.HandleFunc("/health", handlers.HealthHandler)
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	// /users (GET, POST)
	// mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handlers.ListUsers(w, r)
	// 	case http.MethodPost:
	// 		handlers.CreateUser(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.ListUsers).Methods("GET")

	// /users/{id}
	// mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handlers.GetUser(w, r)
	// 	case http.MethodPut:
	// 		handlers.UpdateUser(w, r)
	// 	case http.MethodDelete:
	// 		handlers.DeleteUser(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	log.Println("Server running on :8000")

	// Apply middleware
	// log.Fatal(http.ListenAndServe(":8000", middleware.Logger(r)))
	log.Fatal(http.ListenAndServe(":8000", r))
}
