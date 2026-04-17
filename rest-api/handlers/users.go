package handlers

import (
	"demo/models"
	"demo/store"
	"encoding/json"
	"net/http"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	user.ID = len(store.Users) + 1
	store.Users = append(store.Users, user)
	store.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(store.Users)
}