package services

import "demo/models"

type UserRepository interface {
	CreateUser(user models.User) models.User
	GetAllUsers() []models.User
	GetUserByID(id int) (models.User, bool)
	UpdateUser(id int, user models.User) (models.User, bool)
	DeleteUser(id int) bool
}