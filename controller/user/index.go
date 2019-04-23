
package controller

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

var prefix = "user"
var routeAll = "/"+prefix+"s"
var routeById = "/"+prefix+"/{id}"

func all(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "All Users Endpoint Hit")
}

func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	id := vars["id"]
	response := "All Users Endpoint By Id:" + id
    fmt.Fprintf(w, response)
}

func new(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New User Endpoint Hit")
}

func delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func update(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update User Endpoint Hit")
}

func AddRouter(myRouter *mux.Router) {
	myRouter.HandleFunc(routeAll, all).Methods("GET")
	myRouter.HandleFunc(routeById, getById).Methods("GET")
    myRouter.HandleFunc(routeById, delete).Methods("DELETE")
    myRouter.HandleFunc(routeById, update).Methods("PUT")
    myRouter.HandleFunc(routeById, new).Methods("POST")
}