package main

import (
	"Go_server_no_dep/src"
	"log"
	"net/http"
	"sync"
)
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type UserStore struct {
    mu sync.RWMutex // prevents data corruption. when multiple request happen.
    users map[string]User // store users 

}
func NewUserStore() *UserStore{
    return &UserStore{
        users: make(map[string]User),
    }
}
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

