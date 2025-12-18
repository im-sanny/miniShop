package item

import (
	"encoding/json"
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var newItem database.Item          // Create a new album struct to hold incoming data
	decoder := json.NewDecoder(r.Body) // Decode JSON request body into newAlbum
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	createItem := database.CreateItem(newItem)
	util.SendData(w, createItem, http.StatusCreated) // Send the created album back to client with 201 Created
}
