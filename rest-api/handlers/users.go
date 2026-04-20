package handlers

//HTTP layer ONLY

import (
	"demo/httphelper"
	"demo/models"
	"demo/services"
	
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

	result, err := services.CreateUser(user)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httphelper.WriteJSON(w, result, http.StatusCreated)
}

// GET /users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetAllUsers()
	httphelper.WriteJSON(w, users, http.StatusOK)
}

// GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByID(id)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	httphelper.WriteJSON(w, user, http.StatusOK)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		httphelper.WriteError(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	user, err := services.UpdateUser(id, updated)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httphelper.WriteJSON(w, user, http.StatusOK)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		httphelper.WriteError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := services.DeleteUser(id); err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
