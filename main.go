package main

import (
    "fmt"
    // "log"
    "net/http"
)
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}
func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/Hello", hello)

	
}