package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"week-2/handler"
	"week-2/logger"
	"week-2/repository"
	"week-2/service"
)

func main() {

	logger.Init()

	repo := repository.NewUserRepo()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	logger.Log.Info("server started on :8000")

	http.ListenAndServe(":8000", r)
}
