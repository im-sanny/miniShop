package rest

import (
	"miniShop/rest/handlers"
	middleware "miniShop/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /items", manager.With(http.HandlerFunc(handlers.GetItemHandler)))
	mux.Handle("GET /items/{itemId}", manager.With(http.HandlerFunc(handlers.GetItemByIDHandler)))
	mux.Handle("POST /items", manager.With(http.HandlerFunc(handlers.CreateItemHandler), middleware.AuthenticateJWT))
	mux.Handle("PUT /items/{itemId}", manager.With(http.HandlerFunc(handlers.UpdateItemHandler), middleware.AuthenticateJWT))
	mux.Handle("DELETE /items/{itemId}", manager.With(http.HandlerFunc(handlers.DeleteItemHandler), middleware.AuthenticateJWT))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUserHandler)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.UserLoginHandler)))
}
