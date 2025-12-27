package item

import (
	middleware "miniShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /items",
		manager.With(
			http.HandlerFunc(h.Get),
		),
	)

	mux.Handle(
		"GET /items/{itemId}",
		manager.With(
			http.HandlerFunc(h.GetByID),
		),
	)

	mux.Handle(
		"POST /items",
		manager.With(
			http.HandlerFunc(h.Create),
			middleware.AuthenticateJWT,
		))

	mux.Handle(
		"PUT /items/{itemId}",
		manager.With(
			http.HandlerFunc(h.Update),
			middleware.AuthenticateJWT,
		))

	mux.Handle(
		"DELETE /items/{itemId}",
		manager.With(
			http.HandlerFunc(h.Delete),
			middleware.AuthenticateJWT,
		))

}
