package handlers

import (
	"encoding/json"
	"miniShop/database"
	"miniShop/util"
	"net/http"
	"strconv"
)

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}

	var newItem database.Item
	error := json.NewDecoder(r.Body).Decode(&newItem)
	if error != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	newItem.ID = id
	database.UpdateItem(newItem)
	util.SendData(w, newItem, http.StatusOK)
}
