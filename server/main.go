package main

import (
	"log"
	"net/http"
	"project1/db"
	"project1/middleware"
	"project1/routes"
)

func main() {
	router := routes.InitUserRoutes()
	handler := middleware.CORS(router)
	
	// Initialize Redis
	db.InitRedis()
	
	if err := db.ConnectToMongo("mongodb+srv://mango26june:mango123@cluster0.ga9pq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}