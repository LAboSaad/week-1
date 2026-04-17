package main

import (
	"demo/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/users", handlers.CreateUser)
	http.HandleFunc("/users/list", handlers.ListUsers)

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}