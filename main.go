package main

import (
    "fmt"
    "log"
    "net/http"
)
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}
func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/Hello", hello)

	server := &http.Server{
        Addr: ":8080",
        Handler :mux,
    }
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(server.ListenAndServe())
}

