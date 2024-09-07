package main

import (
	"encoding/json"
	"net/http"
)

// handler for sending a notification
func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notificationRequest NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&notificationRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Call sendMessage function with the global producer
	if err := sendMessage(notificationRequest.Message); err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	// handle notification separately
	handleNotification(notificationRequest.Message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification sent successfully"))
}

// NotificationRequest represents the structure of the notification request
type NotificationRequest struct {
	Message string `json:"message"`
}
