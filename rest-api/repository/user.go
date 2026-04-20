package repository

//data layer
import (
	"demo/models"
	"demo/store"
)

func CreateUser(user models.User) models.User {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	user.ID = len(store.Users) + 1
	store.Users = append(store.Users, user)

	return user
}

func GetAllUsers() []models.User {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	return store.Users
}

func GetUserByID(id int) (models.User, bool) {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	for _, u := range store.Users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func UpdateUser(id int, updated models.User) (models.User, bool) {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	for i, u := range store.Users {
		if u.ID == id {
			u.Name = updated.Name
			store.Users[i] = u
			return u, true
		}
	}
	return models.User{}, false
}

func DeleteUser(id int) bool {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	for i, u := range store.Users {
		if u.ID == id {
			store.Users = append(store.Users[:i], store.Users[i+1:]...)
			return true
		}
	}
	return false
}
