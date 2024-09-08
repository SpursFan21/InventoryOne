package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route definitions
	r.HandleFunc("/users", UserHandler).Methods("POST", "GET")                     // Create and List users
	r.HandleFunc("/users/{id}", SingleUserHandler).Methods("GET", "PUT", "DELETE") // Get, Update, Delete single user

	// Start server
	log.Println("User service is running on port 8086...")
	err := http.ListenAndServe(":8086", r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
