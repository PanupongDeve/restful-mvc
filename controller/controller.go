package controller

import (
    "log"
    "net/http"

	"github.com/gorilla/mux"
	"restful-mvc/config"
	user"restful-mvc/controller/user"
	cat"restful-mvc/controller/cat"
)

var PORT = ":"+config.PORT

func Run() {
	myRouter := mux.NewRouter().StrictSlash(true)
	user.AddRouter(myRouter);
	cat.AddRouter(myRouter);
    log.Fatal(http.ListenAndServe(PORT, myRouter))
}