package item

import (
	"miniShop/util"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	itemList, err := h.itemRepo.Get()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}
	util.SendData(w, http.StatusOK, itemList)
}
