package main

import (
	"demo/handlers"
	"demo/repository"
	"demo/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo := &repository.InMemoryUserRepo{}
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users", handler.ListUsers).Methods("GET")

	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}