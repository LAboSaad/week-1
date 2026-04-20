package services

import (
	"demo/models"
	"demo/validation"
	"errors"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	if err := validation.ValidateUser(user); err != nil {
		return models.User{}, err
	}

	for _, u := range s.repo.GetAllUsers() {
		if u.Name == user.Name {
			return models.User{}, errors.New("user already exists")
		}
	}

	return s.repo.CreateUser(user), nil
}

func (s *UserService) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	user, found := s.repo.GetUserByID(id)
	if !found {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, updated models.User) (models.User, error) {
	if err := validation.ValidateUser(updated); err != nil {
		return models.User{}, err
	}

	user, found := s.repo.UpdateUser(id, updated)
	if !found {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	if ok := s.repo.DeleteUser(id); !ok {
		return errors.New("user not found")
	}
	return nil
}