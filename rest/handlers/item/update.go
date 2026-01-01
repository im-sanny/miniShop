package item

import (
	"encoding/json"
	"miniShop/domain"
	"miniShop/util"
	"net/http"
	"strconv"
)

type reqUpdateItem struct {
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	itemId := r.PathValue("itemId")

	id, err := strconv.Atoi(itemId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	var req reqUpdateItem
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Name == "" || req.Brand == "" || req.Price <= 0 || req.Image == "" {
		util.SendError(w, http.StatusBadRequest, "invalid input")
		return
	}

	item, err := h.svc.Update(domain.Item{
		ID:    id,
		Name:  req.Name,
		Brand: req.Brand,
		Price: req.Price,
		Image: req.Image,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to update item")
		return
	}

	util.SendData(w, http.StatusOK, item)
}
