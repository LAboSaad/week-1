package service
//business logic
import (
	
	"week-2/models"
	"week-2/repository"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(r repository.UserRepo) *UserService {
	return &UserService{repo: r}
}

//Define a method called CreateUser that belongs to UserService, takes name and email, and returns a User
func (s *UserService) CreateUser(name, email string) models.User{
	user:=models.User{
		Name: name,
		Email: email,
	}
	return s.repo.Create(user)
}
