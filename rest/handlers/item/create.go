package item

import (
	"encoding/json"
	"miniShop/domain"
	"miniShop/util"
	"net/http"
)

type reqCreateItem struct {
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req reqCreateItem
	decoder := json.NewDecoder(r.Body) // Decode JSON request body into newAlbum
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if req.Name == "" || req.Brand == "" || req.Price <= 0 || req.Image == "" {
		util.SendError(w, http.StatusBadRequest, "invalid input")
		return
	}

	createdItem, err := h.svc.Create(domain.Item{
		Name:  req.Name,
		Brand: req.Brand,
		Price: req.Price,
		Image: req.Image,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "failed to create item")
		return
	}
	util.SendData(w, http.StatusCreated, createdItem)
}
