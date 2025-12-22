package item

import (
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("itemId") // All values coming from the URL are strings. So we collect the itemId as a string first, then convert it to int so database queries can use it.

	id, err := strconv.Atoi(idStr) // converting idStr into int which was in string form then it'll store in id
	if err != nil {                // id error qual not nil then it has error
		util.SendError(w, http.StatusBadRequest, "Invalid req body") // after getting error this line will return an error response
		return                                                       // from here req will go back or stop
	}

	item, err := h.svc.GetByID(id)
	if item == nil {
		util.SendError(w, http.StatusNotFound, "item not found")
	}

	util.SendData(w, http.StatusOK, item)
}
