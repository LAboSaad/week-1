package handler

import (
	"encoding/json"
	"net/http"
	"week-2/logger"
	"week-2/service"
	"week-2/validation"
)

//HTTP layer (controllers)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req validation.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"event": "invalid_json",
			"error": err.Error(),
		}).Error("failed to parse request")

		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	if err := validation.Validate.Struct(req); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"event": "validation_failed",
			"error": err.Error(),
		}).Warn("invalid request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(map[string]interface{}{
		"event": "create_user_request",
		"email": req.Email,
	}).Info("processing request")

	user := h.service.CreateUser(req.Name, req.Email)
	logger.Log.WithFields(map[string]interface{}{
		"event":   "user_created",
		"user_id": user.ID,
	}).Info("user created")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
