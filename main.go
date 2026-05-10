package main

import (
    
    "log"
    "net/http"
    "Go_server_no_dep/src"
)


func main() {
    mux := http.NewServeMux()
	// mux.HandleFunc("/Hello", src.Hello)
    mux.HandleFunc("POST /users", src.CreateUser)
    mux.HandleFunc("GET /users", src.UserReturn)
    mux.HandleFunc("GET /users/{id}", src.UserWithID)


	server := &http.Server{
        Addr: ":8080",
        Handler :mux,
    }
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(server.ListenAndServe())
}
// broswer sends a get request -> when it needs something

