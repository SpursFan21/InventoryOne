package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", UserHandler)            // Create, Get, Update, Delete users
	http.HandleFunc("/users/{id}", SingleUserHandler) // Get single user, Update, Delete

	log.Println("User service is running on port 8086...")
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
