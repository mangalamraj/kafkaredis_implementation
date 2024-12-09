package main

import (
	"log"
	"net/http"
	"project1/middleware"
	"project1/routes"
)

func main() {
	router := routes.InitUserRoutes()
	handler := middleware.CORS(router)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}