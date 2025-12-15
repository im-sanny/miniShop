package rest

import (
	"miniShop/rest/handlers"
	middleware "miniShop/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /items", manager.With(http.HandlerFunc(handlers.GetItems)))
	mux.Handle("GET /items/{itemId}", manager.With(http.HandlerFunc(handlers.GetItemById)))
	mux.Handle("POST /items", manager.With(http.HandlerFunc(handlers.PostItem)))
}
