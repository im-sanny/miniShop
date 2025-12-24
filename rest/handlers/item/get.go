package item

import (
	"miniShop/domain"
	"miniShop/util"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		items []*domain.Item
		count int64
		err1  error
		err2  error
	)

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

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		items, err1 = h.svc.Get(page, limit)
	}()

	go func() {
		defer wg.Done()
		count, err2 = h.svc.Count()
	}()

	wg.Wait()

	if err1 != nil || err2 != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	util.SendPage(w, items, page, limit, count)
}
