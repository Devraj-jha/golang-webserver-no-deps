package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func UserReturn(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Will return a list of users")


}
func CreateUser(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Will create users")

}
func UserWithID(w http.ResponseWriter, r *http.Request){
     id := r.PathValue("id")

     fmt.Print("will return user with id: %s", id)


}

// takes a go value, converts it into json, sends it back to the clinet. 
// with an http status code.
func writeJSON(w http.ResponseWriter, status int, v any) error {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(status)
     return json.NewEncoder(w).Encode(v)
}
// It takes JSON from the request body and converts it into a Go struct or variable.
func readJSON(r *http.Request, dest any) error {

    return json.NewDecoder(r.Body).Decode(dest)
    
}


