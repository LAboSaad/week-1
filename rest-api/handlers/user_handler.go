package handlers

import (
	"demo/httphelper"
	"demo/models"
	"demo/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httphelper.WriteError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateUser(user)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httphelper.WriteJSON(w, res, http.StatusCreated)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	httphelper.WriteJSON(w, h.service.GetAllUsers(), http.StatusOK)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		httphelper.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	httphelper.WriteJSON(w, user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var updated models.User
	json.NewDecoder(r.Body).Decode(&updated)

	user, err := h.service.UpdateUser(id, updated)
	if err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httphelper.WriteJSON(w, user, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.service.DeleteUser(id); err != nil {
		httphelper.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}