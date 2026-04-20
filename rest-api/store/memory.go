package store

import (
	"demo/models"
	"sync"
)

var (
	Users = make([]models.User, 0)
	// Mu    sync.Mutex
	Mu sync.RWMutex
)
