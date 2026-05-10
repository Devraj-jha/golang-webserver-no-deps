package src

import (
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



