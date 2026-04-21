package repository
//data access
import (
	"sync"
	"week-2/models"
)

type UserRepo interface {
	Create(user models.User) models.User
}

type InMemoryUserRepo struct {
	mu    sync.Mutex
	users []models.User
}

func NewUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{}
}

func (r *InMemoryUserRepo) Create(user models.User) models.User {

	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user

}
