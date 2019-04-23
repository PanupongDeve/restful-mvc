package main

import (
	"fmt"
    "net/http"

	controller"restful-mvc/controller"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "All Users Endpoint Hit")
}

func newUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New User Endpoint Hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update User Endpoint Hit")
}

func handleRequests() {
    controller.Run()
}


func main() {
    fmt.Println("Go ORM Tutorial")

    // Handle Subsequent requests
    handleRequests()
}