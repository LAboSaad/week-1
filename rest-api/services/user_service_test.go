package services

//To run the tests we write go test ./services -v

import (
	"demo/models"
	"testing"
)

type MockRepo struct {
	users []models.User
}

func (m *MockRepo) CreateUser(user models.User) models.User {
	user.ID = len(m.users) + 1
	m.users = append(m.users, user)
	return user
}

func (m *MockRepo) GetAllUsers() []models.User {
	return m.users
}

func (m *MockRepo) GetUserByID(id int) (models.User, bool) {
	for _, u := range m.users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func (m *MockRepo) UpdateUser(id int, updated models.User) (models.User, bool) {
	for i, u := range m.users {
		if u.ID == id {
			u.Name = updated.Name
			m.users[i] = u
			return u, true
		}
	}
	return models.User{}, false
}

func (m *MockRepo) DeleteUser(id int) bool {
	for i, u := range m.users {
		if u.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return true
		}
	}
	return false
}

func TestCreateUser(t *testing.T) {
	repo := &MockRepo{}
	service := NewUserService(repo)

	user := models.User{Name: "Layla"}

	result, err := service.CreateUser(user)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "Layla" {
		t.Errorf("expected Layla, got %s", result.Name)
	}
}

func TestGetAllUsers(t *testing.T) {
	repo := &MockRepo{
		users: []models.User{
			{ID: 1, Name: "Layla"},
			{ID: 2, Name: "John"},
		},
	}

	service := NewUserService(repo)

	users := service.GetAllUsers()

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}
func TestGetUserByID_Success(t *testing.T) {
	repo := &MockRepo{
		users: []models.User{
			{ID: 1, Name: "Layla"},
		},
	}

	service := NewUserService(repo)

	user, err := service.GetUserByID(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Name != "Layla" {
		t.Errorf("expected Layla, got %s", user.Name)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	repo := &MockRepo{}
	service := NewUserService(repo)

	_, err := service.GetUserByID(1)

	if err == nil {
		t.Error("expected error, got nil")
	}
}
func TestUpdateUser_Success(t *testing.T) {
	repo := &MockRepo{
		users: []models.User{
			{ID: 1, Name: "Old"},
		},
	}

	service := NewUserService(repo)

	updated := models.User{Name: "New"}

	user, err := service.UpdateUser(1, updated)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Name != "New" {
		t.Errorf("expected New, got %s", user.Name)
	}
}

func TestUpdateUser_NotFound(t *testing.T) {
	repo := &MockRepo{}
	service := NewUserService(repo)

	_, err := service.UpdateUser(1, models.User{Name: "Test"})

	if err == nil {
		t.Error("expected error, got nil")
	}
}
func TestDeleteUser_Success(t *testing.T) {
	repo := &MockRepo{
		users: []models.User{
			{ID: 1, Name: "Layla"},
		},
	}

	service := NewUserService(repo)

	err := service.DeleteUser(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.users) != 0 {
		t.Error("user was not deleted")
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	repo := &MockRepo{}
	service := NewUserService(repo)

	err := service.DeleteUser(1)

	if err == nil {
		t.Error("expected error, got nil")
	}
}
