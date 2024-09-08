package main

import (
	"net/http"
)

// Handler for creating, listing, and updating users
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createUser(w, r) // Admins only
	case "GET":
		getUsers(w, r) // Get all users
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler for getting, updating, and deleting a specific user by ID
func SingleUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUser(w, r)
	case "PUT":
		updateUser(w, r) // Admins only
	case "DELETE":
		deleteUser(w, r) // Admins only
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
