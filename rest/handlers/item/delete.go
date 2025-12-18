package item

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, "invalid req", http.StatusBadRequest)
		return
	}

	database.DeleteItemById(id)
	util.SendData(w, "Item has been deleted successfully", http.StatusOK)
}
