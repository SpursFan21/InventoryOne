package main

import (
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createUserHandler(w, r)
	case "GET":
		getUsersHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SingleUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUserHandler(w, r)
	case "PUT":
		updateUserHandler(w, r)
	case "DELETE":
		deleteUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := createUser(user)
	if err != nil {
		if err.Error() == "User already exists" {
			http.Error(w, "User already exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := getUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Other handler functions (get single user, update user, delete user) go here
