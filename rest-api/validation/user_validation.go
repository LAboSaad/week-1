package validation

import (
	"demo/models"
	"errors"
)

func ValidateUser(u models.User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if len(u.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}
	return nil
}