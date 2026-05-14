package src

import (
	"fmt"
	"net/http"
	"time"
)

// UserReturn handles GET /users - returns all users
func (s *UserStore) UserReturn(w http.ResponseWriter, r *http.Request) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	// Convert map to slice for JSON array
	userList := make([]User, 0, len(s.Users))
	for _, user := range s.Users {
		userList = append(userList, user)
	}

	WriteJSON(w, http.StatusOK, userList)
}

// CreateUser handles POST /users - creates a new user
func (s *UserStore) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	
	// Read JSON from request body
	if err := ReadJSON(r, &newUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Generate unique ID using timestamp
	newUser.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	// Store the user (write lock needed)
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Users[newUser.ID] = newUser

	// Return 201 Created with the user data
	WriteJSON(w, http.StatusCreated, newUser)
}

// UserWithID handles GET /users/{id} - returns a single user
func (s *UserStore) UserWithID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.Mu.RLock()
	defer s.Mu.RUnlock()

	user, exists := s.Users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

// DeleteUser handles DELETE /users/{id} - deletes a user
func (s *UserStore) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.Mu.Lock()
	defer s.Mu.Unlock()

	_, exists := s.Users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(s.Users, id)
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}