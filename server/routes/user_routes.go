package routes

import (
	"github.com/gorilla/mux"
	"project1/controller"
)

func InitUserRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/adduser", controller.AddUser).Methods("POST")
	return router
}
