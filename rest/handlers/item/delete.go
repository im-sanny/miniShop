package item

import (
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid req")
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, "Item deleted successfully")
}
