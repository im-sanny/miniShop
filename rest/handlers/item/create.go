package item

import (
	"encoding/json"
	"miniShop/repo"
	"miniShop/util"
	"net/http"
)

type ReqCreateItem struct {
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateItem
	decoder := json.NewDecoder(r.Body) // Decode JSON request body into newAlbum
	if err := decoder.Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	createdItem, err := h.itemRepo.Create(repo.Item{
		Name:  req.Name,
		Brand: req.Brand,
		Price: req.Price,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusCreated, createdItem)
}
