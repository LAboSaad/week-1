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

//Mocking cases
// func TestCreateUser(t *testing.T) {
// 	repo := &MockRepo{}
// 	service := NewUserService(repo)

// 	user := models.User{Name: "Layla"}

// 	result, err := service.CreateUser(user)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if result.Name != "Layla" {
// 		t.Errorf("expected Layla, got %s", result.Name)
// 	}
// }

// func TestGetAllUsers(t *testing.T) {
// 	repo := &MockRepo{
// 		users: []models.User{
// 			{ID: 1, Name: "Layla"},
// 			{ID: 2, Name: "John"},
// 		},
// 	}

// 	service := NewUserService(repo)

// 	users := service.GetAllUsers()

// 	if len(users) != 2 {
// 		t.Errorf("expected 2 users, got %d", len(users))
// 	}
// }
// func TestGetUserByID_Success(t *testing.T) {
// 	repo := &MockRepo{
// 		users: []models.User{
// 			{ID: 1, Name: "Layla"},
// 		},
// 	}

// 	service := NewUserService(repo)

// 	user, err := service.GetUserByID(1)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if user.Name != "Layla" {
// 		t.Errorf("expected Layla, got %s", user.Name)
// 	}
// }

// func TestGetUserByID_NotFound(t *testing.T) {
// 	repo := &MockRepo{}
// 	service := NewUserService(repo)

// 	_, err := service.GetUserByID(1)

// 	if err == nil {
// 		t.Error("expected error, got nil")
// 	}
// }
// func TestUpdateUser_Success(t *testing.T) {
// 	repo := &MockRepo{
// 		users: []models.User{
// 			{ID: 1, Name: "Old"},
// 		},
// 	}

// 	service := NewUserService(repo)

// 	updated := models.User{Name: "New"}

// 	user, err := service.UpdateUser(1, updated)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if user.Name != "New" {
// 		t.Errorf("expected New, got %s", user.Name)
// 	}
// }

// func TestUpdateUser_NotFound(t *testing.T) {
// 	repo := &MockRepo{}
// 	service := NewUserService(repo)

// 	_, err := service.UpdateUser(1, models.User{Name: "Test"})

// 	if err == nil {
// 		t.Error("expected error, got nil")
// 	}
// }
// func TestDeleteUser_Success(t *testing.T) {
// 	repo := &MockRepo{
// 		users: []models.User{
// 			{ID: 1, Name: "Layla"},
// 		},
// 	}

// 	service := NewUserService(repo)

// 	err := service.DeleteUser(1)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if len(repo.users) != 0 {
// 		t.Error("user was not deleted")
// 	}
// }

// func TestDeleteUser_NotFound(t *testing.T) {
// 	repo := &MockRepo{}
// 	service := NewUserService(repo)

// 	err := service.DeleteUser(1)

// 	if err == nil {
// 		t.Error("expected error, got nil")
// 	}
// }

//Table-driven tests

// =======================
// CREATE USER TEST
// =======================
func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		input    models.User
		existing []models.User
		wantErr  bool
	}{
		{
			name:     "valid user",
			input:    models.User{Name: "Layla"},
			existing: []models.User{},
			wantErr:  false,
		},
		{
			name:  "duplicate user",
			input: models.User{Name: "Layla"},
			existing: []models.User{
				{ID: 1, Name: "Layla"},
			},
			wantErr: true,
		},
		{
			name:     "invalid user (empty name)",
			input:    models.User{Name: ""},
			existing: []models.User{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepo{users: tt.existing}
			service := NewUserService(repo)

			_, err := service.CreateUser(tt.input)

			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// =======================
// GET ALL USERS
// =======================
func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name     string
		existing []models.User
		expected int
	}{
		{
			name: "multiple users",
			existing: []models.User{
				{ID: 1, Name: "Layla"},
				{ID: 2, Name: "John"},
			},
			expected: 2,
		},
		{
			name:     "no users",
			existing: []models.User{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepo{users: tt.existing}
			service := NewUserService(repo)

			users := service.GetAllUsers()

			if len(users) != tt.expected {
				t.Errorf("expected %d users, got %d", tt.expected, len(users))
			}
		})
	}
}

// =======================
// GET USER BY ID
// =======================
func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name     string
		existing []models.User
		id       int
		wantErr  bool
	}{
		{
			name: "user exists",
			existing: []models.User{
				{ID: 1, Name: "Layla"},
			},
			id:      1,
			wantErr: false,
		},
		{
			name:     "user not found",
			existing: []models.User{},
			id:       1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepo{users: tt.existing}
			service := NewUserService(repo)

			_, err := service.GetUserByID(tt.id)

			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// =======================
// UPDATE USER
// =======================
func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name     string
		existing []models.User
		id       int
		input    models.User
		wantErr  bool
	}{
		{
			name: "success update",
			existing: []models.User{
				{ID: 1, Name: "Old"},
			},
			id:      1,
			input:   models.User{Name: "New"},
			wantErr: false,
		},
		{
			name:     "user not found",
			existing: []models.User{},
			id:       1,
			input:    models.User{Name: "Test"},
			wantErr:  true,
		},
		{
			name:     "invalid data",
			existing: []models.User{{ID: 1, Name: "Old"}},
			id:       1,
			input:    models.User{Name: ""},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepo{users: tt.existing}
			service := NewUserService(repo)

			_, err := service.UpdateUser(tt.id, tt.input)

			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// =======================
// DELETE USER
// =======================
func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name     string
		existing []models.User
		id       int
		wantErr  bool
	}{
		{
			name: "delete success",
			existing: []models.User{
				{ID: 1, Name: "Layla"},
			},
			id:      1,
			wantErr: false,
		},
		{
			name:     "user not found",
			existing: []models.User{},
			id:       1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepo{users: tt.existing}
			service := NewUserService(repo)

			err := service.DeleteUser(tt.id)

			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
