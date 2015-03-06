package main

import (
    "net/http"
    "github.com/jasonmichels/out-of-office/controllers"
    "github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	AccountController := new(controllers.AccountController)
	DashboardController := new(controllers.DashboardController)
	StatusController := new(controllers.StatusController)

    router.HandleFunc("/login", AccountController.LoginHandler).Methods("GET")
    router.HandleFunc("/login", AccountController.AuthenticateHandler).Methods("POST")
    router.HandleFunc("/dashboard", DashboardController.DashboardHandler).Methods("GET")
    router.HandleFunc("/status/create", StatusController.CreateHandler).Methods("GET")

	http.Handle("/", router)
    http.ListenAndServe(":9005", nil)
}