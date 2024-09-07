package main

import (
	"encoding/json"
	"net/http"
)

// OrderRequest represents the structure of an order request
type OrderRequest struct {
	OrderID  string `json:"order_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var orderRequest OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call sendMessage function to send order to Kafka
	if err := sendMessage(orderRequest); err != nil {
		http.Error(w, "Failed to send order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order processed successfully"))
}
