package src

import (
	"sync"
)

// User represents a user in our system
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserStore holds the in-memory database with thread-safe access
type UserStore struct {
	Mu    sync.RWMutex
	Users map[string]User
}

// NewUserStore creates a new UserStore with initialized map
func NewUserStore() *UserStore {
	return &UserStore{
		Users: make(map[string]User),
	}
}