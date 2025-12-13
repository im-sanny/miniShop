package cmd

import (
	"miniShop/handlers"
	"net/http"
)

func initRoutes(mux *http.ServeMux) {
	mux.Handle("GET /items", http.HandlerFunc(handlers.GetItems))
	mux.Handle("GET /items/{itemId}", http.HandlerFunc(handlers.GetItemById))
	mux.Handle("POST /items", http.HandlerFunc(handlers.PostItem))
}
