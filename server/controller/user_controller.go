package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project1/db"
	"project1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userCollection := db.GetCollection("user_app", "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"id":      result.InsertedID,
	})
	
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	// Try to get from cache first
	cachedUsers, err := db.GetCache("users")
	fmt.Println("hiii",cachedUsers)
	if err == nil {
		// Cache hit
		w.Write([]byte(cachedUsers))
		return
	}

	// Cache miss - get from MongoDB
	userCollection := db.GetCollection("user_app", "users")
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, "Error decoding users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Println("users",users)
	response := map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   users,
	}

	// Convert to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Store in cache for 5 minutes
	err = db.SetCache("users", string(jsonResponse), 5*time.Minute)
	if err != nil {
		log.Printf("Error setting cache: %v", err)
	}

	w.Write(jsonResponse)
}
