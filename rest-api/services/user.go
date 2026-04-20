package services
//business logic layer
import (
	"demo/models"
	"demo/repository"
	"demo/validation"
	"errors"
)

func CreateUser(user models.User) (models.User, error) {
	if err := validation.ValidateUser(user); err != nil {
		return models.User{}, err
	}

	// Example business rule (optional)
	// check duplicate names
	users := repository.GetAllUsers()
	for _, u := range users {
		if u.Name == user.Name {
			return models.User{}, errors.New("user already exists")
		}
	}

	return repository.CreateUser(user), nil
}

func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}

func GetUserByID(id int) (models.User, error) {
	user, found := repository.GetUserByID(id)
	if !found {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func UpdateUser(id int, updated models.User) (models.User, error) {
	if err := validation.ValidateUser(updated); err != nil {
		return models.User{}, err
	}

	user, found := repository.UpdateUser(id, updated)
	if !found {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func DeleteUser(id int) error {
	if ok := repository.DeleteUser(id); !ok {
		return errors.New("user not found")
	}
	return nil
}