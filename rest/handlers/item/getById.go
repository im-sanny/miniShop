package item

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("itemId") // All values coming from the URL are strings. So we collect the itemId as a string first, then convert it to int so database queries can use it.

	id, err := strconv.Atoi(idStr) // converting idStr into int which was in string form then it'll store in id
	if err != nil {                // id error qual not nil then it has error
		http.Error(w, "give me valid item id", http.StatusBadRequest) // after getting error this line will return an error response
		return                                                        // from here req will go back or stop
	}

	item := database.GetItemById(id)
	if item == nil {
		util.SendError(w, "item not found", http.StatusNotFound)
	}
	util.SendData(w, item, http.StatusOK)
}
