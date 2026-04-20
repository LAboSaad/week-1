package handlers

import (
	"demo/httphelper"
	"demo/models"
	"demo/store"
	"demo/validation"
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
		httphelper.WriteError(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateUser(user); err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	user.ID = len(store.Users) + 1
	store.Users = append(store.Users, user)
	store.Mu.Unlock()

	httphelper.WriteJSON(w, user, http.StatusCreated)
}

// GET /users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	httphelper.WriteJSON(w, store.Users, http.StatusOK)
}

// GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	store.Mu.RLock()
	defer store.Mu.RUnlock()

	for _, user := range store.Users {
		if user.ID == id {
			httphelper.WriteJSON(w, user, http.StatusOK)
			return
		}
	}

	httphelper.WriteError(w, "user not found", http.StatusNotFound)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		httphelper.WriteError(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateUser(updated); err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	defer store.Mu.Unlock()

	for i, user := range store.Users {
		if user.ID == id {
			user.Name = updated.Name
			store.Users[i] = user
			httphelper.WriteJSON(w, user, http.StatusOK)
			return
		}
	}

	httphelper.WriteError(w, "user not found", http.StatusNotFound)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
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

	httphelper.WriteError(w, "user not found", http.StatusNotFound)
}
