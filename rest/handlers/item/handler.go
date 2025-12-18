package item

import (
	"miniShop/repo"
	middleware "miniShop/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	itemRepo repo.ItemRepo
}

func NewHandler(middlewares *middleware.Middlewares, itemRepo repo.ItemRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		itemRepo: itemRepo,
	}
}
