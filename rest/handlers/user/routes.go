package user

import (
	middleware "miniShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.Create)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}
