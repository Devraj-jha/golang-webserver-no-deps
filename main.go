package main

import (
	"Go_server_no_dep/src"
	"log"
	"net/http"
)

func main() {
	// Create the user store
	store := src.NewUserStore()

	// Set up the router
	mux := http.NewServeMux()
	
	// Register all routes with their handlers
	mux.HandleFunc("GET /users", store.UserReturn)
	mux.HandleFunc("POST /users", store.CreateUser)
	mux.HandleFunc("GET /users/{id}", store.UserWithID)
	mux.HandleFunc("DELETE /users/{id}", store.DeleteUser)

	// Configure the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	log.Println("Available endpoints:")
	log.Println("  GET    /users           - List all users")
	log.Println("  POST   /users           - Create a new user")
	log.Println("  GET    /users/{id}      - Get a specific user")
	log.Println("  DELETE /users/{id}      - Delete a user")
	
	log.Fatal(server.ListenAndServe())
}