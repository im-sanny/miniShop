package item

import (
	"errors"
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")

	id, err := strconv.Atoi(itemID)
	if err != nil || id <= 0 {
		util.SendError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	if err = h.svc.Delete(id); err != nil {
		if errors.Is(err, util.ErrorNotFound) {
			util.SendError(w, http.StatusNotFound, "item not found")
			return
		}
	}
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "failed to delete item")
		return
	}

	util.SendData(w, http.StatusOK, "Item deleted successfully")
}
