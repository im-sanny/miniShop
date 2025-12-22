package item

import (
	"encoding/json"
	"miniShop/domain"
	"miniShop/util"
	"net/http"
	"strconv"
)

type ReqUpdateItem struct {
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	var req ReqUpdateItem
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid json")
		return
	}

	item, err := h.svc.Update(domain.Item{
		ID:    id,
		Name:  req.Name,
		Brand: req.Brand,
		Price: req.Price,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to update item")
		return
	}

	util.SendData(w, http.StatusOK, item)
}
