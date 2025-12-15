package handlers

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
	"strconv"
)

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, "invalid req", http.StatusBadRequest)
		return
	}

	database.DeleteItemById(id)
	util.SendData(w, "Item has been deleted successfully", http.StatusOK)
}
