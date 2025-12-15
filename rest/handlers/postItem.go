package handlers

import (
	"encoding/json"
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func PostItem(w http.ResponseWriter, r *http.Request) {
	var newItem database.Item // Create a new album struct to hold incoming data
	decoder := json.NewDecoder(r.Body) // Decode JSON request body into newAlbum
	if err := decoder.Decode(&newItem)
	err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	newItem.ID = len(database.ItemList) + 1                // Assign a simple sequential ID (demo only — not production-safe)
	database.ItemList = append(database.ItemList, newItem) // Store in in-memory "database" (global slice)
	util.SendData(w, newItem, http.StatusCreated)          // Send the created album back to client with 201 Created
}
