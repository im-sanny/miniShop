package item

import (
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	itemList1, err := h.svc.Get(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}
	util.SendPage(w, itemList1, page, limit, cnt)
}
