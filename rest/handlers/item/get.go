package item

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.GetAllItem(), http.StatusOK)
}
