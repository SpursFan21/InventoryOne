package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/inventory", GetInventory).Methods("GET")
	router.HandleFunc("/inventory", CreateInventory).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
