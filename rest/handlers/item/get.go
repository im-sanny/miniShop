package item

import (
	"miniShop/domain"
	"miniShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	type itemRes struct {
		items []*domain.Item
		err   error
	}

	type countRes struct {
		count int64
		err   error
	}

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err != nil || limit <= 0 {
		limit = 10
	}

	itemCh := make(chan itemRes, 1)
	countCh := make(chan countRes, 1)

	go func() {
		items, err := h.svc.Get(page, limit)
		itemCh <- itemRes{items, err}
	}()

	go func() {
		count, err := h.svc.Count()
		countCh <- countRes{count, err}
	}()

	items := <-itemCh
	count := <-countCh

	if items.err != nil || count.err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendPage(w, items.items, page, limit, count.count)
}
