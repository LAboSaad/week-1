package handlers

import (
	"demo/models"
	"demo/store"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Lock() should be used for writes while RLock() should be used for reads
// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	user.ID = len(store.Users) + 1
	store.Users = append(store.Users, user)
	store.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GET /users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.Users)
}

// GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	// idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	// id, _ := strconv.Atoi(idStr)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "something went wrong maybee?", http.StatusBadRequest)
		return
	}

	store.Mu.RLock()
	defer store.Mu.RUnlock()

	for _, user := range store.Users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	// id, _ := strconv.Atoi(idStr)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	defer store.Mu.Unlock()

	for i, user := range store.Users {
		if user.ID == id {
			user.Name = updated.Name
			store.Users[i] = user
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	// id, _ := strconv.Atoi(idStr)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	defer store.Mu.Unlock()

	for i, user := range store.Users {
		if user.ID == id {
			store.Users = append(store.Users[:i], store.Users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
