package main

import (
	"encoding/json"
	"net/http"
)

type InventoryItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

var inventory = make(map[string]InventoryItem)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	items := make([]InventoryItem, 0, len(inventory))
	for _, item := range inventory {
		items = append(items, item)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func CreateInventory(w http.ResponseWriter, r *http.Request) {
	var item InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inventory[item.ID] = item
	w.WriteHeader(http.StatusCreated)
}
