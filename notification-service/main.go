package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up Kafka producer or consumer here

	// Example HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Notification Service is running"))
	})

	log.Println("Notification Service is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
